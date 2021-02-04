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

// INCOMPAT_FLAGS type
type INCOMPAT_FLAGS uint8

const (
	// MAGIC_NUMBER_V2 const value for common use
	MAGIC_NUMBER_V2       packet.MAGIC_NUMBER = 0xfd
	INCOMPAT_FLAGS_SIGNED INCOMPAT_FLAGS      = 0b00000001
)

// Packet is a wire type for encoding/decoding mavlink messages.
// use the ToPacket() and FromPacket() routines on specific message
// types to convert them to/from the Message type.
type packet2 struct {
	incompatFlags uint8             // incompat flags
	compatFlags   uint8             // compat flags
	seqID         uint8             // Sequence of packet
	sysID         uint8             // ID of message sender system/aircraft
	compID        uint8             // ID of the message sender component
	msgID         message.MessageID // ID of message in payload
	payload       []byte
	checksum      uint16
}

func NewPacketV2(sysID uint8, compID uint8, seqID uint8, message message.Message) (packet.Packet, error) {
	payload, err := message.Marshal()
	if err != nil {
		return nil, err
	}
	return &packet2{
		seqID:   seqID,
		sysID:   sysID,
		compID:  compID,
		msgID:   message.MsgID(),
		payload: payload,
	}, nil
}

// IsNil returns true if packet is nil
func (p *packet2) IsNil() bool {
	return p == nil
}

// IsSigned checks whether the frame contains a signature. It does not validate the signature
func (p *packet2) IsSigned() bool {
	return (INCOMPAT_FLAGS(p.incompatFlags) & INCOMPAT_FLAGS_SIGNED) != 0
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
func (p *packet2) MsgID() message.MessageID {
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
		return errors.ErrNilPointerReference
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
func (p *packet2) Copy() packet.Packet {
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
		return errors.ErrNilPointerReference
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
	return errors.ErrNoNewData
}

// Marshal trying to serialize byte slice from packet
func (p *packet2) Marshal() ([]byte, error) {
	if p == nil {
		return nil, errors.ErrNilPointerReference
	}
	bytes := make([]byte, 0, 12+len(p.payload))
	// prepare
	if err := p.prepare(); err != nil {
		return nil, err
	}
	// header
	bytes = append(bytes,
		byte(MAGIC_NUMBER_V2),
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
	bytes = append(bytes, helpers.U16ToBytes(p.checksum)...)
	return bytes, nil
}

func (p *packet2) prepare() error {
	if p == nil {
		return errors.ErrNilPointerReference
	}
	// payload minify
	payloadLen := len(p.payload)
	for payloadLen > 1 && p.payload[payloadLen-1] == 0 {
		payloadLen--
	}
	p.payload = p.payload[:payloadLen]
	info, err := register.Info(p.msgID)
	if err != nil {
		return err
	}
	crc := crc.NewX25()
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
	crc.WriteByte(info.Extra)
	p.checksum = crc.Sum16()
	return nil
}

// Message function produce message from packet
func (p *packet2) Message() (message.Message, error) {
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
func (p *packet2) String() string {
	if p == nil {
		return "nil"
	}
	return fmt.Sprintf(
		"&mavlink2.Packet{ incompatFlags: %08b, compatFlags: %08b, seqID: %d, sysID: %d, compID: %d, msgID: %d, payload: %s, checksum: %d }",
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
