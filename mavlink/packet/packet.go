/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package packet

import (
	"github.com/asmyasnikov/go-mavlink/mavlink/message"
	"github.com/asmyasnikov/go-mavlink/mavlink/signature"
	"github.com/asmyasnikov/go-mavlink/mavlink/version"
	"time"
)

// MAGIC_NUMBER type
type MAGIC_NUMBER uint8

// Packet is the interface implemented by frames of every supported version.
type Packet interface {
	// Version returns version of packet framing
	Version() version.MAVLINK_VERSION
	// IsNil returns true if packet is nil
	IsNil() bool
	// IsSigned checks whether the frame contains a signature. It does not validate the signature
	IsSigned() bool
	// SysID returns system id
	SysID() uint8
	// CompID returns component id
	CompID() uint8
	// MsgID returns message id
	MsgID() message.MessageID
	// Checksum returns packet checksum
	Checksum() uint16
	// SeqID returns packet sequence number
	SeqID() uint8
	// Payload returns packet payload
	Payload() []byte
	// Signature returns packet signature
	Signature() *signature.Signature
	// Copy returns deep copy of packet
	Copy() Packet
	// Message returns dialect message
	Message() (message.Message, error)
	// String returns string representation of packet
	String() string
	// Marshal encodes Packet to byte slice
	Marshal() ([]byte, error)
	// MarshalWithSignature encodes Packet to byte slice including signature at appendix
	MarshalWithSignature(linkID byte, timestamp time.Time, secretKey [32]byte) ([]byte, error)
	// Unmarshal parses PAYLOAD and stores the result in Packet
	Unmarshal(payload []byte) error
}
