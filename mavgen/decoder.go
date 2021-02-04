/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

// decoderTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func decoderTemplate() string {
	var tmpl = "package mavlink\n" +
		"\n" +
		"import (\n" +
		"\t\"bufio\"\n" +
		"\t\"io\"\n" +
		"    \"{{.CommonPackageURL}}/packet\"\n" +
		"    \"{{.CommonPackageURL}}/parser\"\n" +
		")\n" +
		"\n" +
		"// Decoder struct provide decoding processor\n" +
		"type Decoder struct {\n" +
		"\treader      io.ByteReader\n" +
		"\tparsers     []parser.Parser\n" +
		"}\n" +
		"\n" +
		"func (d *Decoder) clearParser(parser parser.Parser) {\n" +
		"\tparser.Destroy()\n" +
		"}\n" +
		"\n" +
		"func (d *Decoder) clearParsers() {\n" +
		"\tfor _, parser := range d.parsers {\n" +
		"\t\td.clearParser(parser)\n" +
		"\t}\n" +
		"\td.parsers = d.parsers[:0]\n" +
		"}\n" +
		"\n" +
		"// Decode decode input stream to packet. Method return error or nil\n" +
		"func (d *Decoder) Decode() (packet.Packet, error) {\n" +
		"\tfor {\n" +
		"\t\tc, err := d.reader.ReadByte()\n" +
		"\t\tif err != nil {\n" +
		"\t\t\treturn nil, err\n" +
		"\t\t}\n" +
		"\t\tswitch packet.MAGIC_NUMBER(c) {\n" +
		"\t\tcase parser.MAGIC_NUMBER_V1: // mavlink1\n" +
		"\t\t\td.parsers = append(d.parsers, parser.NewParserV1())\n" +
		"\t\tcase parser.MAGIC_NUMBER_V2: // mavlink2\n" +
		"\t\t\td.parsers = append(d.parsers, parser.NewParserV2())\n" +
		"\t\t}\n" +
		"\t\tparsers := make([]parser.Parser, 0, len(d.parsers))\n" +
		"\t\tfor _, parser := range d.parsers {\n" +
		"\t\t\tif p, err := parser.ParseChar(c); err != nil {\n" +
		"\t\t\t\td.clearParser(parser)\n" +
		"\t\t\t} else if !p.Nil() {\n" +
		"\t\t\t\td.clearParsers()\n" +
		"\t\t\t\treturn p, nil\n" +
		"\t\t\t} else {\n" +
		"\t\t\t\tparsers = append(parsers, parser)\n" +
		"\t\t\t}\n" +
		"\t\t}\n" +
		"\t\td.parsers = parsers\n" +
		"\t}\n" +
		"}\n" +
		"\n" +
		"func byteReader(r io.Reader) io.ByteReader {\n" +
		"\tif rb, ok := r.(io.ByteReader); ok {\n" +
		"\t\treturn rb\n" +
		"\t}\n" +
		"    return bufio.NewReader(r)\n" +
		"}\n" +
		"\n" +
		"// NewDecoder function create decoder instance\n" +
		"func NewDecoder(r io.Reader) *Decoder {\n" +
		"\treturn &Decoder{\n" +
		"\t\treader:  byteReader(r),\n" +
		"\t\tparsers: make([]parser.Parser, 0),\n" +
		"\t}\n" +
		"}\n" +
		""
	return tmpl
}
