/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package mavlink

import (
	"errors"
)

var (
	// ErrUnknownMsgID define
	ErrUnknownMsgID = errors.New("unknown msg id")
	// ErrCrcFail define
	ErrCrcFail = errors.New("checksum did not match")
	// ErrNoNewData define
	ErrNoNewData = errors.New("no new data")
	// ErrNilPointerReference define
	ErrNilPointerReference = errors.New("nil pointer reference")
	// ErrPayloadTooSmall define
	ErrPayloadTooSmall = errors.New("payload too small")
	// zeroTail is a cache of zero slice for auto append tail to
	// payload in Mavlink2 messages with trimmed payload (variable length)
	zeroTail = make([]byte, 256)
	// currentSeqNum
	currentSeqNum uint8
)
