/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package signature

import (
	"fmt"
	"time"
)

// Signature type
type Signature []byte

const (
	// SIGNATURE_LEN constant
	SIGNATURE_LEN = 13
)

// NewSignature creates signature
func NewSignature(linkID byte, timestamp time.Time, signature [6]byte) Signature {
	s := make([]byte, SIGNATURE_LEN)
	s[0] = linkID
	t := timestamp.Sub(signatureReferenceDate) / (time.Microsecond * 10)
	s[1] = byte(t)
	s[2] = byte(t >> 8)
	s[3] = byte(t >> 16)
	s[4] = byte(t >> 24)
	s[5] = byte(t >> 32)
	s[6] = byte(t >> 40)
	copy(s[7:], signature[:])
	return s
}

// LinkID returns link id
func (s Signature) LinkID() byte {
	return s[0]
}

// 1st January 2015 GMT https://mavlink.io/en/guide/message_signing.html#timestamps
var signatureReferenceDate = time.Date(2015, 01, 01, 0, 0, 0, 0, time.UTC)

// Timestamp returns timestamp of signature
func (s Signature) Timestamp() time.Time {
	return signatureReferenceDate.Add(time.Duration(uint64(s[6])<<40|uint64(s[5])<<32|uint64(s[4])<<24|uint64(s[3])<<16|uint64(s[2])<<8|uint64(s[1])) * (10 * time.Microsecond))
}

// Signature returns signature slice
func (s Signature) Signature() (signature [6]byte) {
	copy(signature[:], s[7:])
	return signature
}

// String function return string view of Packet struct
func (s Signature) String() string {
	return fmt.Sprintf(
		"&Signature{ linkID: 0x%02X, timestamp: \"%+v\", signature: \"%06X\" }",
		s.LinkID(),
		s.Timestamp(),
		s.Signature(),
	)
}
