/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

// encoderTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func encoderTemplate() string {
	var tmpl = "package mavlink\n" +
		"\n" +
		"import (\n" +
		"\t\"fmt\"\n" +
		"\t\"io\"\n" +
		"    \"{{.CommonPackageURL}}/parser\"\n" +
		"    \"{{.CommonPackageURL}}/packet\"\n" +
		"    \"{{.CommonPackageURL}}/message\"\n" +
		"    \"{{.CommonPackageURL}}/version\"\n" +
		")\n" +
		"\n" +
		"// Encoder struct provide decoding processor\n" +
		"type Encoder struct {\n" +
		"\twriter  io.Writer\n" +
		"\tcurrentSeqNum byte\n" +
		"}\n" +
		"\n" +
		"func (e *Encoder) nextSeqNum() byte {\n" +
		"\te.currentSeqNum++\n" +
		"\treturn e.currentSeqNum\n" +
		"}\n" +
		"\n" +
		"// Encode encode packet to output stream. Method return error or nil\n" +
		"func (e *Encoder) Encode(v version.MAVLINK_VERSION, sysID uint8, compID uint8, message message.Message) error {\n" +
		"    p, err := MakePacket(v, sysID, compID, e.nextSeqNum(), message)\n" +
		"\tif err != nil {\n" +
		"\t    return err\n" +
		"\t}\n" +
		"\tb, err := p.Marshal()\n" +
		"\tif err != nil {\n" +
		"\t    return err\n" +
		"\t}\n" +
		"\tn, err := e.writer.Write(b)\n" +
		"\tif len(b) != n {\n" +
		"\t\treturn fmt.Errorf(\"writed %d bytes, but need to write %d bytes\", n, len(b))\n" +
		"\t}\n" +
		"\treturn err\n" +
		"}\n" +
		"\n" +
		"// NewEncoder function create encoder instance\n" +
		"func NewEncoder(w io.Writer) *Encoder {\n" +
		"\treturn &Encoder{\n" +
		"\t\twriter:        w,\n" +
		"\t\tcurrentSeqNum: 0,\n" +
		"\t}\n" +
		"}\n" +
		"\n" +
		"\n" +
		"func MakePacket(v version.MAVLINK_VERSION, sysID uint8, compID uint8, seqID uint8, message message.Message) (packet.Packet, error) {\n" +
		"\tswitch v {\n" +
		"\tcase version.MAVLINK_V1:\n" +
		"\t\treturn parser.MakePacketV1(sysID, compID, seqID, message)\n" +
		"\tcase version.MAVLINK_V2:\n" +
		"\t\treturn parser.MakePacketV2(sysID, compID, seqID, message)\n" +
		"\tdefault:\n" +
		"\t\treturn nil, fmt.Errorf(\"Unknown mavlink version %d\", v)\n" +
		"\t}\n" +
		"}\n" +
		""
	return tmpl
}
