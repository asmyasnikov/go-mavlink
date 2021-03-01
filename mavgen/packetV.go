/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

// packetVTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func packetVTemplate() string {
	var tmpl = "package parser\n" +
		"\n" +
		"import (\n" +
		"    \"fmt\"\n" +
		"    \"time\"\n" +
		"    \"{{.CommonPackageURL}}/version\"\n" +
		"    \"{{.CommonPackageURL}}/signature\"\n" +
		"    \"{{.CommonPackageURL}}/packet\"\n" +
		"    \"{{.CommonPackageURL}}/register\"\n" +
		"    \"{{.CommonPackageURL}}/errors\"\n" +
		"    \"{{.CommonPackageURL}}/helpers\"\n" +
		"    \"{{.CommonPackageURL}}/message\"\n" +
		"    \"{{.CommonPackageURL}}/crc\"\n" +
		")\n" +
		"\n" +
		"{{if eq .MavlinkVersion 2 -}}\n" +
		"// MAVLINK_IFLAG type\n" +
		"type MAVLINK_IFLAG uint8\n" +
		"{{- end}}\n" +
		"\n" +
		"const (\n" +
		"    // MAGIC_NUMBER_V{{.MavlinkVersion}} const value for common use\n" +
		"    MAGIC_NUMBER_V{{.MavlinkVersion}} packet.MAGIC_NUMBER = {{if eq .MavlinkVersion 2 -}} 0xfd {{- else -}} 0xfe {{- end}}\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"    // MAVLINK_IFLAG_SIGNED constant\n" +
		"    MAVLINK_IFLAG_SIGNED MAVLINK_IFLAG = 0b00000001\n" +
		"{{- end}}\n" +
		")\n" +
		"\n" +
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
		"\tmsgID         message.MessageID // ID of message in payload\n" +
		"\tpayload       []byte\n" +
		"\tchecksum      uint16\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"    signature     *signature.Signature\n" +
		"{{- end}}\n" +
		"}\n" +
		"\n" +
		"// NewPacketV{{.MavlinkVersion}} creates new mavlink{{.MavlinkVersion}}.Packet\n" +
		"func NewPacketV{{.MavlinkVersion}}(sysID uint8, compID uint8, seqID uint8, message message.Message) (packet.Packet, error) {\n" +
		"\tpayload, err := message.Marshal()\n" +
		"\tif err != nil {\n" +
		"\t\treturn nil, err\n" +
		"\t}\n" +
		"    return &packet{{.MavlinkVersion}}{\n" +
		"        seqID:   seqID,\n" +
		"        sysID:   sysID,\n" +
		"        compID:  compID,\n" +
		"        msgID:   message.MsgID(),\n" +
		"        payload: payload,\n" +
		"    }, nil\n" +
		"}\n" +
		"\n" +
		"// Version returns version of packet framing\n" +
		"func (p *packet{{.MavlinkVersion}}) Version() version.MAVLINK_VERSION {\n" +
		"    return version.MAVLINK_V{{.MavlinkVersion}}\n" +
		"}\n" +
		"\n" +
		"// IsNil returns true if packet is nil\n" +
		"func (p *packet{{.MavlinkVersion}}) IsNil() bool {\n" +
		"    return p == nil\n" +
		"}\n" +
		"\n" +
		"// IsSigned checks whether the frame contains a signature. It does not validate the signature\n" +
		"func (p *packet{{.MavlinkVersion}}) IsSigned() bool {\n" +
		"\treturn {{if eq .MavlinkVersion 2 -}} (MAVLINK_IFLAG(p.incompatFlags) & MAVLINK_IFLAG_SIGNED) != 0 {{- else -}} false {{- end}}\n" +
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
		"func (p *packet{{.MavlinkVersion}}) MsgID() message.MessageID {\n" +
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
		"// Signature returns packet signature\n" +
		"func (p *packet{{.MavlinkVersion}}) Signature() *signature.Signature {\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"    return p.signature.Copy()\n" +
		"{{- else}}\n" +
		"    return nil\n" +
		"{{- end}}\n" +
		"}\n" +
		"\n" +
		"func (p *packet{{.MavlinkVersion}}) assign(rhs *packet{{.MavlinkVersion}}) error {\n" +
		"    if p == nil {\n" +
		"        return errors.ErrNilPointerReference\n" +
		"    }\n" +
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
		"{{- if eq .MavlinkVersion 2}}\n" +
		"    p.signature = rhs.signature.Copy()\n" +
		"{{- end}}\n" +
		"    return nil\n" +
		"}\n" +
		"\n" +
		"// Copy returns deep copy of packet\n" +
		"func (p *packet{{.MavlinkVersion}}) Copy() packet.Packet {\n" +
		"    return p.copy()\n" +
		"}\n" +
		"\n" +
		"// Copy returns deep copy of packet\n" +
		"func (p *packet{{.MavlinkVersion}}) copy() *packet{{.MavlinkVersion}} {\n" +
		"    if p == nil {\n" +
		"        return nil\n" +
		"    }\n" +
		"    c := &packet{{.MavlinkVersion}}{}\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"    c.incompatFlags = p.incompatFlags\n" +
		"    c.compatFlags = p.compatFlags\n" +
		"{{- end}}\n" +
		"    c.seqID = p.seqID\n" +
		"    c.sysID = p.sysID\n" +
		"    c.compID = p.compID\n" +
		"    c.msgID = p.msgID\n" +
		"    c.checksum = p.checksum\n" +
		"    c.payload = append([]byte(nil), p.payload...)\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"    c.signature = p.signature.Copy()\n" +
		"{{- end}}\n" +
		"    return c\n" +
		"}\n" +
		"\n" +
		"// Unmarshal trying to de-serialize byte slice to packet\n" +
		"func (p *packet{{.MavlinkVersion}}) Unmarshal(buffer []byte) error {\n" +
		"    if p == nil {\n" +
		"        return errors.ErrNilPointerReference\n" +
		"    }\n" +
		"\tparser := _parsersPoolV{{.MavlinkVersion}}.Get().(*parser{{.MavlinkVersion}}).reset()\n" +
		"\tdefer parser.Destroy()\n" +
		"\tfor _, c := range buffer {\n" +
		"\t\tpacket, err := parser.parseChar(c)\n" +
		"\t\tif err != nil {\n" +
		"\t\t\treturn fmt.Errorf(\"Unmarshal fail with error \\\"%+v\\\". Parser %+v\", err, parser)\n" +
		"\t\t}\n" +
		"\t\tif packet != nil {\n" +
		"\t\t\treturn p.assign(packet)\n" +
		"\t\t}\n" +
		"\t}\n" +
		"\treturn fmt.Errorf(\"Unmarshal fail with error \\\"%+v\\\". Parser %+v\", errors.ErrNoNewData, parser)\n" +
		"}\n" +
		"\n" +
		"// Marshal trying to serialize byte slice from packet\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"// After Marshal() need to append signature if packet is signed\n" +
		"//\tif p.IsSigned() {\n" +
		"//\t\tif err := p.signature.Fix(bytes, secretKey); err != nil {\n" +
		"//\t\t    return nil, err\n" +
		"//\t\t}\n" +
		"//\t\tb, err := p.signature.Marshal()\n" +
		"//\t\tif err != nil {\n" +
		"//\t\t\treturn nil, err\n" +
		"//\t\t}\n" +
		"//\t\tbytes = append(bytes,b...)\n" +
		"//\t}\n" +
		"{{- end}}\n" +
		"func (p *packet{{.MavlinkVersion}}) Marshal() ([]byte, error) {\n" +
		"\tif p == nil {\n" +
		"\t\treturn nil, errors.ErrNilPointerReference\n" +
		"\t}\n" +
		"    bytes := make([]byte, 0, {{if eq .MavlinkVersion 2 -}} 12 {{- else -}} 8 {{- end}}+len(p.payload))\n" +
		"\t// prepare\n" +
		"\tif err := p.prepare(); err != nil {\n" +
		"\t\treturn nil, err\n" +
		"\t}\n" +
		"    // header\n" +
		"    bytes = append(bytes,\n" +
		"\t    byte(MAGIC_NUMBER_V{{.MavlinkVersion}}),\n" +
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
		"\tbytes = append(bytes, helpers.U16ToBytes(p.checksum)...)\n" +
		"\treturn bytes, nil\n" +
		"}\n" +
		"\n" +
		"// Marshal trying to serialize byte slice from packet including signature at appendix\n" +
		"func (p *packet{{.MavlinkVersion}}) MarshalWithSignature(linkID byte, timestamp time.Time, secretKey [32]byte) ([]byte, error) {\n" +
		"    bytes, err := p.Marshal()\n" +
		"    if err != nil {\n" +
		"        return nil, err\n" +
		"    }\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"    if p.IsSigned() {\n" +
		"        p.signature, err = signature.New(linkID, timestamp, secretKey, bytes)\n" +
		"        if err != nil {\n" +
		"            return nil, err\n" +
		"        }\n" +
		"        b, err := p.signature.Marshal()\n" +
		"        if err != nil {\n" +
		"            return nil, err\n" +
		"        }\n" +
		"        bytes = append(bytes, b...)\n" +
		"    }\n" +
		"{{- end}}\n" +
		"    return bytes, err\n" +
		"}\n" +
		"\n" +
		"func (p *packet{{.MavlinkVersion}}) prepare() error {\n" +
		"    if p == nil {\n" +
		"        return errors.ErrNilPointerReference\n" +
		"    }\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"    // payload minify\n" +
		"\tpayloadLen := len(p.payload)\n" +
		"\tfor payloadLen > 1 && p.payload[payloadLen-1] == 0 {\n" +
		"\t\tpayloadLen--\n" +
		"\t}\n" +
		"\tp.payload = p.payload[:payloadLen]\n" +
		"{{- end}}\n" +
		"    info, err := register.Info(p.msgID)\n" +
		"    if err != nil {\n" +
		"\t\treturn err\n" +
		"    }\n" +
		"\tcrc := crc.NewX25()\n" +
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
		"\tcrc.WriteByte(info.Extra)\n" +
		"\tp.checksum = crc.Sum16()\n" +
		"\treturn nil\n" +
		"}\n" +
		"\n" +
		"// Message function produce message from packet\n" +
		"func (p *packet{{.MavlinkVersion}}) Message() (message.Message, error) {\n" +
		"    if p == nil {\n" +
		"        return nil, errors.ErrNilPointerReference\n" +
		"    }\n" +
		"    info, err := register.Info(p.msgID)\n" +
		"\tif err != nil {\n" +
		"\t\treturn nil, err\n" +
		"\t}\n" +
		"\treturn info.Constructor(p.payload)\n" +
		"}\n" +
		"\n" +
		"// String function return string view of Packet struct\n" +
		"func (p *packet{{.MavlinkVersion}}) String() string {\n" +
		"    if p == nil {\n" +
		"        return \"nil\"\n" +
		"    }\n" +
		"\treturn fmt.Sprintf(\n" +
		"\t\t\"&mavlink{{.MavlinkVersion}}.Packet{ {{ if eq .MavlinkVersion 2 }}incompatFlags: %08b, compatFlags: %08b, {{ end }}seqID: %d, sysID: %d, compID: %d, msgID: %d, payload: %s, checksum: %d%s }\",\n" +
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
		"{{- if eq .MavlinkVersion 2}}\n" +
		"        func() string {\n" +
		"            if p.IsSigned() {\n" +
		"                return fmt.Sprintf(\", signature: %+v\", *p.signature)\n" +
		"            }\n" +
		"            return \"\"\n" +
		"        }(),\n" +
		"{{- else}}\n" +
		"        \"\",\n" +
		"{{- end}}\n" +
		"\t)\n" +
		"}\n" +
		""
	return tmpl
}
