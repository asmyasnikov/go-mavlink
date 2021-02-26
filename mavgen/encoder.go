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
		"\t\"time\"\n" +
		"    \"{{.CommonPackageURL}}/parser\"\n" +
		"    \"{{.CommonPackageURL}}/packet\"\n" +
		"    \"{{.CommonPackageURL}}/message\"\n" +
		"    \"{{.CommonPackageURL}}/version\"\n" +
		")\n" +
		"\n" +
		"// Encoder struct provide decoding processor\n" +
		"type Encoder struct {\n" +
		"\twriter  io.Writer\n" +
		"}\n" +
		"\n" +
		"var currentSeqNum uint8\n" +
		"\n" +
		"func nextSeqNum() byte {\n" +
		"\tcurrentSeqNum++\n" +
		"\treturn currentSeqNum\n" +
		"}\n" +
		"\n" +
		"// Encode encode packet to output stream. Method return error or nil on success\n" +
		"func (e *Encoder) Encode(p packet.Packet) error {\n" +
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
		"// EncodeWithSignature encode packet with signature to output stream. Method return error or nil on success\n" +
		"func (e *Encoder) EncodeWithSignature(p packet.Packet, linkID byte, timestamp time.Time, secretKey [32]byte) error {\n" +
		"\tb, err := p.MarshalWithSignature(linkID, timestamp, secretKey)\n" +
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
		"// NewEncoder function creates encoder instance\n" +
		"func NewEncoder(w io.Writer) *Encoder {\n" +
		"\treturn &Encoder{\n" +
		"\t\twriter:        w,\n" +
		"\t}\n" +
		"}\n" +
		"\n" +
		"\n" +
		"// NewPacket function creates packet\n" +
		"func NewPacket(v version.MAVLINK_VERSION, sysID uint8, compID uint8, message message.Message) (packet.Packet, error) {\n" +
		"\tswitch v {\n" +
		"\tcase version.MAVLINK_V1:\n" +
		"\t\treturn parser.NewPacketV1(sysID, compID, nextSeqNum(), message)\n" +
		"\tcase version.MAVLINK_V2:\n" +
		"\t\treturn parser.NewPacketV2(sysID, compID, nextSeqNum(), message)\n" +
		"\tdefault:\n" +
		"\t\treturn nil, fmt.Errorf(\"Unknown mavlink version %d\", v)\n" +
		"\t}\n" +
		"}\n" +
		""
	return tmpl
}
