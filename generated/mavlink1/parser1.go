/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package mavlink

import "sync"

// MAVLINK1_PARSE_STATE typedef
type MAVLINK1_PARSE_STATE int

// MAVLINK1_PARSE_STATES
const (
	MAVLINK1_PARSE_STATE_UNINIT           MAVLINK1_PARSE_STATE = iota
	MAVLINK1_PARSE_STATE_IDLE             MAVLINK1_PARSE_STATE = iota
	MAVLINK1_PARSE_STATE_GOT_STX          MAVLINK1_PARSE_STATE = iota
	MAVLINK1_PARSE_STATE_GOT_LENGTH       MAVLINK1_PARSE_STATE = iota
	MAVLINK1_PARSE_STATE_GOT_SEQ          MAVLINK1_PARSE_STATE = iota
	MAVLINK1_PARSE_STATE_GOT_SYSID        MAVLINK1_PARSE_STATE = iota
	MAVLINK1_PARSE_STATE_GOT_COMPID       MAVLINK1_PARSE_STATE = iota
	MAVLINK1_PARSE_STATE_GOT_MSGID1       MAVLINK1_PARSE_STATE = iota
	MAVLINK1_PARSE_STATE_GOT_PAYLOAD      MAVLINK1_PARSE_STATE = iota
	MAVLINK1_PARSE_STATE_GOT_CRC1         MAVLINK1_PARSE_STATE = iota
	MAVLINK1_PARSE_STATE_GOT_BAD_CRC      MAVLINK1_PARSE_STATE = iota
	MAVLINK1_PARSE_STATE_GOT_GOOD_MESSAGE MAVLINK1_PARSE_STATE = iota
)

// parser1 is a state machine which parse bytes to mavlink.Packet
type parser1 struct {
	packet1
	state MAVLINK1_PARSE_STATE
	crc   *X25
}

var _parsersPool_v1 = &sync.Pool{
	New: func() interface{} {
		return new(parser1)
	},
}

// Reset set parser to idle state
func NewParserV1() Parser {
	return _parsersPool_v1.Get().(Parser)
}

// Reset set parser to idle state
func (p *parser1) Destroy() {
	p.state = MAVLINK1_PARSE_STATE_UNINIT
	if p.crc != nil {
		p.crc.Reset()
		p.crc = nil
	}
	_parsersPool_v1.Put(p)
}

func (p *parser1) ParseChar(c byte) (Packet, error) {
	switch p.state {
	case MAVLINK1_PARSE_STATE_UNINIT,
		MAVLINK1_PARSE_STATE_IDLE,
		MAVLINK1_PARSE_STATE_GOT_BAD_CRC,
		MAVLINK1_PARSE_STATE_GOT_GOOD_MESSAGE:
		if c == 0xfe {
			p.crc = NewX25()
			p.state = MAVLINK1_PARSE_STATE_GOT_STX
		}
	case MAVLINK1_PARSE_STATE_GOT_STX:
		p.payload = make([]byte, 0, c)
		p.crc.WriteByte(c)
		p.state = MAVLINK1_PARSE_STATE_GOT_LENGTH
	case MAVLINK1_PARSE_STATE_GOT_LENGTH:
		p.seqID = c
		p.crc.WriteByte(c)
		p.state = MAVLINK1_PARSE_STATE_GOT_SEQ
	case MAVLINK1_PARSE_STATE_GOT_SEQ:
		p.sysID = c
		p.crc.WriteByte(c)
		p.state = MAVLINK1_PARSE_STATE_GOT_SYSID
	case MAVLINK1_PARSE_STATE_GOT_SYSID:
		p.compID = c
		p.crc.WriteByte(c)
		p.state = MAVLINK1_PARSE_STATE_GOT_COMPID
	case MAVLINK1_PARSE_STATE_GOT_COMPID:
		p.msgID = MessageID(c)
		p.crc.WriteByte(c)
		p.state = MAVLINK1_PARSE_STATE_GOT_MSGID1
	case MAVLINK1_PARSE_STATE_GOT_MSGID1:
		p.payload = append(p.payload, c)
		p.crc.WriteByte(c)
		if len(p.payload) == cap(p.payload) {
			p.state = MAVLINK1_PARSE_STATE_GOT_PAYLOAD
		}
	case MAVLINK1_PARSE_STATE_GOT_PAYLOAD:
		if msg, ok := supported[p.msgID]; !ok {
			return nil, ErrUnknownMsgID
		} else {
			p.crc.WriteByte(msg.Extra)
		}
		if c != uint8(p.crc.Sum16()&0xFF) {
			p.state = MAVLINK1_PARSE_STATE_GOT_BAD_CRC
			return nil, ErrCrcFail
		}
		p.state = MAVLINK1_PARSE_STATE_GOT_CRC1
	case MAVLINK1_PARSE_STATE_GOT_CRC1:
		if c == uint8(p.crc.Sum16()>>8) {
			p.checksum = p.crc.Sum16()
			p.state = MAVLINK1_PARSE_STATE_GOT_GOOD_MESSAGE
			result := packet1{}
			result.Assign(p)
			return &result, nil
		}
		p.state = MAVLINK1_PARSE_STATE_GOT_BAD_CRC
		return nil, ErrCrcFail
	}
	return nil, nil
}
