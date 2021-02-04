/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package message

// MessageID typedef
type MessageID uint32

// Message is a basic type for encoding/decoding mavlink messages.
// use the Pack() and Unpack() routines on specific message
// types to convert them to/from the Packet type.
type Message interface {
	// MsgID returns message id
	MsgID() MessageID
	// String returns human-readable representation on Message
	String() string
	// Marshal encodes Packet to byte slice
	Marshal() ([]byte, error)
	// Unmarshal parses PAYLOAD and stores the result in Packet
	Unmarshal(payload []byte) error
}
