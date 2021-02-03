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
		"func (e *Encoder) makePacket(version MAVLINK_VERSION, sysID uint8, compID uint8, message Message) (Packet, error) {\n" +
		"\tpayload, err := message.Marshal()\n" +
		"\tif err != nil {\n" +
		"\t    return nil, err\n" +
		"\t}\n" +
		"    switch version {\n" +
		"    case MAVLINK_V1:\n" +
		"        return &packet1{\n" +
		"            seqID: e.nextSeqNum(),\n" +
		"            sysID: sysID,\n" +
		"            compID: compID,\n" +
		"            msgID: message.MsgID(),\n" +
		"            payload: payload,\n" +
		"        }, nil\n" +
		"    case MAVLINK_V2:\n" +
		"        return &packet2{\n" +
		"            seqID: e.nextSeqNum(),\n" +
		"            sysID: sysID,\n" +
		"            compID: compID,\n" +
		"            msgID: message.MsgID(),\n" +
		"            payload: payload,\n" +
		"        }, nil\n" +
		"    default:\n" +
		"        return nil, fmt.Errorf(\"Undefined mavlink version %d\", version)\n" +
		"    }\n" +
		"}\n" +
		"\n" +
		"// Encode encode packet to output stream. Method return error or nil\n" +
		"func (e *Encoder) Encode(version MAVLINK_VERSION, sysID uint8, compID uint8, message Message) error {\n" +
		"    packet, err := e.makePacket(version, sysID, compID, message)\n" +
		"\tif err != nil {\n" +
		"\t    return err\n" +
		"\t}\n" +
		"\tb, err := packet.Marshal()\n" +
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
		""
	return tmpl
}
