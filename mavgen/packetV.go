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
		"    \"{{.CommonPackageURL}}/packet\"\n" +
		"    \"{{.CommonPackageURL}}/register\"\n" +
		"    \"{{.CommonPackageURL}}/errors\"\n" +
		"    \"{{.CommonPackageURL}}/helpers\"\n" +
		"    \"{{.CommonPackageURL}}/message\"\n" +
		"    \"{{.CommonPackageURL}}/signature\"\n" +
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
		"\tsignature     signature.Signature\n" +
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
		"func (p *packet{{.MavlinkVersion}}) Signature() signature.Signature {\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"    return append([]byte(nil), p.signature...)\n" +
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
		"    p.signature = append([]byte(nil), []byte(rhs.signature)...)\n" +
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
		"    copy.payload = append([]byte(nil), p.payload...)\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"    copy.signature = append([]byte(nil), []byte(p.signature)...)\n" +
		"{{- end}}\n" +
		"    return copy\n" +
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
		"{{- if eq .MavlinkVersion 2}}\n" +
		"\tif p.IsSigned() {\n" +
		"\t    bytes = append(bytes, []byte(p.signature)...)\n" +
		"\t}\n" +
		"{{- end}}\n" +
		"\treturn bytes, nil\n" +
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
		"\treturn info.Constructor(p)\n" +
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
		"                return fmt.Sprintf(\", signature: %+v\", p.signature)\n" +
		"            }\n" +
		"            return \"\"\n" +
		"        }(),\n" +
		"{{- else}}\n" +
		"        \"\",\n" +
		"{{- end}}\n" +
		"\t)\n" +
		"}\n" +
		"\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"func (p *packet{{.MavlinkVersion}}) GenSignature(key *V2Key) *V2Signature {\n" +
		"\tmsg := f.GetMessage().(*msg.MessageRaw)\n" +
		"\th := sha256.New()\n" +
		"\n" +
		"\t// secret key\n" +
		"\th.Write(key[:])\n" +
		"\n" +
		"\t// the signature covers the whole message, excluding the signature itself\n" +
		"\tbuf := make([]byte, 6)\n" +
		"\th.Write([]byte{V2MagicByte})\n" +
		"\th.Write([]byte{byte(len(msg.Content))})\n" +
		"\th.Write([]byte{f.IncompatibilityFlag})\n" +
		"\th.Write([]byte{f.CompatibilityFlag})\n" +
		"\th.Write([]byte{f.SequenceID})\n" +
		"\th.Write([]byte{f.SystemID})\n" +
		"\th.Write([]byte{f.ComponentID})\n" +
		"\th.Write(uint24Encode(buf, f.Message.GetID()))\n" +
		"\th.Write(msg.Content)\n" +
		"\tbinary.LittleEndian.PutUint16(buf, f.Checksum)\n" +
		"\th.Write(buf[:2])\n" +
		"\th.Write([]byte{f.SignatureLinkID})\n" +
		"\th.Write(uint48Encode(buf, f.SignatureTimestamp))\n" +
		"\n" +
		"\tsig := new(V2Signature)\n" +
		"\tcopy(sig[:], h.Sum(nil)[:6])\n" +
		"\treturn sig\n" +
		"}\n" +
		"{{- end}}\n" +
		""
	return tmpl
}
