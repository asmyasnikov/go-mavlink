/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package parser

import (
	"fmt"
	"github.com/asmyasnikov/go-mavlink/mavlink/crc"
	"github.com/asmyasnikov/go-mavlink/mavlink/errors"
	"github.com/asmyasnikov/go-mavlink/mavlink/message"
	"github.com/asmyasnikov/go-mavlink/mavlink/packet"
	"github.com/asmyasnikov/go-mavlink/mavlink/register"
	"sync"
)

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

func (s MAVLINK1_PARSE_STATE) String() string {
	switch s {
	case MAVLINK1_PARSE_STATE_UNINIT:
		return "MAVLINK1_PARSE_STATE_UNINIT"
	case MAVLINK1_PARSE_STATE_IDLE:
		return "MAVLINK1_PARSE_STATE_IDLE"
	case MAVLINK1_PARSE_STATE_GOT_STX:
		return "MAVLINK1_PARSE_STATE_GOT_STX"
	case MAVLINK1_PARSE_STATE_GOT_LENGTH:
		return "MAVLINK1_PARSE_STATE_GOT_LENGTH"
	case MAVLINK1_PARSE_STATE_GOT_SEQ:
		return "MAVLINK1_PARSE_STATE_GOT_SEQ"
	case MAVLINK1_PARSE_STATE_GOT_SYSID:
		return "MAVLINK1_PARSE_STATE_GOT_SYSID"
	case MAVLINK1_PARSE_STATE_GOT_COMPID:
		return "MAVLINK1_PARSE_STATE_GOT_COMPID"
	case MAVLINK1_PARSE_STATE_GOT_MSGID1:
		return "MAVLINK1_PARSE_STATE_GOT_MSGID1"
	case MAVLINK1_PARSE_STATE_GOT_PAYLOAD:
		return "MAVLINK1_PARSE_STATE_GOT_PAYLOAD"
	case MAVLINK1_PARSE_STATE_GOT_CRC1:
		return "MAVLINK1_PARSE_STATE_GOT_CRC1"
	case MAVLINK1_PARSE_STATE_GOT_BAD_CRC:
		return "MAVLINK1_PARSE_STATE_GOT_BAD_CRC"
	case MAVLINK1_PARSE_STATE_GOT_GOOD_MESSAGE:
		return "MAVLINK1_PARSE_STATE_GOT_GOOD_MESSAGE"
	default:
		return fmt.Sprintf("MAVLINK1_PARSE_STATE_UNKNOWN%d", s)
	}
}

// parser1 is a state machine which parse bytes to packet.Packet
type parser1 struct {
	packet1
	state MAVLINK1_PARSE_STATE
	crc   *crc.X25
	tail  []byte
}

var _parsersPoolV1 = &sync.Pool{
	New: func() interface{} {
		return new(parser1)
	},
}

// NewParserV1 returns Parser from inner pool
func NewParserV1() Parser {
	return _parsersPoolV1.Get().(Parser)
}

// String returns string representation
func (p *parser1) String() string {
	return fmt.Sprintf("mavlink1.Parser{ state: %+v, packet: %+v }", p.state, p.packet1)
}

// Destroy set parser to idle state and return it into inner pool
func (p *parser1) Destroy() {
	_parsersPoolV1.Put(p.reset())
}

func (p *parser1) reset() *parser1 {
	p.state = MAVLINK1_PARSE_STATE_UNINIT
	if p.crc != nil {
		p.crc = nil
	}
	return p
}

// ParseChar parse char to packet
func (p *parser1) ParseChar(c byte) (packet.Packet, error) {
	return p.parseChar(c)
}

func (p *parser1) parseChar(c byte) (*packet1, error) {
	switch p.state {
	case MAVLINK1_PARSE_STATE_UNINIT,
		MAVLINK1_PARSE_STATE_IDLE,
		MAVLINK1_PARSE_STATE_GOT_BAD_CRC,
		MAVLINK1_PARSE_STATE_GOT_GOOD_MESSAGE:
		if c == byte(MAGIC_NUMBER_V1) {
			p.crc = crc.NewX25()
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
		p.msgID = message.MessageID(c)
		p.crc.WriteByte(c)
		p.state = MAVLINK1_PARSE_STATE_GOT_MSGID1
	case MAVLINK1_PARSE_STATE_GOT_MSGID1:
		p.payload = append(p.payload, c)
		p.crc.WriteByte(c)
		if len(p.payload) == cap(p.payload) {
			p.state = MAVLINK1_PARSE_STATE_GOT_PAYLOAD
		}
	case MAVLINK1_PARSE_STATE_GOT_PAYLOAD:
		info, err := register.Info(p.msgID)
		if err != nil {
			return nil, err
		}
		p.crc.WriteByte(info.Extra)
		if c != uint8(p.crc.Sum16()&0xFF) {
			p.state = MAVLINK1_PARSE_STATE_GOT_BAD_CRC
			return nil, errors.ErrCrcFail
		}
		p.state = MAVLINK1_PARSE_STATE_GOT_CRC1
	case MAVLINK1_PARSE_STATE_GOT_CRC1:
		if c == uint8(p.crc.Sum16()>>8) {
			p.checksum = p.crc.Sum16()
			p.state = MAVLINK1_PARSE_STATE_GOT_GOOD_MESSAGE
			return p.copy(), nil
		}
		p.state = MAVLINK1_PARSE_STATE_GOT_BAD_CRC
		return nil, errors.ErrCrcFail
	}
	return nil, nil
}
