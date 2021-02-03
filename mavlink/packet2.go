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
type packet2 struct {
	incompatFlags uint8     // incompat flags
	compatFlags   uint8     // compat flags
	seqID         uint8     // Sequence of packet
	sysID         uint8     // ID of message sender system/aircraft
	compID        uint8     // ID of the message sender component
	msgID         MessageID // ID of message in payload
	payload       []byte
	checksum      uint16
}

// Nil returns true if packet is nil
func (p *packet2) Nil() bool {
	return p == nil
}

// SysID returns system id
func (p *packet2) SysID() uint8 {
	return p.sysID
}

// CompID returns component id
func (p *packet2) CompID() uint8 {
	return p.compID
}

// MsgID returns message id
func (p *packet2) MsgID() MessageID {
	return p.msgID
}

// Checksum returns packet checksum
func (p *packet2) Checksum() uint16 {
	return p.checksum
}

// SeqID returns packet sequence number
func (p *packet2) SeqID() uint8 {
	return p.seqID
}

// Payload returns packet payload
func (p *packet2) Payload() []byte {
	return append([]byte(nil), p.payload...)
}

func (p *packet2) assign(rhs *packet2) error {
	if p == nil {
		return ErrNilPointerReference
	}
	p.incompatFlags = rhs.incompatFlags
	p.compatFlags = rhs.compatFlags
	p.seqID = rhs.seqID
	p.sysID = rhs.sysID
	p.compID = rhs.compID
	p.msgID = rhs.msgID
	p.checksum = rhs.checksum
	p.payload = append([]byte(nil), rhs.payload...)
	return nil
}

// Copy returns deep copy of packet
func (p *packet2) Copy() Packet {
	return p.copy()
}

// Copy returns deep copy of packet
func (p *packet2) copy() *packet2 {
	if p == nil {
		return nil
	}
	copy := &packet2{}
	copy.incompatFlags = p.incompatFlags
	copy.compatFlags = p.compatFlags
	copy.seqID = p.seqID
	copy.sysID = p.sysID
	copy.compID = p.compID
	copy.msgID = p.msgID
	copy.checksum = p.checksum
	copy.payload = append([]byte(nil), p.payload...)
	return copy
}

// Unmarshal trying to de-serialize byte slice to packet
func (p *packet2) Unmarshal(buffer []byte) error {
	if p == nil {
		return ErrNilPointerReference
	}
	parser := _parsersPool_v2.Get().(*parser2)
	defer parser.Destroy()
	for _, c := range buffer {
		packet, err := parser.parseChar(c)
		if err != nil {
			return err
		}
		if packet != nil {
			return p.assign(packet)
		}
	}
	return ErrNoNewData
}

// Marshal trying to serialize byte slice from packet
func (p *packet2) Marshal() ([]byte, error) {
	if p == nil {
		return nil, ErrNilPointerReference
	}
	if err := p.fixChecksum(); err != nil {
		return nil, err
	}
	bytes := make([]byte, 0, 12+len(p.payload))
	// header
	bytes = append(bytes,
		0xfd,
		byte(len(p.payload)),
		uint8(p.incompatFlags),
		uint8(p.compatFlags),
		p.seqID,
		p.sysID,
		p.compID,
		uint8(p.msgID),
		uint8(p.msgID>>8),
		uint8(p.msgID>>16),
	)
	// payload
	bytes = append(bytes, p.payload...)
	// crc
	bytes = append(bytes, u16ToBytes(p.checksum)...)
	return bytes, nil
}

func (p *packet2) fixChecksum() error {
	if p == nil {
		return ErrNilPointerReference
	}
	msg, ok := supported[p.msgID]
	if !ok {
		return ErrUnknownMsgID
	}
	crc := NewX25()
	crc.WriteByte(byte(len(p.payload)))
	crc.WriteByte(p.incompatFlags)
	crc.WriteByte(p.compatFlags)
	crc.WriteByte(p.seqID)
	crc.WriteByte(p.sysID)
	crc.WriteByte(p.compID)
	crc.WriteByte(byte(p.msgID >> 0))
	crc.WriteByte(byte(p.msgID >> 8))
	crc.WriteByte(byte(p.msgID >> 16))
	crc.Write(p.payload)
	crc.WriteByte(msg.Extra)
	p.checksum = crc.Sum16()
	return nil
}

// Message function produce message from packet
func (p *packet2) Message() (Message, error) {
	if p == nil {
		return nil, ErrNilPointerReference
	}
	msg, ok := supported[p.msgID]
	if !ok {
		return nil, ErrUnknownMsgID
	}
	return msg.Constructor(p)
}

// String function return string view of Packet struct
func (p *packet2) String() string {
	if p == nil {
		return "nil"
	}
	return fmt.Sprintf(
		"&packet2{ incompatFlags: %08b, compatFlags: %08b, seqID: %d, sysID: %d, compID: %d, msgID: %d, payload: %s, checksum: %d }",
		p.incompatFlags,
		p.compatFlags,
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
