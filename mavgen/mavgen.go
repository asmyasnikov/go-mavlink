package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/howeyc/crc16"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

var (
	upperCaseWords = map[string]string{
		"Id":   "ID",
		"Cpu":  "CPU",
		"Uri":  "URI",
		"Url":  "URL",
		"Uid":  "UID",
		"Http": "HTTP",
		"Udp":  "UDP",
	}
)

// Dialect described root tag of schema
type Dialect struct {
	FilePath string
	XMLName  xml.Name   `xml:"mavlink"`
	Version  string     `xml:"version"`
	Include  []string   `xml:"include"`
	Enums    []*Enum    `xml:"enums>enum"`
	Messages []*Message `xml:"messages>message"`
}

// Enum described schema tag enum
type Enum struct {
	Name        string       `xml:"name,attr"`
	Description string       `xml:"description"`
	Entries     []*EnumEntry `xml:"entry"`
}

// EnumEntry described schema tag entry
type EnumEntry struct {
	Value       string            `xml:"value,attr"`
	Name        string            `xml:"name,attr"`
	Description string            `xml:"description"`
	Params      []*EnumEntryParam `xml:"param"`
}

// EnumEntryParam described schema tag param
type EnumEntryParam struct {
	Index       uint8  `xml:"index,attr"`
	Description string `xml:",innerxml"`
}

// Message described schema tag message
type Message struct {
	// use uint32 instead of uint8 so that we can filter
	// msgids from mavlink v2, which are 24 bits wide.
	// see filtering in ParseDialect()
	ID          uint32          `xml:"id,attr"`
	Name        string          `xml:"name,attr"`
	Description string          `xml:"description"`
	Fields      []*MessageField `xml:"field"`

	// this field is only used during ParseDialect phase,
	// it contains an empty string after ParseDialect returns
	Raw         string `xml:",innerxml"`
	DialectName string
}

// MessageField described schema tag filed
type MessageField struct {
	CType       string `xml:"type,attr"`
	Name        string `xml:"name,attr"`
	Enum        string `xml:"enum,attr"`
	Display     string `xml:"display,attr"`
	Description string `xml:",innerxml"`
	GoType      string
	Tags        map[string]string
	BitSize     int
	ArrayLen    int
	ByteOffset  int // from beginning of payload
}

var funcMap = template.FuncMap{
	"UpperCamelCase":   UpperCamelCase,
	"IsByteArrayField": IsByteArrayField,
	"IsStringField":    IsStringField,
}

// SizeInBytes function calculate size in bytes of message field
func (f *MessageField) SizeInBytes() int {
	if f.ArrayLen > 0 {
		return f.BitSize / 8 * f.ArrayLen
	}
	return f.BitSize / 8
}

// Size function calculate size in bytes of message
func (m *Message) Size() int {
	sz := 0
	for _, f := range m.Fields {
		sz += f.SizeInBytes()
	}
	return sz
}

// CRCExtra calculation: http://www.mavlink.org/mavlink/crc_extra_calculation
func (m *Message) CRCExtra() uint8 {
	hash := crc16.New(crc16.CCITTTable)

	fmt.Fprint(hash, m.Name+" ")
	for _, f := range m.Fields {
		cType := f.CType
		if cType == "uint8_t_mavlink_version" {
			cType = "uint8_t"
		}
		// type name for crc extra purposes does not include array portion
		if idx := strings.IndexByte(cType, '['); idx >= 0 {
			cType = cType[:idx]
		}
		fmt.Fprint(hash, cType+" "+f.Name+" ")
		if f.ArrayLen > 0 {
			_, _ = hash.Write([]byte{byte(f.ArrayLen)})
		}
	}

	crc := hash.Sum16()
	return uint8((crc & 0xFF) ^ (crc >> 8))
}

// implementation of sort.Interface for Message
func (m *Message) Len() int {
	return len(m.Fields)
}

func (m *Message) Less(i, j int) bool {
	return m.Fields[i].BitSize < m.Fields[j].BitSize
}

func (m *Message) Swap(i, j int) {
	m.Fields[i], m.Fields[j] = m.Fields[j], m.Fields[i]
}

