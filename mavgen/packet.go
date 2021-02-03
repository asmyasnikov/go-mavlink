/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

// packetTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func packetTemplate() string {
	var tmpl = "package mavlink\n" +
		"\n" +
		"import (\n" +
		"    \"fmt\"\n" +
		")\n" +
		"\n" +
		"// Packet is a wire type for encoding/decoding mavlink messages.\n" +
		"// use the ToPacket() and FromPacket() routines on specific message\n" +
		"// types to convert them to/from the Message type.\n" +
		"type packet{{.MavlinkVersion}} struct {\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"\tincompatFlags uint8     // incompat flags\n" +
		"\tcompatFlags   uint8     // compat flags\n" +
		"{{- end}}\n" +
		"\tseqID         uint8     // Sequence of packet\n" +
		"\tsysID         uint8     // ID of message sender system/aircraft\n" +
		"\tcompID        uint8     // ID of the message sender component\n" +
		"\tmsgID         MessageID // ID of message in payload\n" +
		"\tpayload       []byte\n" +
		"\tchecksum      uint16\n" +
		"}\n" +
		"\n" +
		"// SysID returns system id\n" +
		"func (p *packet{{.MavlinkVersion}}) SysID() uint8 {\n" +
		"    return p.sysID\n" +
		"}\n" +
		"\n" +
		"// CompID returns component id\n" +
		"func (p *packet{{.MavlinkVersion}}) CompID() uint8 {\n" +
		"    return p.compID\n" +
		"}\n" +
		"\n" +
		"// MsgID returns message id\n" +
		"func (p *packet{{.MavlinkVersion}}) MsgID() MessageID {\n" +
		"    return p.msgID\n" +
		"}\n" +
		"\n" +
		"// Checksum returns packet checksum\n" +
		"func (p *packet{{.MavlinkVersion}}) Checksum() uint16 {\n" +
		"    return p.checksum\n" +
		"}\n" +
		"\n" +
		"// SeqID returns packet sequence number\n" +
		"func (p *packet{{.MavlinkVersion}}) SeqID() uint8 {\n" +
		"    return p.seqID\n" +
		"}\n" +
		"\n" +
		"// Payload returns packet payload\n" +
		"func (p *packet{{.MavlinkVersion}}) Payload() []byte {\n" +
		"    return append([]byte(nil), p.payload...)\n" +
		"}\n" +
		"\n" +
		"func (p *packet{{.MavlinkVersion}}) assign(rhs *packet{{.MavlinkVersion}}) error {\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"    p.incompatFlags = rhs.incompatFlags\n" +
		"    p.compatFlags = rhs.compatFlags\n" +
		"{{- end}}\n" +
		"    p.seqID = rhs.seqID\n" +
		"    p.sysID = rhs.sysID\n" +
		"    p.compID = rhs.compID\n" +
		"    p.msgID = rhs.msgID\n" +
		"    p.checksum = rhs.checksum\n" +
		"    p.payload = append([]byte(nil), rhs.payload...)\n" +
		"    return nil\n" +
		"}\n" +
		"\n" +
		"/*\n" +
		"// Assign assign internal fields from right hand side packet\n" +
		"func (p *packet{{.MavlinkVersion}}) Assign(rhs Packet) error {\n" +
		"    packet, ok := rhs.(*packet{{.MavlinkVersion}})\n" +
		"    if !ok {\n" +
		"        return fmt.Errorf(\"cast interface '%+v' to '*packet{{.MavlinkVersion}}' fail\", rhs)\n" +
		"    }\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"    p.incompatFlags = rhs.InncompatFlags()\n" +
		"    p.compatFlags = rhs.CompatFlags()\n" +
		"{{- end}}\n" +
		"    p.seqID = p.SeqID()\n" +
		"    p.sysID = p.SysID()\n" +
		"    p.compID = p.CompID()\n" +
		"    p.msgID = p.MsgID()\n" +
		"    p.checksum = p.Checksum()\n" +
		"    p.payload = p.Payload()\n" +
		"    return nil\n" +
		"}\n" +
		"*/\n" +
		"\n" +
		"// Copy returns deep copy of packet\n" +
		"func (p *packet{{.MavlinkVersion}}) Copy() Packet {\n" +
		"    return p.copy()\n" +
		"}\n" +
		"\n" +
		"// Copy returns deep copy of packet\n" +
		"func (p *packet{{.MavlinkVersion}}) copy() *packet{{.MavlinkVersion}} {\n" +
		"    copy := &packet{{.MavlinkVersion}}{}\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"    copy.incompatFlags = p.incompatFlags\n" +
		"    copy.compatFlags = p.compatFlags\n" +
		"{{- end}}\n" +
		"    copy.seqID = p.seqID\n" +
		"    copy.sysID = p.sysID\n" +
		"    copy.compID = p.compID\n" +
		"    copy.msgID = p.msgID\n" +
		"    copy.checksum = p.checksum\n" +
		"    copy.payload = p.payload\n" +
		"    return copy\n" +
		"}\n" +
		"\n" +
		"\n" +
		"/*\n" +
		"// Encode trying to encode message to packet\n" +
		"func (p *packet{{.MavlinkVersion}}) encode(sysID, compID uint8, m Message) error {\n" +
		"\tp.seqID = p.nextSeqNum()\n" +
		"\tp.sysID = sysID\n" +
		"\tp.compID = compID\n" +
		"\treturn p.Encode(m)\n" +
		"}\n" +
		"\n" +
		"// Encode trying to encode message to packet\n" +
		"func (p *packet{{.MavlinkVersion}}) Encode(m Message) error {\n" +
		"\tif err := m.Pack(p); err != nil {\n" +
		"\t\treturn err\n" +
		"\t}\n" +
		"{{- if eq .MavlinkVersion 2 }}\n" +
		"\tpayloadLen := len(p.payload)\n" +
		"\tfor payloadLen > 1 && p.payload[payloadLen-1] == 0 {\n" +
		"\t\tpayloadLen--\n" +
		"\t}\n" +
		"\tp.payload = p.payload[:payloadLen]\n" +
		"{{- end }}\n" +
		"\tif err := p.fixChecksum(); err != nil {\n" +
		"\t\treturn err\n" +
		"\t}\n" +
		"\treturn nil\n" +
		"}\n" +
		"\n" +
		"// Decode trying to decode message to packet\n" +
		"func (p *packet{{.MavlinkVersion}}) Decode(m Message) error {\n" +
		"{{- if eq .MavlinkVersion 2 }}\n" +
		"    if msg, ok := supported[p.msgID]; !ok {\n" +
		"        return ErrUnknownMsgID\n" +
		"    } else if len(p.payload) < msg.Size {\n" +
		"\t\tp.payload = append(p.payload, zeroTail[:msg.Size-len(p.payload)]...)\n" +
		"\t}\n" +
		"{{- end }}\n" +
		"    return m.Unpack(p)\n" +
		"}\n" +
		"*/\n" +
		"\n" +
		"// Unmarshal trying to de-serialize byte slice to packet\n" +
		"func (p *packet{{.MavlinkVersion}}) Unmarshal(buffer []byte) error {\n" +
		"\tparser := _parsersPool_v{{.MavlinkVersion}}.Get().(*parser{{.MavlinkVersion}})\n" +
		"\tdefer parser.Destroy()\n" +
		"\tfor _, c := range buffer {\n" +
		"\t\tpacket, err := parser.parseChar(c)\n" +
		"\t\tif err != nil {\n" +
		"\t\t\treturn err\n" +
		"\t\t}\n" +
		"\t\tif packet != nil {\n" +
		"\t\t\treturn p.assign(packet)\n" +
		"\t\t}\n" +
		"\t}\n" +
		"\treturn ErrNoNewData\n" +
		"}\n" +
		"\n" +
		"// Marshal trying to serialize byte slice from packet\n" +
		"func (p *packet{{.MavlinkVersion}}) Marshal() ([]byte, error) {\n" +
		"\tif p == nil {\n" +
		"\t\treturn nil, ErrNilPointerReference\n" +
		"\t}\n" +
		"    bytes := make([]byte, 0, {{if eq .MavlinkVersion 2 -}} 12 {{- else -}} 8 {{- end}}+len(p.payload))\n" +
		"    // header\n" +
		"    bytes = append(bytes,\n" +
		"\t    {{if eq .MavlinkVersion 2 -}} 0xfd {{- else -}} 0xfe {{- end}},\n" +
		"\t    byte(len(p.payload)),\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"\t    uint8(p.incompatFlags),\n" +
		"\t    uint8(p.compatFlags),\n" +
		"{{- end}}\n" +
		"\t    p.seqID,\n" +
		"\t    p.sysID,\n" +
		"\t    p.compID,\n" +
		"\t    uint8(p.msgID),\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"\t    uint8(p.msgID >> 8),\n" +
		"\t    uint8(p.msgID >> 16),\n" +
		"{{- end}}\n" +
		"    )\n" +
		"    // payload\n" +
		"\tbytes = append(bytes, p.payload...)\n" +
		"\t// crc\n" +
		"\tbytes = append(bytes, p.u16ToBytes(p.checksum)...)\n" +
		"\treturn bytes, nil\n" +
		"}\n" +
		"\n" +
		"func (p *packet{{.MavlinkVersion}}) fixChecksum() error {\n" +
		"    msg, ok := supported[p.msgID]\n" +
		"    if !ok {\n" +
		"\t\treturn ErrUnknownMsgID\n" +
		"    }\n" +
		"\tcrc := NewX25()\n" +
		"\tcrc.WriteByte(byte(len(p.payload)))\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"\tcrc.WriteByte(p.incompatFlags)\n" +
		"\tcrc.WriteByte(p.compatFlags)\n" +
		"{{- end}}\n" +
		"\tcrc.WriteByte(p.seqID)\n" +
		"\tcrc.WriteByte(p.sysID)\n" +
		"\tcrc.WriteByte(p.compID)\n" +
		"\tcrc.WriteByte(byte(p.msgID >> 0 ))\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"\tcrc.WriteByte(byte(p.msgID >> 8 ))\n" +
		"\tcrc.WriteByte(byte(p.msgID >> 16))\n" +
		"{{- end}}\n" +
		"\tcrc.Write(p.payload)\n" +
		"\tcrc.WriteByte(msg.Extra)\n" +
		"\tp.checksum = crc.Sum16()\n" +
		"\treturn nil\n" +
		"}\n" +
		"\n" +
		"func (p *packet{{.MavlinkVersion}}) u16ToBytes(v uint16) []byte {\n" +
		"\treturn []byte{byte(v & 0xff), byte(v >> 8)}\n" +
		"}\n" +
		"\n" +
		"// Message function produce message from packet\n" +
		"func (p *packet{{.MavlinkVersion}}) Message() (Message, error) {\n" +
		"    msg, ok := supported[p.msgID]\n" +
		"\tif !ok {\n" +
		"\t\treturn nil, ErrUnknownMsgID\n" +
		"\t}\n" +
		"\treturn msg.Constructor(p)\n" +
		"}\n" +
		"\n" +
		"// String function return string view of Packet struct\n" +
		"func (p *packet{{.MavlinkVersion}}) String() string {\n" +
		"\treturn fmt.Sprintf(\n" +
		"\t\t\"&packet{{.MavlinkVersion}}{ {{ if eq .MavlinkVersion 2 }}incompatFlags: %08b, compatFlags: %08b, {{ end }}seqID: %d, sysID: %d, compID: %d, msgID: %d, payload: %s, checksum: %d }\",\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"    \tp.incompatFlags,\n" +
		"\t    p.compatFlags,\n" +
		"{{- end}}\n" +
		"\t\tp.seqID,\n" +
		"\t\tp.sysID,\n" +
		"\t\tp.compID,\n" +
		"\t\tint64(p.msgID),\n" +
		"\t\tfunc() string {\n" +
		"\t\t\tmsg, err := p.Message()\n" +
		"\t\t\tif err != nil {\n" +
		"\t\t\t\treturn fmt.Sprintf(\"%0X\", p.payload)\n" +
		"\t\t\t}\n" +
		"\t\t\treturn msg.String()\n" +
		"\t\t}(),\n" +
		"\t\tp.checksum,\n" +
		"\t)\n" +
		"}\n" +
		""
	return tmpl
}
