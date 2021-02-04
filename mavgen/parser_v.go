/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

// parser_vTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func parser_vTemplate() string {
	var tmpl = "package parser\n" +
		"\n" +
		"import (\n" +
		"    \"sync\"\n" +
		"    \"{{.CommonPackageURL}}/packet\"\n" +
		"    \"{{.CommonPackageURL}}/register\"\n" +
		"    \"{{.CommonPackageURL}}/errors\"\n" +
		"    \"{{.CommonPackageURL}}/message\"\n" +
		"    \"{{.CommonPackageURL}}/crc\"\n" +
		")\n" +
		"\n" +
		"// MAVLINK{{.MavlinkVersion}}_PARSE_STATE typedef\n" +
		"type MAVLINK{{.MavlinkVersion}}_PARSE_STATE int\n" +
		"\n" +
		"// MAVLINK{{.MavlinkVersion}}_PARSE_STATES\n" +
		"const (\n" +
		"\tMAVLINK{{.MavlinkVersion}}_PARSE_STATE_UNINIT             MAVLINK{{.MavlinkVersion}}_PARSE_STATE = iota\n" +
		"\tMAVLINK{{.MavlinkVersion}}_PARSE_STATE_IDLE               MAVLINK{{.MavlinkVersion}}_PARSE_STATE = iota\n" +
		"\tMAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_STX            MAVLINK{{.MavlinkVersion}}_PARSE_STATE = iota\n" +
		"\tMAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_LENGTH         MAVLINK{{.MavlinkVersion}}_PARSE_STATE = iota\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"\tMAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_INCOMPAT_FLAGS MAVLINK{{.MavlinkVersion}}_PARSE_STATE = iota\n" +
		"\tMAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_COMPAT_FLAGS   MAVLINK{{.MavlinkVersion}}_PARSE_STATE = iota\n" +
		"{{- end}}\n" +
		"\tMAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_SEQ            MAVLINK{{.MavlinkVersion}}_PARSE_STATE = iota\n" +
		"\tMAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_SYSID          MAVLINK{{.MavlinkVersion}}_PARSE_STATE = iota\n" +
		"\tMAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_COMPID         MAVLINK{{.MavlinkVersion}}_PARSE_STATE = iota\n" +
		"\tMAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_MSGID1         MAVLINK{{.MavlinkVersion}}_PARSE_STATE = iota\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"\tMAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_MSGID2         MAVLINK{{.MavlinkVersion}}_PARSE_STATE = iota\n" +
		"\tMAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_MSGID3         MAVLINK{{.MavlinkVersion}}_PARSE_STATE = iota\n" +
		"{{- end}}\n" +
		"\tMAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_PAYLOAD        MAVLINK{{.MavlinkVersion}}_PARSE_STATE = iota\n" +
		"\tMAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_CRC1           MAVLINK{{.MavlinkVersion}}_PARSE_STATE = iota\n" +
		"\tMAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_BAD_CRC        MAVLINK{{.MavlinkVersion}}_PARSE_STATE = iota\n" +
		"\tMAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_GOOD_MESSAGE   MAVLINK{{.MavlinkVersion}}_PARSE_STATE = iota\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"\tMAVLINK{{.MavlinkVersion}}_PARSE_STATE_WAIT_SIGNATURE     MAVLINK{{.MavlinkVersion}}_PARSE_STATE = iota\n" +
		"{{- end}}\n" +
		")\n" +
		"\n" +
		"// parser{{.MavlinkVersion}} is a state machine which parse bytes to packet.Packet\n" +
		"type parser{{.MavlinkVersion}} struct {\n" +
		"    packet{{.MavlinkVersion}}\n" +
		"\tstate       MAVLINK{{.MavlinkVersion}}_PARSE_STATE\n" +
		"\tcrc         *crc.X25\n" +
		"}\n" +
		"\n" +
		"var _parsersPool_v{{.MavlinkVersion}} = &sync.Pool{\n" +
		"\tNew: func() interface{} {\n" +
		"\t\treturn new(parser{{.MavlinkVersion}})\n" +
		"\t},\n" +
		"}\n" +
		"\n" +
		"// Reset set parser to idle state\n" +
		"func NewParserV{{.MavlinkVersion}}() Parser {\n" +
		"    return _parsersPool_v{{.MavlinkVersion}}.Get().(Parser)\n" +
		"}\n" +
		"\n" +
		"// Reset set parser to idle state\n" +
		"func (p *parser{{.MavlinkVersion}}) Destroy() {\n" +
		"\tp.state = MAVLINK{{.MavlinkVersion}}_PARSE_STATE_UNINIT\n" +
		"\tif p.crc != nil {\n" +
		"\t\tp.crc = nil\n" +
		"\t}\n" +
		"\t_parsersPool_v{{.MavlinkVersion}}.Put(p)\n" +
		"}\n" +
		"\n" +
		"// ParseChar parse char to packet\n" +
		"func (p *parser{{.MavlinkVersion}}) ParseChar(c byte) (packet.Packet, error) {\n" +
		"    return p.parseChar(c)\n" +
		"}\n" +
		"\n" +
		"func (p *parser{{.MavlinkVersion}}) parseChar(c byte) (*packet{{.MavlinkVersion}}, error) {\n" +
		"\tswitch p.state {\n" +
		"\tcase MAVLINK{{.MavlinkVersion}}_PARSE_STATE_UNINIT,\n" +
		"\t\t MAVLINK{{.MavlinkVersion}}_PARSE_STATE_IDLE,\n" +
		"\t\t MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_BAD_CRC,\n" +
		"\t\t MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_GOOD_MESSAGE :\n" +
		"\t\tif c == byte(MAGIC_NUMBER_V{{.MavlinkVersion}}) {\n" +
		"\t\t\tp.crc = crc.NewX25()\n" +
		"\t\t\tp.state = MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_STX\n" +
		"\t\t}\n" +
		"\tcase MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_STX:\n" +
		"\t\tp.payload = make([]byte, 0, c)\n" +
		"\t\tp.crc.WriteByte(c)\n" +
		"\t\tp.state = MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_LENGTH\n" +
		"\tcase MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_LENGTH:\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"\t\tp.incompatFlags = c\n" +
		"\t\tp.crc.WriteByte(c)\n" +
		"\t\tp.state = MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_INCOMPAT_FLAGS\n" +
		"\tcase MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_INCOMPAT_FLAGS:\n" +
		"\t\tp.compatFlags = c\n" +
		"\t\tp.crc.WriteByte(c)\n" +
		"\t\tp.state = MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_COMPAT_FLAGS\n" +
		"\tcase MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_COMPAT_FLAGS:\n" +
		"{{- end}}\n" +
		"\t\tp.seqID = c\n" +
		"\t\tp.crc.WriteByte(c)\n" +
		"\t\tp.state = MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_SEQ\n" +
		"\tcase MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_SEQ:\n" +
		"\t\tp.sysID = c\n" +
		"\t\tp.crc.WriteByte(c)\n" +
		"\t\tp.state = MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_SYSID\n" +
		"\tcase MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_SYSID:\n" +
		"\t\tp.compID = c\n" +
		"\t\tp.crc.WriteByte(c)\n" +
		"\t\tp.state = MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_COMPID\n" +
		"\tcase MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_COMPID:\n" +
		"\t\tp.msgID = message.MessageID(c)\n" +
		"\t\tp.crc.WriteByte(c)\n" +
		"\t\tp.state = MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_MSGID1\n" +
		"\tcase MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_MSGID1:\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"\t\tp.msgID += message.MessageID(c) << 8\n" +
		"\t\tp.crc.WriteByte(c)\n" +
		"\t\tp.state = MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_MSGID2\n" +
		"\tcase MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_MSGID2:\n" +
		"\t\tp.msgID += message.MessageID(c) << 8 * 2\n" +
		"\t\tp.crc.WriteByte(c)\n" +
		"\t\tp.state = MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_MSGID3\n" +
		"\tcase MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_MSGID3:\n" +
		"{{- end}}\n" +
		"\t\tp.payload = append(p.payload, c)\n" +
		"\t\tp.crc.WriteByte(c)\n" +
		"\t\tif len(p.payload) == cap(p.payload) {\n" +
		"\t\t\tp.state = MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_PAYLOAD\n" +
		"\t\t}\n" +
		"\tcase MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_PAYLOAD:\n" +
		"        if info, err := register.Info(p.msgID); err != nil {\n" +
		"            return nil, err\n" +
		"        } else {\n" +
		"            p.crc.WriteByte(info.Extra)\n" +
		"        }\n" +
		"\t\tif c != uint8(p.crc.Sum16()&0xFF) {\n" +
		"\t\t\tp.state = MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_BAD_CRC\n" +
		"\t\t\treturn nil, errors.ErrCrcFail\n" +
		"\t\t}\n" +
		"        p.state = MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_CRC1\n" +
		"\tcase MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_CRC1:\n" +
		"\t\tif c == uint8(p.crc.Sum16()>>8) {\n" +
		"\t\t\tp.checksum = p.crc.Sum16()\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"            if !p.IsSigned() {\n" +
		"                p.state = MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_GOOD_MESSAGE\n" +
		"                return p.copy(), nil\n" +
		"            }\n" +
		"            p.state = MAVLINK{{.MavlinkVersion}}_PARSE_STATE_WAIT_SIGNATURE\n" +
		"{{- else}}\n" +
		"\t\t\tp.state = MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_GOOD_MESSAGE\n" +
		"\t\t\treturn p.copy(), nil\n" +
		"{{- end}}\n" +
		"\t\t} else {\n" +
		"            p.state = MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_BAD_CRC\n" +
		"            return nil, errors.ErrCrcFail\n" +
		"        }\n" +
		"{{- if eq .MavlinkVersion 2}}\n" +
		"\tcase MAVLINK{{.MavlinkVersion}}_PARSE_STATE_WAIT_SIGNATURE:\n" +
		"\t    p.signature = append(p.signature, c)\n" +
		"\t    if len([]byte(p.signature)) == SIGNATURE_LEN {\n" +
		"\t\t\tp.state = MAVLINK{{.MavlinkVersion}}_PARSE_STATE_GOT_GOOD_MESSAGE\n" +
		"\t\t\treturn p.copy(), nil\n" +
		"\t    }\n" +
		"{{- end}}\n" +
		"\t}\n" +
		"\treturn nil, nil\n" +
		"}\n" +
		"\n" +
		""
	return tmpl
}
