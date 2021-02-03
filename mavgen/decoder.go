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
		")\n" +
		"\n" +
		"// Packet is the interface implemented by frames of every supported version.\n" +
		"type Packet interface {\n" +
		"\t// Nil returns true if packet is nil\n" +
		"\tNil() bool\n" +
		"\t// SysID returns system id\n" +
		"\tSysID() uint8\n" +
		"\t// CompID returns component id\n" +
		"\tCompID() uint8\n" +
		"\t// MsgID returns message id\n" +
		"\tMsgID() MessageID\n" +
		"\t// Checksum returns packet checksum\n" +
		"\tChecksum() uint16\n" +
		"\t// SeqID returns packet sequence number\n" +
		"\tSeqID() uint8\n" +
		"\t// Payload returns packet payload\n" +
		"\tPayload() []byte\n" +
		"\t// Assign assign internal fields from right hand side packet\n" +
		"\t//Assign(rhs Packet) error\n" +
		"\t// Copy returns deep copy of packet\n" +
		"\tCopy() Packet\n" +
		"\t// Message returns dialect message\n" +
		"\tMessage() (Message, error)\n" +
		"\t// String returns string representation of packet\n" +
		"\tString() string\n" +
		"    // Marshal encodes Packet to byte slice\n" +
		"    Marshal() ([]byte, error)\n" +
		"    // Unmarshal parses PAYLOAD and stores the result in Packet\n" +
		"    Unmarshal(payload []byte) error\n" +
		"}\n" +
		"\n" +
		"// Parser interface is abstract of parsers\n" +
		"type Parser interface {\n" +
		"\tParseChar(c byte) (Packet, error)\n" +
		"    Destroy()\n" +
		"}\n" +
		"\n" +
		"// Decoder struct provide decoding processor\n" +
		"type Decoder struct {\n" +
		"\treader      io.ByteReader\n" +
		"\tparsers     []Parser\n" +
		"}\n" +
		"\n" +
		"func (d *Decoder) clearParser(parser Parser) {\n" +
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
		"func (d *Decoder) Decode() (Packet, error) {\n" +
		"\tfor {\n" +
		"\t\tc, err := d.reader.ReadByte()\n" +
		"\t\tif err != nil {\n" +
		"\t\t\treturn nil, err\n" +
		"\t\t}\n" +
		"\t\tswitch c {\n" +
		"\t\tcase 0xfe: // mavlink1\n" +
		"\t\t\td.parsers = append(d.parsers, NewParserV1())\n" +
		"\t\tcase 0xfd: // mavlink2\n" +
		"\t\t\td.parsers = append(d.parsers, NewParserV2())\n" +
		"\t\t}\n" +
		"\t\tparsers := make([]Parser, 0, len(d.parsers))\n" +
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
		"\t\tparsers: make([]Parser, 0),\n" +
		"\t}\n" +
		"}\n" +
		"\n" +
		"func u16ToBytes(v uint16) []byte {\n" +
		"\treturn []byte{byte(v & 0xff), byte(v >> 8)}\n" +
		"}\n" +
		""
	return tmpl
}
