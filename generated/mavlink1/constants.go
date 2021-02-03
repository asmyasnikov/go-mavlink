/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package mavlink

import (
	"errors"
)

// magicNumber constant for mavlink 1.0
const (
	magicNumber = 0xfe
)

// MessageID typedef
type MessageID uint8

// MAVLINK_PARSE_STATE typedef
type MAVLINK_PARSE_STATE int

// MAVLINK_PARSE_STATES
const (
	MAVLINK_PARSE_STATE_UNINIT           MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_IDLE             MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_STX          MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_LENGTH       MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_SEQ          MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_SYSID        MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_COMPID       MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_MSGID1       MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_PAYLOAD      MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_CRC1         MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_BAD_CRC      MAVLINK_PARSE_STATE = iota
	MAVLINK_PARSE_STATE_GOT_GOOD_MESSAGE MAVLINK_PARSE_STATE = iota
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
