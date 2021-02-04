/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package parser

import (
	"github.com/asmyasnikov/go-mavlink/mavlink/crc"
	"github.com/asmyasnikov/go-mavlink/mavlink/errors"
	"github.com/asmyasnikov/go-mavlink/mavlink/message"
	"github.com/asmyasnikov/go-mavlink/mavlink/packet"
	"github.com/asmyasnikov/go-mavlink/mavlink/register"
	"sync"
)

// MAVLINK2_PARSE_STATE typedef
type MAVLINK2_PARSE_STATE int

// MAVLINK2_PARSE_STATES
const (
	MAVLINK2_PARSE_STATE_UNINIT             MAVLINK2_PARSE_STATE = iota
	MAVLINK2_PARSE_STATE_IDLE               MAVLINK2_PARSE_STATE = iota
	MAVLINK2_PARSE_STATE_GOT_STX            MAVLINK2_PARSE_STATE = iota
	MAVLINK2_PARSE_STATE_GOT_LENGTH         MAVLINK2_PARSE_STATE = iota
	MAVLINK2_PARSE_STATE_GOT_INCOMPAT_FLAGS MAVLINK2_PARSE_STATE = iota
	MAVLINK2_PARSE_STATE_GOT_COMPAT_FLAGS   MAVLINK2_PARSE_STATE = iota
	MAVLINK2_PARSE_STATE_GOT_SEQ            MAVLINK2_PARSE_STATE = iota
	MAVLINK2_PARSE_STATE_GOT_SYSID          MAVLINK2_PARSE_STATE = iota
	MAVLINK2_PARSE_STATE_GOT_COMPID         MAVLINK2_PARSE_STATE = iota
	MAVLINK2_PARSE_STATE_GOT_MSGID1         MAVLINK2_PARSE_STATE = iota
	MAVLINK2_PARSE_STATE_GOT_MSGID2         MAVLINK2_PARSE_STATE = iota
	MAVLINK2_PARSE_STATE_GOT_MSGID3         MAVLINK2_PARSE_STATE = iota
	MAVLINK2_PARSE_STATE_GOT_PAYLOAD        MAVLINK2_PARSE_STATE = iota
	MAVLINK2_PARSE_STATE_GOT_CRC1           MAVLINK2_PARSE_STATE = iota
	MAVLINK2_PARSE_STATE_GOT_BAD_CRC        MAVLINK2_PARSE_STATE = iota
	MAVLINK2_PARSE_STATE_GOT_GOOD_MESSAGE   MAVLINK2_PARSE_STATE = iota
	MAVLINK2_PARSE_STATE_WAIT_SIGNATURE     MAVLINK2_PARSE_STATE = iota
)

// parser2 is a state machine which parse bytes to packet.Packet
type parser2 struct {
	packet2
	state MAVLINK2_PARSE_STATE
	crc   *crc.X25
}

var _parsersPool_v2 = &sync.Pool{
	New: func() interface{} {
		return new(parser2)
	},
}

// Reset set parser to idle state
func NewParserV2() Parser {
	return _parsersPool_v2.Get().(Parser)
}

// Reset set parser to idle state
func (p *parser2) Destroy() {
	p.state = MAVLINK2_PARSE_STATE_UNINIT
	if p.crc != nil {
		p.crc = nil
	}
	_parsersPool_v2.Put(p)
}

// ParseChar parse char to packet
func (p *parser2) ParseChar(c byte) (packet.Packet, error) {
	return p.parseChar(c)
}

