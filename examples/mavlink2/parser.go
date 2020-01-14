/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package mavlink

import "sync"

// Parser is a finally state machine which parse bytes to mavlink.Packet
type Parser struct {
	state  MAVLINK_PARSE_STATE
	packet Packet
	crc    *X25
}

var parsersPool = &sync.Pool{
	New: func() interface{} {
		return new(Parser)
	},
}

// Reset set parser to idle state
func (p *Parser) Reset() {
	p.state = MAVLINK_PARSE_STATE_UNINIT
	p.crc.Reset()
	p.crc = nil
}

func (p *Parser) parseChar(c byte) (*Packet, error) {
	switch p.state {
	case MAVLINK_PARSE_STATE_UNINIT,
		MAVLINK_PARSE_STATE_IDLE,
		MAVLINK_PARSE_STATE_GOT_BAD_CRC,
		MAVLINK_PARSE_STATE_GOT_GOOD_MESSAGE:
		if c == magicNumber {
			p.crc = NewX25()
			p.state = MAVLINK_PARSE_STATE_GOT_STX
		}
	case MAVLINK_PARSE_STATE_GOT_STX:
		p.packet.Payload = make([]byte, 0, c)
		p.crc.WriteByte(c)
		p.state = MAVLINK_PARSE_STATE_GOT_LENGTH
	case MAVLINK_PARSE_STATE_GOT_LENGTH:
		p.packet.InCompatFlags = c
		p.crc.WriteByte(c)
		p.state = MAVLINK_PARSE_STATE_GOT_INCOMPAT_FLAGS
	case MAVLINK_PARSE_STATE_GOT_INCOMPAT_FLAGS:
		p.packet.CompatFlags = c
		p.crc.WriteByte(c)
		p.state = MAVLINK_PARSE_STATE_GOT_COMPAT_FLAGS
	case MAVLINK_PARSE_STATE_GOT_COMPAT_FLAGS:
		p.packet.SeqID = c
		p.crc.WriteByte(c)
		p.state = MAVLINK_PARSE_STATE_GOT_SEQ
	case MAVLINK_PARSE_STATE_GOT_SEQ:
		p.packet.SysID = c
		p.crc.WriteByte(c)
		p.state = MAVLINK_PARSE_STATE_GOT_SYSID
	case MAVLINK_PARSE_STATE_GOT_SYSID:
		p.packet.CompID = c
		p.crc.WriteByte(c)
		p.state = MAVLINK_PARSE_STATE_GOT_COMPID
	case MAVLINK_PARSE_STATE_GOT_COMPID:
		p.packet.MsgID = MessageID(c)
		p.crc.WriteByte(c)
		p.state = MAVLINK_PARSE_STATE_GOT_MSGID1
	case MAVLINK_PARSE_STATE_GOT_MSGID1:
		p.packet.MsgID += MessageID(c) << 8
		p.crc.WriteByte(c)
		p.state = MAVLINK_PARSE_STATE_GOT_MSGID2
	case MAVLINK_PARSE_STATE_GOT_MSGID2:
		p.packet.MsgID += MessageID(c) << 8 * 2
		p.crc.WriteByte(c)
		p.state = MAVLINK_PARSE_STATE_GOT_MSGID3
	case MAVLINK_PARSE_STATE_GOT_MSGID3:
		p.packet.Payload = append(p.packet.Payload, c)
		p.crc.WriteByte(c)
		if len(p.packet.Payload) == cap(p.packet.Payload) {
			p.state = MAVLINK_PARSE_STATE_GOT_PAYLOAD
		}
	case MAVLINK_PARSE_STATE_GOT_PAYLOAD:
		crcExtra, err := dialects.findCrcX(p.packet.MsgID)
		if err != nil {
			crcExtra = 0
		}
		p.crc.WriteByte(crcExtra)
		if c != uint8(p.crc.Sum16()&0xFF) {
			p.state = MAVLINK_PARSE_STATE_GOT_BAD_CRC
			p.packet = Packet{}
			return nil, ErrCrcFail
		}
		p.state = MAVLINK_PARSE_STATE_GOT_CRC1
	case MAVLINK_PARSE_STATE_GOT_CRC1:
		if c == uint8(p.crc.Sum16()>>8) {
			p.packet.Checksum = p.crc.Sum16()
			p.state = MAVLINK_PARSE_STATE_GOT_GOOD_MESSAGE
			result := p.packet
			p.packet = Packet{}
			return &result, nil
		}
		p.state = MAVLINK_PARSE_STATE_GOT_BAD_CRC
		p.packet = Packet{}
		return nil, ErrCrcFail
	}
	return nil, nil
}