// IsByteArrayField function check field is bytearray
func IsByteArrayField(v interface{}) bool {
	if field, ok := v.(*MessageField); ok {
		if field.ArrayLen == 0 {
			return false
		}
		t := strings.ToLower(field.GoType)
		for _, s := range []string{"byte", "uint8"} {
			if strings.Contains(t, s) {
				return true
			}
		}
	}
	return false
}

// IsStringField function check field is bytearray
func IsStringField(v interface{}) bool {
	if field, ok := v.(*MessageField); ok {
		return field.GoType == "string"
	}
	return false
}

// UpperCamelCase function convert names to upper camel case
func UpperCamelCase(s string) string {
	var b bytes.Buffer
	for _, fragment := range strings.Split(s, "_") {
		if len(fragment) > 0 {
			word := strings.ToUpper(fragment[:1]) + strings.ToLower(fragment[1:])
			if replacement, ok := upperCaseWords[word]; ok {
				word = replacement
			}
			b.WriteString(word)
		}
	}
	return b.String()
}

// helper to pack a single element into a payload.
// can be called for a single field, or an element within a field's array.
func (f *MessageField) payloadPackPrimitive(offset, name string) string {

	if f.BitSize == 8 {
		return fmt.Sprintf("payload[%s] = byte(%s)", offset, name)
	}

	if f.IsFloat() {
		switch f.BitSize {
		case 32, 64:
			return fmt.Sprintf("binary.LittleEndian.PutUint%d(payload[%s:], math.Float%dbits(%s))", f.BitSize, offset, f.BitSize, name)
		}
	} else {
		switch f.BitSize {
		case 16, 32, 64:
			return fmt.Sprintf("binary.LittleEndian.PutUint%d(payload[%s:], uint%d(%s))", f.BitSize, offset, f.BitSize, name)
		}
	}

	panic("unhandled bitsize")
}

// PayloadPackSequence function produce a string that will pack
// this message's fields into a byte slice called 'payload'
func (f *MessageField) PayloadPackSequence() string {
	name := UpperCamelCase(f.Name)

	if f.ArrayLen > 0 {
		// optimize to copy() if possible
		if f.GoType == "string" || strings.HasSuffix(f.GoType, "uint8") || strings.HasSuffix(f.GoType, "byte") {
			return fmt.Sprintf(`copy(payload[%d:], m.%s)`, f.ByteOffset, name)
		}

		// pack each element in the array
		s := fmt.Sprintf("for i, v := range m.%s {\n", name)
		off := fmt.Sprintf("%d + i * %d", f.ByteOffset, f.BitSize/8)
		s += f.payloadPackPrimitive(off, "v") + "\n"
		s += fmt.Sprintf("}")
		return s
	}

	// pack a single field
	return f.payloadPackPrimitive(fmt.Sprintf("%d", f.ByteOffset), "m."+name)
}

func (f *MessageField) payloadUnpackPrimitive(offset string) string {

	if f.BitSize == 8 {
		if len(f.Enum) > 0 {
			return fmt.Sprintf("%s(payload[%s])", goArrayType(f.Enum), offset)
		}
		return fmt.Sprintf("%s(payload[%s])", goArrayType(f.GoType), offset)
	}

	if f.IsFloat() {
		switch f.BitSize {
		case 32, 64:
			return fmt.Sprintf("math.Float%dfrombits(binary.LittleEndian.Uint%d(payload[%s:]))", f.BitSize, f.BitSize, offset)
		}
	} else {
		switch f.BitSize {
		case 16, 32, 64:
			if len(f.Enum) > 0 {
				return fmt.Sprintf("%s(binary.LittleEndian.Uint%d(payload[%s:]))", goArrayType(f.Enum), f.BitSize, offset)
			}
			return fmt.Sprintf("%s(binary.LittleEndian.Uint%d(payload[%s:]))", goArrayType(f.GoType), f.BitSize, offset)
		}
	}

	panic("unhandled bitsize")
}

