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

// Nil returns true if packet is nil
func (p *packet1) Nil() bool {
	return p == nil
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
	return append([]byte(nil), p.payload...)
}

func (p *packet1) assign(rhs *packet1) error {
	if p == nil {
		return ErrNilPointerReference
	}
	p.seqID = rhs.seqID
	p.sysID = rhs.sysID
	p.compID = rhs.compID
	p.msgID = rhs.msgID
	p.checksum = rhs.checksum
	p.payload = append([]byte(nil), rhs.payload...)
	return nil
}

/*
// Assign assign internal fields from right hand side packet
func (p *packet1) Assign(rhs Packet) error {
    if p == nil {
        return ErrNilPointerReference
    }
    packet, ok := rhs.(*packet1)
    if !ok {
        return fmt.Errorf("cast interface '%+v' to '*packet1' fail", rhs)
    }
    p.seqID = p.SeqID()
    p.sysID = p.SysID()
    p.compID = p.CompID()
    p.msgID = p.MsgID()
    p.checksum = p.Checksum()
    p.payload = p.Payload()
    return nil
}
*/

// Copy returns deep copy of packet
func (p *packet1) Copy() Packet {
	return p.copy()
}

// Copy returns deep copy of packet
func (p *packet1) copy() *packet1 {
	if p == nil {
		return nil
	}
	copy := &packet1{}
	copy.seqID = p.seqID
	copy.sysID = p.sysID
	copy.compID = p.compID
	copy.msgID = p.msgID
	copy.checksum = p.checksum
	copy.payload = p.payload
	return copy
}

/*
// Encode trying to encode message to packet
func (p *packet1) encode(sysID, compID uint8, m Message) error {
    if p == nil {
        return ErrNilPointerReference
    }
	p.seqID = p.nextSeqNum()
	p.sysID = sysID
	p.compID = compID
	return p.Encode(m)
}

// Encode trying to encode message to packet
func (p *packet1) Encode(m Message) error {
    if p == nil {
        return ErrNilPointerReference
    }
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
    if p == nil {
        return ErrNilPointerReference
    }
    return m.Unpack(p)
}
*/

// Unmarshal trying to de-serialize byte slice to packet
func (p *packet1) Unmarshal(buffer []byte) error {
	if p == nil {
		return ErrNilPointerReference
	}
	parser := _parsersPool_v1.Get().(*parser1)
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
func (p *packet1) Marshal() ([]byte, error) {
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
	if p == nil {
		return ErrNilPointerReference
	}
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
func (p *packet1) String() string {
	if p == nil {
		return "nil"
	}
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
