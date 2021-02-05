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
	"github.com/asmyasnikov/go-mavlink/mavlink/helpers"
	"github.com/asmyasnikov/go-mavlink/mavlink/message"
	"github.com/asmyasnikov/go-mavlink/mavlink/packet"
	"github.com/asmyasnikov/go-mavlink/mavlink/register"
)

const (
	// MAGIC_NUMBER_V1 const value for common use
	MAGIC_NUMBER_V1 packet.MAGIC_NUMBER = 0xfe
)

// Packet is a wire type for encoding/decoding mavlink messages.
// use the ToPacket() and FromPacket() routines on specific message
// types to convert them to/from the Message type.
type packet1 struct {
	seqID    uint8             // Sequence of packet
	sysID    uint8             // ID of message sender system/aircraft
	compID   uint8             // ID of the message sender component
	msgID    message.MessageID // ID of message in payload
	payload  []byte
	checksum uint16
}

// NewPacketV1 creates new mavlink1.Packet
func NewPacketV1(sysID uint8, compID uint8, seqID uint8, message message.Message) (packet.Packet, error) {
	payload, err := message.Marshal()
	if err != nil {
		return nil, err
	}
	return &packet1{
		seqID:   seqID,
		sysID:   sysID,
		compID:  compID,
		msgID:   message.MsgID(),
		payload: payload,
	}, nil
}

// IsNil returns true if packet is nil
func (p *packet1) IsNil() bool {
	return p == nil
}

// IsSigned checks whether the frame contains a signature. It does not validate the signature
func (p *packet1) IsSigned() bool {
	return false
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
func (p *packet1) MsgID() message.MessageID {
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
		return errors.ErrNilPointerReference
	}
	p.seqID = rhs.seqID
	p.sysID = rhs.sysID
	p.compID = rhs.compID
	p.msgID = rhs.msgID
	p.checksum = rhs.checksum
	p.payload = append([]byte(nil), rhs.payload...)
	return nil
}

// Copy returns deep copy of packet
func (p *packet1) Copy() packet.Packet {
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
	copy.payload = append([]byte(nil), p.payload...)
	return copy
}

// Unmarshal trying to de-serialize byte slice to packet
func (p *packet1) Unmarshal(buffer []byte) error {
	if p == nil {
		return errors.ErrNilPointerReference
	}
	parser := _parsersPoolV1.Get().(*parser1).reset()
	defer parser.Destroy()
	for _, c := range buffer {
		packet, err := parser.parseChar(c)
		if err != nil {
			return fmt.Errorf("Unmarshal fail with error \"%+v\". Parser %+v", err, parser)
		}
		if packet != nil {
			return p.assign(packet)
		}
	}
	return fmt.Errorf("Unmarshal fail with error \"%+v\". Parser %+v", errors.ErrNoNewData, parser)
}

// Marshal trying to serialize byte slice from packet
func (p *packet1) Marshal() ([]byte, error) {
	if p == nil {
		return nil, errors.ErrNilPointerReference
	}
	bytes := make([]byte, 0, 8+len(p.payload))
	// prepare
	if err := p.prepare(); err != nil {
		return nil, err
	}
	// header
	bytes = append(bytes,
		byte(MAGIC_NUMBER_V1),
		byte(len(p.payload)),
		p.seqID,
		p.sysID,
		p.compID,
		uint8(p.msgID),
	)
	// payload
	bytes = append(bytes, p.payload...)
	bytes = append(bytes, helpers.U16ToBytes(p.checksum)...)
	return bytes, nil
}

func (p *packet1) prepare() error {
	if p == nil {
		return errors.ErrNilPointerReference
	}
	info, err := register.Info(p.msgID)
	if err != nil {
		return err
	}
	crc := crc.NewX25()
	crc.WriteByte(byte(len(p.payload)))
	crc.WriteByte(p.seqID)
	crc.WriteByte(p.sysID)
	crc.WriteByte(p.compID)
	crc.WriteByte(byte(p.msgID >> 0))
	crc.Write(p.payload)
	crc.WriteByte(info.Extra)
	p.checksum = crc.Sum16()
	return nil
}

// Message function produce message from packet
func (p *packet1) Message() (message.Message, error) {
	if p == nil {
		return nil, errors.ErrNilPointerReference
	}
	info, err := register.Info(p.msgID)
	if err != nil {
		return nil, err
	}
	return info.Constructor(p)
}

// String function return string view of Packet struct
func (p *packet1) String() string {
	if p == nil {
		return "nil"
	}
	return fmt.Sprintf(
		"&mavlink1.Packet{ seqID: %d, sysID: %d, compID: %d, msgID: %d, payload: %s, checksum: %d%s }",
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
		"",
	)
}