func (p *parser2) parseChar(c byte) (*packet2, error) {
	switch p.state {
	case MAVLINK2_PARSE_STATE_UNINIT,
		MAVLINK2_PARSE_STATE_IDLE,
		MAVLINK2_PARSE_STATE_GOT_BAD_CRC,
		MAVLINK2_PARSE_STATE_GOT_GOOD_MESSAGE:
		if c == byte(MAGIC_NUMBER_V2) {
			p.crc = crc.NewX25()
			p.state = MAVLINK2_PARSE_STATE_GOT_STX
		}
	case MAVLINK2_PARSE_STATE_GOT_STX:
		p.payload = make([]byte, 0, c)
		p.crc.WriteByte(c)
		p.state = MAVLINK2_PARSE_STATE_GOT_LENGTH
	case MAVLINK2_PARSE_STATE_GOT_LENGTH:
		p.incompatFlags = c
		p.crc.WriteByte(c)
		p.state = MAVLINK2_PARSE_STATE_GOT_INCOMPAT_FLAGS
	case MAVLINK2_PARSE_STATE_GOT_INCOMPAT_FLAGS:
		p.compatFlags = c
		p.crc.WriteByte(c)
		p.state = MAVLINK2_PARSE_STATE_GOT_COMPAT_FLAGS
	case MAVLINK2_PARSE_STATE_GOT_COMPAT_FLAGS:
		p.seqID = c
		p.crc.WriteByte(c)
		p.state = MAVLINK2_PARSE_STATE_GOT_SEQ
	case MAVLINK2_PARSE_STATE_GOT_SEQ:
		p.sysID = c
		p.crc.WriteByte(c)
		p.state = MAVLINK2_PARSE_STATE_GOT_SYSID
	case MAVLINK2_PARSE_STATE_GOT_SYSID:
		p.compID = c
		p.crc.WriteByte(c)
		p.state = MAVLINK2_PARSE_STATE_GOT_COMPID
	case MAVLINK2_PARSE_STATE_GOT_COMPID:
		p.msgID = message.MessageID(c)
		p.crc.WriteByte(c)
		p.state = MAVLINK2_PARSE_STATE_GOT_MSGID1
	case MAVLINK2_PARSE_STATE_GOT_MSGID1:
		p.msgID += message.MessageID(c) << 8
		p.crc.WriteByte(c)
		p.state = MAVLINK2_PARSE_STATE_GOT_MSGID2
	case MAVLINK2_PARSE_STATE_GOT_MSGID2:
		p.msgID += message.MessageID(c) << 8 * 2
		p.crc.WriteByte(c)
		p.state = MAVLINK2_PARSE_STATE_GOT_MSGID3
	case MAVLINK2_PARSE_STATE_GOT_MSGID3:
		p.payload = append(p.payload, c)
		p.crc.WriteByte(c)
		if len(p.payload) == cap(p.payload) {
			p.state = MAVLINK2_PARSE_STATE_GOT_PAYLOAD
		}
	case MAVLINK2_PARSE_STATE_GOT_PAYLOAD:
		if info, err := register.Info(p.msgID); err != nil {
			return nil, err
		} else {
			p.crc.WriteByte(info.Extra)
		}
		if c != uint8(p.crc.Sum16()&0xFF) {
			p.state = MAVLINK2_PARSE_STATE_GOT_BAD_CRC
			return nil, errors.ErrCrcFail
		}
		p.state = MAVLINK2_PARSE_STATE_GOT_CRC1
	case MAVLINK2_PARSE_STATE_GOT_CRC1:
		if c == uint8(p.crc.Sum16()>>8) {
			p.checksum = p.crc.Sum16()
			if !p.IsSigned() {
				p.state = MAVLINK2_PARSE_STATE_GOT_GOOD_MESSAGE
				return p.copy(), nil
			}
			p.state = MAVLINK2_PARSE_STATE_WAIT_SIGNATURE
		} else {
			p.state = MAVLINK2_PARSE_STATE_GOT_BAD_CRC
			return nil, errors.ErrCrcFail
		}
	case MAVLINK2_PARSE_STATE_WAIT_SIGNATURE:
		p.signature = append(p.signature, c)
		if len([]byte(p.signature)) == SIGNATURE_LEN {
			p.state = MAVLINK2_PARSE_STATE_GOT_GOOD_MESSAGE
			return p.copy(), nil
		}
	}
	return nil, nil
}
