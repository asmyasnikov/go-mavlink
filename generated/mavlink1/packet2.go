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
	// todo: maybe COPY???
	return p.payload
}

// Assign assign internal fields from right hand side packet
func (p *packet2) Assign(rhs Packet) error {
	packet, ok := rhs.(*packet2)
	if !ok {
		return fmt.Errorf("cast interface '%+v' to '*packet2' fail", rhs)
	}
	packet.incompatFlags = p.incompatFlags
	packet.compatFlags = p.compatFlags
	packet.seqID = p.seqID
	packet.sysID = p.sysID
	packet.compID = p.compID
	packet.msgID = p.msgID
	packet.checksum = p.checksum
	packet.payload = append(packet.payload[:0], p.payload...)
	return nil
}

func (p *packet2) nextSeqNum() byte {
	currentSeqNum++
	return currentSeqNum
}

// Encode trying to encode message to packet
func (p *packet2) encode(sysID, compID uint8, m Message) error {
	p.seqID = p.nextSeqNum()
	p.sysID = sysID
	p.compID = compID
	return p.Encode(m)
}

// Encode trying to encode message to packet
func (p *packet2) Encode(m Message) error {
	if err := m.Pack(p); err != nil {
		return err
	}
	payloadLen := len(p.payload)
	for payloadLen > 1 && p.payload[payloadLen-1] == 0 {
		payloadLen--
	}
	p.payload = p.payload[:payloadLen]
	if err := p.fixChecksum(); err != nil {
		return err
	}
	return nil
}

// Decode trying to decode message to packet
func (p *packet2) Decode(m Message) error {
	if msg, ok := supported[p.msgID]; !ok {
		return ErrUnknownMsgID
	} else if len(p.payload) < msg.Size {
		p.payload = append(p.payload, zeroTail[:msg.Size-len(p.payload)]...)
	}
	return m.Unpack(p)
}

// Unmarshal trying to de-serialize byte slice to packet
func unmarshal2(buffer []byte, p *packet2) error {
	parser := _parsersPool_v2.Get().(*parser2)
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
func marshal2(p *packet2) ([]byte, error) {
	if p == nil {
		return nil, ErrNilPointerReference
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
	bytes = append(bytes, p.u16ToBytes(p.checksum)...)
	return bytes, nil
}

func (p *packet2) fixChecksum() error {
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

func (p *packet2) u16ToBytes(v uint16) []byte {
	return []byte{byte(v & 0xff), byte(v >> 8)}
}

// Message function produce message from packet
func (p *packet2) Message() (Message, error) {
	msg, ok := supported[p.msgID]
	if !ok {
		return nil, ErrUnknownMsgID
	}
	return msg.Constructor(p), nil
}

// String function return string view of Packet struct
func (p *packet2) String() string {
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