// PayloadUnpackSequence function produce a string that will unpack
// this message's fields into a byte slice called 'payload'
func (f *MessageField) PayloadUnpackSequence() string {
	name := UpperCamelCase(f.Name)
	if f.ArrayLen > 0 {
		if f.GoType == "string" {
			return fmt.Sprintf("m.%s = string(payload[%d:%d])", name, f.ByteOffset, f.ByteOffset+f.ArrayLen)
		}
		// optimize to copy() if possible
		if strings.HasSuffix(f.GoType, "uint8") || strings.HasSuffix(f.GoType, "byte") {
			return fmt.Sprintf("copy(m.%s[:], payload[%d:%d])", name, f.ByteOffset, f.ByteOffset+f.ArrayLen)
		}

		// unpack each element in the array
		s := fmt.Sprintf("for i := 0; i < len(m.%s); i++ {\n", name)
		off := fmt.Sprintf("%d + i * %d", f.ByteOffset, f.BitSize/8)
		s += fmt.Sprintf("m.%s[i] = %s\n", name, f.payloadUnpackPrimitive(off))
		s += fmt.Sprintf("}")
		return s
	}
	return fmt.Sprintf("m.%s = %s", name, f.payloadUnpackPrimitive(fmt.Sprintf("%d", f.ByteOffset)))
}

// Return the number of non-extension fields, that are contained in rawmsg, where rawmsg
// is raw XML content of "message" element.
func numBaseFields(rawmsg string) int {
	dec := xml.NewDecoder(strings.NewReader(rawmsg))

	fields := 0
	for {
		t, err := dec.Token()
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			return fields
		}
		if se, ok := t.(xml.StartElement); ok {
			switch se.Name.Local {
			case "field":
				fields++
			case "extensions":
				return fields
			}
			if err := dec.Skip(); err != nil {
				panic(err)
			}
		}
	}
}

// parseDialect read in an xml-based dialect stream,
// and populate a Dialect struct with its contents
func parseDialect(in io.Reader) (*Dialect, error) {
	filebytes, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}

	dialect := &Dialect{}

	if err := xml.Unmarshal(filebytes, &dialect); err != nil {
		return nil, err
	}

	// filter out messages with MSG_ID > 256. these are from
	// mavlink v2 and do not fit in uint8
	filteredMessages := make([]*Message, len(dialect.Messages))
	n := 0
	for _, msg := range dialect.Messages {
		if msg.ID <= 0xff {
			// we ignore field extensions, since they are from
			// mavlink v2
			ind := numBaseFields(msg.Raw)
			if ind >= 0 {
				msg.Fields = msg.Fields[:ind]
			}
			filteredMessages[n] = msg
			n++
		}
		msg.Raw = ""
	}
	dialect.Messages = filteredMessages[:n]

	return dialect, nil
}

// ParseDialect read in an xml-based dialect file,
// and populate a Dialect struct with its contents
func ParseDialect(schemeFile string) (*Dialect, error) {
	in, err := os.Open(schemeFile)
	if err != nil {
		return nil, err
	}
	defer in.Close()
	d, err := parseDialect(in)
	if err != nil {
		return nil, err
	}
	for _, i := range d.Include {
		includedPath := filepath.Join(filepath.Dir(schemeFile), i)
		included, err := ParseDialect(includedPath)
		if err != nil {
			return nil, err
		}
		included.FilePath = includedPath
		if err := d.merge(included); err != nil {
			return nil, err
		}
	}
	return d, nil
}

var reCType = regexp.MustCompile(`^(.+?)([0-9]+)_t$`)

// convert a C primitive type to its corresponding Go type.
// do not handle arrays or other constructs...just primitives.
func c2goPrimitive(ctype string) (string, int) {
	switch ctype {
	case "uint8_t", "uint16_t", "uint32_t", "uint64_t",
		"int8_t", "int16_t", "int32_t", "int64_t":
		if matches := reCType.FindStringSubmatch(ctype); matches != nil {
			return matches[1] + matches[2], func() int {
				v, _ := strconv.Atoi(matches[2])
				return v
			}()
		}
	case "char":
		return "byte", 8
	case "float":
		return "float32", 32
	case "double":
		return "float64", 64
	case "uint8_t_mavlink_version":
		return "uint8", 8
	}
	panic(fmt.Sprintf("c2goPrimitive: unhandled primitive type - %s", ctype))
}

func goArrayType(s string) string {
	idx := strings.IndexByte(s, ']')
	if idx < 0 {
		return s
	}
	return s[idx+1:]
}

