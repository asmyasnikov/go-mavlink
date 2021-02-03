/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package mavlink

import (
	"fmt"
)

// Packet is a wire type for encoding/decoding mavlink messages.
// use the ToPacket() and FromPacket() routines on specific message
// types to convert them to/from the Message type.
type packet1 struct {
	seqID    uint8     // Sequence of packet
	sysID    uint8     // ID of message sender system/aircraft
	compID   uint8     // ID of the message sender component
	msgID    MessageID // ID of message in payload
	payload  []byte
	checksum uint16
}

// SysID returns system id
func (p *packet1) SysID() uint8 {
	return p.sysID
}

// CompID returns component id
func (p *packet1) CompID() uint8 {
	return p.compID
}

// MsgID returns message id
func (p *packet1) MsgID() MessageID {
	return p.msgID
}

// Checksum returns packet checksum
func (p *packet1) Checksum() uint16 {
	return p.checksum
}

// SeqID returns packet sequence number
func (p *packet1) SeqID() uint8 {
	return p.seqID
}

// Payload returns packet payload
func (p *packet1) Payload() []byte {
	// todo: maybe COPY???
	return p.payload
}

// Assign assign internal fields from right hand side packet
func (p *packet1) Assign(rhs Packet) error {
	packet, ok := rhs.(*packet1)
	if !ok {
		return fmt.Errorf("cast interface '%+v' to '*packet1' fail", rhs)
	}
	packet.seqID = p.seqID
	packet.sysID = p.sysID
	packet.compID = p.compID
	packet.msgID = p.msgID
	packet.checksum = p.checksum
	packet.payload = append(packet.payload[:0], p.payload...)
	return nil
}

func (p *packet1) nextSeqNum() byte {
	currentSeqNum++
	return currentSeqNum
}

// Encode trying to encode message to packet
func (p *packet1) encode(sysID, compID uint8, m Message) error {
	p.seqID = p.nextSeqNum()
	p.sysID = sysID
	p.compID = compID
	return p.Encode(m)
}

// Encode trying to encode message to packet
func (p *packet1) Encode(m Message) error {
	if err := m.Pack(p); err != nil {
		return err
	}
	if err := p.fixChecksum(); err != nil {
		return err
	}
	return nil
}

// Decode trying to decode message to packet
func (p *packet1) Decode(m Message) error {
	return m.Unpack(p)
}

// Unmarshal trying to de-serialize byte slice to packet
func unmarshal1(buffer []byte, p *packet1) error {
	parser := _parsersPool_v1.Get().(*parser1)
	defer parser.Destroy()
	for _, c := range buffer {
		packet, err := parser.ParseChar(c)
		if err != nil {
			return err
		}
		if packet != nil {
			return p.Assign(packet)
		}
	}
	return ErrNoNewData
}

// Marshal trying to serialize byte slice from packet
func marshal1(p *packet1) ([]byte, error) {
	if p == nil {
		return nil, ErrNilPointerReference
	}
	bytes := make([]byte, 0, 8+len(p.payload))
	// header
	bytes = append(bytes,
		0xfe,
		byte(len(p.payload)),
		p.seqID,
		p.sysID,
		p.compID,
		uint8(p.msgID),
	)
	// payload
	bytes = append(bytes, p.payload...)
	// crc
	bytes = append(bytes, p.u16ToBytes(p.checksum)...)
	return bytes, nil
}

func (p *packet1) fixChecksum() error {
	msg, ok := supported[p.msgID]
	if !ok {
		return ErrUnknownMsgID
	}
	crc := NewX25()
	crc.WriteByte(byte(len(p.payload)))
	crc.WriteByte(p.seqID)
	crc.WriteByte(p.sysID)
	crc.WriteByte(p.compID)
	crc.WriteByte(byte(p.msgID >> 0))
	crc.Write(p.payload)
	crc.WriteByte(msg.Extra)
	p.checksum = crc.Sum16()
	return nil
}

func (p *packet1) u16ToBytes(v uint16) []byte {
	return []byte{byte(v & 0xff), byte(v >> 8)}
}

// Message function produce message from packet
func (p *packet1) Message() (Message, error) {
	msg, ok := supported[p.msgID]
	if !ok {
		return nil, ErrUnknownMsgID
	}
	return msg.Constructor(p), nil
}

// String function return string view of Packet struct
func (p *packet1) String() string {
	return fmt.Sprintf(
		"&packet1{ seqID: %d, sysID: %d, compID: %d, msgID: %d, payload: %s, checksum: %d }",
		p.seqID,
		p.sysID,
		p.compID,
		int64(p.msgID),
		func() string {
			msg, err := p.Message()
			if err != nil {
				return fmt.Sprintf("%0X", p.payload)
			}
			return msg.String()
		}(),
		p.checksum,
	)
}
