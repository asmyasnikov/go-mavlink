/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package signature

import (
	"../helpers"
	"crypto/sha256"
	"fmt"
	"time"
)

// Signature type
type Signature struct {
	linkID    byte
	timestamp time.Time
	crc       [6]byte
}

const (
	// SIGNATURE_LEN constant
	SIGNATURE_LEN = 13
)

// New returns new Signature struct instance
func New(linkID byte, timestamp time.Time, secret [32]byte, packet []byte) (*Signature, error) {
	if timestamp.Sub(signatureReferenceDate) < 0 {
		return nil, fmt.Errorf("bad timestamp %+v (must be newer than %+v)", timestamp, signatureReferenceDate)
	}
	timestamp = signatureReferenceDate.Add((timestamp.Sub(signatureReferenceDate) / (time.Microsecond * 10)) * (time.Microsecond * 10))
	s := &Signature{
		linkID:    linkID,
		timestamp: timestamp,
	}
	t := uint64(s.timestamp.Sub(signatureReferenceDate) / (time.Microsecond * 10))
	h := sha256.New()
	h.Write(secret[:])
	h.Write(packet[:])
	h.Write([]byte{s.linkID})
	h.Write(helpers.U48ToBytes(t))
	copy(s.crc[:], h.Sum(nil)[:6])
	return s, nil
}

// Copy returns copy of Signature struct instance
func (s *Signature) Copy() *Signature {
	if s == nil {
		return nil
	}
	c := &Signature{
		linkID:    s.linkID,
		timestamp: s.timestamp,
	}
	copy(c.crc[:], s.crc[:])
	return c
}

// Unmarshal returns de-serialized bytes to Signature struct
func (s *Signature) Unmarshal(buffer []byte) error {
	s.linkID = buffer[0]
	s.timestamp = signatureReferenceDate.Add(time.Duration(helpers.BytesToU48(buffer[1:7])) * (time.Microsecond * 10))
	copy(s.crc[:], buffer[7:13])
	return nil
}

// Marshal returns serialized bytes from Signature struct
func (s *Signature) Marshal() ([]byte, error) {
	bytes := make([]byte, 0, SIGNATURE_LEN)
	bytes = append(bytes, byte(s.linkID))
	bytes = append(bytes, helpers.U48ToBytes(uint64(s.timestamp.Sub(signatureReferenceDate)/(time.Microsecond*10)))...)
	bytes = append(bytes, s.crc[:]...)
	return bytes, nil
}

// LinkID returns link id
func (s *Signature) LinkID() byte {
	return s.linkID
}

// 1st January 2015 GMT https://mavlink.io/en/guide/message_signing.html#timestamps
var signatureReferenceDate = time.Date(2015, 01, 01, 0, 0, 0, 0, time.UTC)

// Timestamp returns timestamp of signature
func (s *Signature) Timestamp() time.Time {
	return s.timestamp
}

// Signature returns signature slice
func (s *Signature) CRC() (crc [6]byte) {
	copy(crc[:], s.crc[:])
	return crc
}

// String function return string view of Packet struct
func (s Signature) String() string {
	return fmt.Sprintf(
		"&Signature{ linkID: 0x%02X, timestamp: \"%+v\", crc: \"%06X\" }",
		s.LinkID(),
		s.Timestamp().UTC(),
		s.CRC(),
	)
}