// IsFloat return check state of validating field as float type
func (f *MessageField) IsFloat() bool {
	return strings.HasPrefix(goArrayType(f.GoType), "float")
}

// GoTypeInfo produce type info string
func GoTypeInfo(s string) (string, int, int) {
	// array? leave the [N] but convert the primitive type name
	if matches := reTypeIsArray.FindStringSubmatch(s); matches != nil {
		arrayLen, _ := strconv.Atoi(matches[2])
		if matches[1] == "char" {
			return "string", 8, arrayLen
		}
		name, bitSize := c2goPrimitive(matches[1])
		name = "[]" + name
		return name, bitSize, arrayLen
	}
	name, bitSize := c2goPrimitive(s)
	return name, bitSize, 0
}

func (d *Dialect) needImportParentMavlink() bool {
	return len(d.Messages) > 0
}

func (d *Dialect) needImportFmt() bool {
	return len(d.Messages) > 0
}

func (d *Dialect) needImportMath() bool {
	for _, m := range d.Messages {
		for _, f := range m.Fields {
			switch f.CType {
			case "float", "double":
				return true
			}
		}
	}
	return false
}

func (d *Dialect) needImportEncodingBinary() bool {
	for _, m := range d.Messages {
		for _, f := range m.Fields {
			switch f.CType {
			case "uint16_t", "uint32_t", "uint64_t":
				return true
			}
		}
	}
	return false
}

