// Code generated by go-mavlink/mavgen
// DO NOT EDIT!

// Code generated by go-mavlink/mavgen
// DO NOT EDIT!

package mavlink

import (
	"errors"
	"github.com/asmyasnikov/go-mavlink/x25"
	"sync"
	"time"
)

const (
	magicNumber = 0xfd
)

type MAVLINK_PARSE_STATE int

// MAVLINK_PARSE_STATES
const (
	MAVLINK_PARSE_STATE_UNINIT             MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_IDLE               MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_STX            MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_LENGTH         MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_INCOMPAT_FLAGS MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_COMPAT_FLAGS   MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_SEQ            MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_SYSID          MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_COMPID         MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_MSGID1         MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_MSGID2         MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_MSGID3         MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_PAYLOAD        MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_CRC1           MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_BAD_CRC        MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_GOOD_MESSAGE   MAVLINK_PARSE_STATE = iota
)

var (
	// ErrUnknownMsgID define
	ErrUnknownMsgID = errors.New("unknown msg id")
	// ErrCrcFail define
	ErrCrcFail = errors.New("checksum did not match")
	// ErrNoNewData define
	ErrNoNewData = errors.New("No new data")
	// currentSeqNum
	currentSeqNum uint8 = 0
	// dialects
	dialects DialectSlice = DialectSlice{DialectCommon}
)

func AddDialect(d *Dialect) {
	dialects.Add(d)
}

func RemoveDialect(d *Dialect) {
	dialects.Remove(d)
}

// MessageID typedef
type MessageID uint32

// Message is a basic type for encoding/decoding mavlink messages.
// use the Pack() and Unpack() routines on specific message
// types to convert them to/from the Packet type.
type Message interface {
	Pack(*Packet) error
	Unpack(*Packet) error
	MsgID() MessageID
	MsgName() string
}

// Packet is a wire type for encoding/decoding mavlink messages.
// use the ToPacket() and FromPacket() routines on specific message
// types to convert them to/from the Message type.
type Packet struct {
	InCompatFlags uint8     // incompat flags
	CompatFlags   uint8     // compat flags
	SeqID         uint8     // Sequence of packet
	SysID         uint8     // ID of message sender system/aircraft
	CompID        uint8     // ID of the message sender component
	MsgID         MessageID // ID of message in payload
	Payload       []byte
	Checksum      uint16
}

func (p *Packet) nextSeqNum() byte {
	currentSeqNum++
	return currentSeqNum
}

func (p *Packet) Encode(sysID, compID uint8, m Message) error {
	p.SeqID = p.nextSeqNum()
	p.SysID = sysID
	p.CompID = compID
	if err := m.Pack(p); err != nil {
		return err
	}
	if err := p.fixChecksum(dialects); err != nil {
		return err
	}
	return nil
}

func (p *Packet) Bytes() []byte {
	bytes := []byte{
		magicNumber,
		byte(len(p.Payload)),
		uint8(p.InCompatFlags),
		uint8(p.CompatFlags),
		p.SeqID,
		p.SysID,
		p.CompID,
		uint8(p.MsgID),
		uint8(p.MsgID >> 8),
		uint8(p.MsgID >> 16),
	} // header
	bytes = append(bytes, p.Payload...)
	bytes = append(bytes, u16ToBytes(p.Checksum)...)
	return bytes
}

type Multicast struct {
	sync.Mutex
	listeners map[chan []byte](chan bool)
}

func NewMulticast() Multicast {
	m := Multicast{
		listeners: make(map[chan []byte](chan bool)),
	}
	return m
}

// Decoder struct provide decoding processor
type Decoder struct {
	Multicast Multicast
	data      chan []byte
	decoded   chan *Packet
}

func (m *Multicast) register(done chan bool) chan []byte {
	data := make(chan []byte)
	m.Lock()
	m.listeners[data] = done
	m.Unlock()
	return data
}

func (m *Multicast) notify(buffer []byte) {
	m.Lock()
	for k, _ := range m.listeners {
		k <- buffer
	}
	m.Unlock()
}

func (m *Multicast) clear(data chan []byte) {
	m.Lock()
	delete(m.listeners, data)
	m.Unlock()
}

type Parser struct {
	state  MAVLINK_PARSE_STATE
	packet Packet
	crc    *x25.X25
}