func (d *Dialect) generateFile(filePath string, packageName string, commonPackage string, imports []string, generator func(w io.Writer) error) error {
	var bb bytes.Buffer
	bb.WriteString(generatedHeader)
	bb.WriteString("package " + strings.ToLower(packageName) + "\n\n")
	if len(imports) > 0 {
		bb.WriteString("import (\n")
		for _, i := range imports {
			bb.WriteString("\"" + i + "\"\n")
		}
		bb.WriteString(")\n")
	}
	if err := generator(&bb); err != nil {
		return err
	}
	formatted, err := format.Source(bb.Bytes())
	if err != nil {
		formatted = bb.Bytes()
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	n, err := f.Write(formatted)
	if err == nil && n != len(formatted) {
		return io.ErrShortWrite
	}
	return err
}

func (d *Dialect) generateGo(dialectPath string, packageName string, commonPackage string) error {
	if err := d.generateFile(
		filepath.Join(dialectPath, "version.go"),
		packageName,
		commonPackage,
		nil,
		d.generateVersion,
	); err != nil {
		return err
	}
	if err := d.generateFile(
		filepath.Join(dialectPath, "enums.go"),
		packageName,
		commonPackage,
		func() []string {
			if len(d.Enums) > 0 {
				return []string{
					"fmt",
				}
			}
			return nil
		}(),
		d.generateEnums,
	); err != nil {
		return err
	}
	if err := d.generateFile(
		filepath.Join(dialectPath, "classes.go"),
		packageName,
		commonPackage,
		append(
			func() []string {
				if len(d.Messages) > 0 {
					return []string{
						"fmt",
						commonPackage + "/message",
					}
				}
				return nil
			}(),
			append(
				func() []string {
					for _, m := range d.Messages {
						for _, f := range m.Fields {
							switch f.CType {
							case "float", "double":
								return []string{"math"}
							}
						}
					}
					return nil
				}(),
				func() []string {
					for _, m := range d.Messages {
						for _, f := range m.Fields {
							switch f.CType {
							case "uint16_t", "uint32_t", "uint64_t":
								return []string{"encoding/binary"}
							}
						}
					}
					return nil
				}()...,
			)...,
		),
		d.generateClasses,
	); err != nil {
		return err
	}
	if err := d.generateFile(
		filepath.Join(dialectPath, "messages.go"),
		packageName,
		commonPackage,
		func() []string {
			if len(d.Messages) > 0 {
				return []string{
					commonPackage + "/message",
				}
			}
			return nil
		}(),
		d.generateMsgIds,
	); err != nil {
		return err
	}
	if err := d.generateFile(
		filepath.Join(dialectPath, "init.go"),
		packageName,
		commonPackage,
		func() []string {
			if len(d.Messages) > 0 {
				return []string{
					commonPackage + "/message",
					commonPackage + "/packet",
					commonPackage + "/register",
				}
			}
			return nil
		}(),
		d.generateInit,
	); err != nil {
		return err
	}
	return nil
}

func (d *Dialect) generateEnums(w io.Writer) error {
	enumTmpl := `
{{range .Enums}}
{{$enumName := .Name}}
// {{$enumName}} type{{if .Description}}. {{.Description}}{{end}}
type {{.Name}} int

const ({{range .Entries}}
	// {{.Name}} enum{{if .Description}}. {{.Description}}{{end}}
	{{.Name}} {{$enumName}} = {{.Value}} {{end}}
)

func (e {{$enumName}}) String() string {
	if name, ok := map[{{$enumName}}]string{ {{range .Entries}}
		{{.Name}}: "{{.Name}}", {{end}} 
	}[e]; ok {
		return name
	}
	return fmt.Sprintf("{{$enumName}}_UNDEFINED_%d", int(e))
}

// Bitmask return string representetion of intersects {{$enumName}} enums 
func (e {{$enumName}}) Bitmask() string {
	bitmap := ""
	for _, entry := range []{{$enumName}}{ {{range .Entries}}
		{{.Name}}, {{end}} 
	} {
		if e & entry > 0 {
			if len(bitmap) > 0 {
				bitmap += " | "
			}
			bitmap += entry.String()
		}
	}
	return bitmap
}
{{end}}
`
	// fill in missing enum values if necessary, and ensure description strings are valid.
	for _, e := range d.Enums {
		e.Description = strings.Replace(e.Description, "\n", " ", -1)
		for i, ee := range e.Entries {
			if ee.Value == "" {
				ee.Value = strconv.Itoa(i)
			}
			ee.Description = strings.Trim(strings.ReplaceAll(ee.Description, "\n", " "), " .")
			if len(ee.Params) > 0 {
				if len(ee.Description) > 0 {
					ee.Description += ". "
				}
				ee.Description += "Params: "
			}
			for _, pp := range ee.Params {
				ee.Description += strconv.Itoa(int(pp.Index)) + ") " + pp.Description + "; "
			}
		}
	}

	return template.Must(template.New("enums").Parse(enumTmpl)).Execute(w, d)
}

func (d *Dialect) generateVersion(w io.Writer) error {
	versionTmpl := `
const (
    // Version of mavgen which generate this code or user defined (by mavgen flag -v) version of dialect
    Version        = "{{- .Version -}}"
)
`
	return template.Must(template.New("version").Parse(versionTmpl)).Execute(w, d)
}

func (d *Dialect) generateMsgIds(w io.Writer) error {
	msgIDTmpl := `
{{if .Messages}}
// Message IDs
const ({{range .Messages}}
	MSG_ID_{{.Name}} message.MessageID = {{.ID}}{{end}}
)
{{end}}
`
	return template.Must(template.New("msgIds").Funcs(funcMap).Parse(msgIDTmpl)).Execute(w, d)
}

func (d *Dialect) generateInit(w io.Writer) error {
	initTmpl := `
{{if .Messages}}
func init() { {{range .Messages}} 
	register.Register(MSG_ID_{{.Name}}, "MSG_ID_{{.Name}}", {{.Size}}, {{.CRCExtra}}, func(p packet.Packet) (message.Message, error) {
		msg := new({{.Name | UpperCamelCase}})
		if err := msg.Unmarshal(p.Payload()); err != nil {
			return nil, err
		}
		return msg, nil
	}){{end}}
} {{end}}
`
	return template.Must(template.New("init").Funcs(funcMap).Parse(initTmpl)).Execute(w, d)
}

var reTypeIsArray = regexp.MustCompile(`^(.+?)\[([0-9]+)\]$`)

// generate class definitions for each msg id.
// for now, pack/unpack payloads via encoding/binary since it
// is expedient and correct. optimize this if/when needed.
func (d *Dialect) generateClasses(w io.Writer) error {
	classesTmpl := `
{{range .Messages}}
{{$name := .Name | UpperCamelCase}}
// {{$name}} struct (generated typeinfo)  
// {{.Description}}
type {{$name}} struct { {{range .Fields}}
  {{.Name | UpperCamelCase}} {{if .Enum}} {{.Enum}} {{ else }} {{.GoType}} {{ end }} {{if .Tags}} ` + "`" + `{{range $k, $v := .Tags}}{{$k}}:"{{$v}}" {{end}}` + "`" + `{{end}}// {{.Description}}{{end}}
}

// MsgID (generated function)
func (m *{{$name}}) MsgID() message.MessageID {
	return MSG_ID_{{.Name}}
}

// String (generated function)
func (m *{{$name}}) String() string {
	return fmt.Sprintf(
		"&{{.DialectName}}.{{$name}}{ {{range $i, $v := .Fields}}{{if gt $i 0}}, {{end}}{{.Name | UpperCamelCase}}: {{if IsStringField .}}\"%s\"{{else}}{{if IsByteArrayField .}}%0X (\"%s\"){{else}}%+v{{if .Enum}}{{if eq .Display "bitmask"}} (%0{{.BitSize}}b){{end}}{{end}}{{end}}{{end}}{{end}} }", 
		{{range .Fields}}m.{{.Name | UpperCamelCase}}{{if .Enum}}{{if eq .Display "bitmask"}}.Bitmask(), uint64(m.{{.Name | UpperCamelCase}}){{end}}{{end}}{{if IsByteArrayField .}}, string(m.{{.Name | UpperCamelCase}}[:]){{end}},
{{end}}
	)
}

// Marshal (generated function)
func (m *{{$name}}) Marshal() ([]byte, error) {
	payload := make([]byte, {{ .Size }}){{range .Fields}}
	{{.PayloadPackSequence}}{{end}}
	return payload, nil
}

// Unmarshal (generated function)
func (m *{{$name}}) Unmarshal(data []byte) error {
	payload := make([]byte, {{.Size}})
	copy(payload[0:], data){{range .Fields}}
	{{.PayloadUnpackSequence}}{{end}}
	return nil
}
{{end}}
`
	for _, m := range d.Messages {
		m.Description = strings.ReplaceAll(m.Description, "\n", "\n// ")
		if len(m.DialectName) == 0 {
			m.DialectName = baseName(d.FilePath)
		}
		for _, f := range m.Fields {
			f.Description = strings.ReplaceAll(f.Description, "\n", " ")
			f.Tags = map[string]string{}
			goname, gosz, golen := GoTypeInfo(f.CType)
			//if len(f.Enum) > 0 {
			//	f.Tags["raw"] = f.CType
			//}
			if golen > 0 {
				f.Tags["len"] = strconv.Itoa(golen)
			}
			f.GoType, f.BitSize, f.ArrayLen = goname, gosz, golen
		}

		// ensure fields are sorted according to their size,
		// http://www.mavlink.org/mavlink/crc_extra_calculation
		sort.Stable(sort.Reverse(m))

		// once sorted, calculate offsets for use in payload packing/unpacking
		offset := 0
		for _, f := range m.Fields {
			f.ByteOffset = offset
			offset += f.SizeInBytes()
		}
	}

	return template.Must(template.New("classesTmpl").Funcs(funcMap).Parse(classesTmpl)).Execute(w, d)
}

func (d *Dialect) merge(rhs *Dialect) error {
	for _, enum := range rhs.Enums {
		if i := d.enumIdx(enum); i < 0 {
			d.Enums = append(d.Enums, enum)
		} else {
			for _, entry := range enum.Entries {
				if j := d.Enums[i].entryIdx(entry); j < 0 {
					d.Enums[i].Entries = append(d.Enums[i].Entries, entry)
				}
			}
		}
	}
	for _, message := range rhs.Messages {
		if i := d.messageIdx(message); i < 0 {
			if len(message.DialectName) == 0 {
				message.DialectName = baseName(rhs.FilePath)
			}
			d.Messages = append(d.Messages, message)
		}
	}
	return nil
}

func (d *Dialect) messageIdx(message *Message) int {
	for i, m := range d.Messages {
		if m.Name == message.Name {
			return i
		}
	}
	return -1
}

func (d *Dialect) enumIdx(enum *Enum) int {
	for i, e := range d.Enums {
		if e.Name == enum.Name {
			return i
		}
	}
	return -1
}

func (enum *Enum) entryIdx(entry *EnumEntry) int {
	for i, e := range enum.Entries {
		if e.Name == entry.Name {
			return i
		}
	}
	return -1
}