func (parser *Parser) parseChar(c byte) (*Packet, error) {
	switch parser.state {
	case MAVLINK_PARSE_STATE_UNINIT,
		MAVLINK_PARSE_STATE_IDLE,
		MAVLINK_PARSE_STATE_GOT_BAD_CRC,
		MAVLINK_PARSE_STATE_GOT_GOOD_MESSAGE:
		if c == magicNumber {
			parser.crc = x25.New()
			parser.state = MAVLINK_PARSE_STATE_GOT_STX
		}
	case MAVLINK_PARSE_STATE_GOT_STX:
		parser.packet.Payload = make([]byte, 0, c)
		parser.crc.WriteByte(c)
		parser.state = MAVLINK_PARSE_STATE_GOT_LENGTH
	case MAVLINK_PARSE_STATE_GOT_LENGTH:
		parser.packet.InCompatFlags = c
		parser.crc.WriteByte(c)
		parser.state = MAVLINK_PARSE_STATE_GOT_INCOMPAT_FLAGS
	case MAVLINK_PARSE_STATE_GOT_INCOMPAT_FLAGS:
		parser.packet.CompatFlags = c
		parser.crc.WriteByte(c)
		parser.state = MAVLINK_PARSE_STATE_GOT_COMPAT_FLAGS
	case MAVLINK_PARSE_STATE_GOT_COMPAT_FLAGS:
		parser.packet.SeqID = c
		parser.crc.WriteByte(c)
		parser.state = MAVLINK_PARSE_STATE_GOT_SEQ
	case MAVLINK_PARSE_STATE_GOT_SEQ:
		parser.packet.SysID = c
		parser.crc.WriteByte(c)
		parser.state = MAVLINK_PARSE_STATE_GOT_SYSID
	case MAVLINK_PARSE_STATE_GOT_SYSID:
		parser.packet.CompID = c
		parser.crc.WriteByte(c)
		parser.state = MAVLINK_PARSE_STATE_GOT_COMPID
	case MAVLINK_PARSE_STATE_GOT_COMPID:
		parser.packet.MsgID = MessageID(c)
		parser.crc.WriteByte(c)
		parser.state = MAVLINK_PARSE_STATE_GOT_MSGID1
	case MAVLINK_PARSE_STATE_GOT_MSGID1:
		parser.packet.MsgID += MessageID(c << 8)
		parser.crc.WriteByte(c)
		parser.state = MAVLINK_PARSE_STATE_GOT_MSGID2
	case MAVLINK_PARSE_STATE_GOT_MSGID2:
		parser.packet.MsgID += MessageID(c << 8 * 2)
		parser.crc.WriteByte(c)
		parser.state = MAVLINK_PARSE_STATE_GOT_MSGID3
	case MAVLINK_PARSE_STATE_GOT_MSGID3:
		parser.packet.Payload = append(parser.packet.Payload, c)
		parser.crc.WriteByte(c)
		if len(parser.packet.Payload) == cap(parser.packet.Payload) {
			parser.state = MAVLINK_PARSE_STATE_GOT_PAYLOAD
		}
	case MAVLINK_PARSE_STATE_GOT_PAYLOAD:
		crcExtra, err := dialects.findCrcX(parser.packet.MsgID)
		if err != nil {
			crcExtra = 0
		}
		parser.crc.WriteByte(crcExtra)
		if c != uint8(parser.crc.Sum16()&0xFF) {
			parser.state = MAVLINK_PARSE_STATE_GOT_BAD_CRC
			parser.packet = Packet{}
			return nil, ErrCrcFail
		} else {
			parser.state = MAVLINK_PARSE_STATE_GOT_CRC1
		}
	case MAVLINK_PARSE_STATE_GOT_CRC1:
		if c == uint8(parser.crc.Sum16()>>8) {
			parser.packet.Checksum = parser.crc.Sum16()
			parser.state = MAVLINK_PARSE_STATE_GOT_GOOD_MESSAGE
			result := parser.packet
			parser.packet = Packet{}
			return &result, nil
		} else {
			parser.state = MAVLINK_PARSE_STATE_GOT_BAD_CRC
			parser.packet = Packet{}
			return nil, ErrCrcFail
		}
	}
	return nil, nil
}

func (d *Decoder) PushData(data []byte) {
	d.data <- data
}

func (d *Decoder) NextPacket(duration time.Duration) *Packet {
	select {
	case packet := <-d.decoded:
		return packet
	case <-time.After(duration):
		return nil
	}
}

func (d *Decoder) Stop() {
	d.Multicast.Lock()
	for k, v := range d.Multicast.listeners {
		v <- true
		close(v)
		close(k)
	}
	d.Multicast.Unlock()
}

// NewChannelDecoder function create decoder instance with default dialect
func NewChannelDecoder() *Decoder {
	d := &Decoder{
		Multicast: NewMulticast(),
		data:      make(chan []byte, 256),
		decoded:   make(chan *Packet),
	}
	go func() {
		for {
			buffer := <-d.data
			d.Multicast.notify(buffer)
			for i, c := range buffer {
				if c == magicNumber {
					done := make(chan bool)
					newBytes := d.Multicast.register(done)
					go func() {
						defer d.Multicast.clear(newBytes)
						var parser Parser
						for {
							select {
							case buffer := <-newBytes:
								for _, c := range buffer {
									packet, err := parser.parseChar(c)
									if err != nil {
										return
									} else if packet != nil {
										d.decoded <- packet
										return
									}
								}
							case <-done:
								return
							}
						}
					}()
					newBytes <- buffer[i:]
				}
			}
		}
	}()
	return d
}

func (p *Packet) fixChecksum(dialects DialectSlice) error {
	crc := x25.New()
	crc.WriteByte(byte(len(p.Payload)))
	crc.WriteByte(p.InCompatFlags)
	crc.WriteByte(p.CompatFlags)
	crc.WriteByte(p.SeqID)
	crc.WriteByte(p.SysID)
	crc.WriteByte(p.CompID)
	crc.WriteByte(byte(p.MsgID >> 0))
	crc.WriteByte(byte(p.MsgID >> 8))
	crc.WriteByte(byte(p.MsgID >> 16))
	crc.Write(p.Payload)
	crcx, err := dialects.findCrcX(p.MsgID)
	if err != nil {
		return err
	}
	crc.WriteByte(crcx)
	p.Checksum = crc.Sum16()
	return nil
}

func u16ToBytes(v uint16) []byte {
	return []byte{byte(v & 0xff), byte(v >> 8)}
}
