/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package helpers

import (
	"math/rand"
)

// U16ToBytes makes bytearray representation
func U16ToBytes(v uint16) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),
	}
}

// RandomByte generate random byte
func RandomByte() byte {
	return RandomBytes(1)[0]
}

// RandomBytes generate random byte
func RandomBytes(size int) []byte {
	buffer := make([]byte, size)
	_, _ = rand.Read(buffer[:])
	return buffer
}

// BytesToU24 makes uint32 (like as uint24) from bytearray
func BytesToU24(v []byte) uint32 {
	return uint32(v[2])<<16 |
		uint32(v[1])<<8 |
		uint32(v[0])
}

// U24ToBytes makes bytearray from uint32 value (like as uin24)
func U24ToBytes(v uint32) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
	}
}

// BytesToU48 makes uint64 value (like as uint48)
func BytesToU48(v []byte) uint64 {
	return uint64(v[5])<<40 |
		uint64(v[4])<<32 |
		uint64(v[3])<<24 |
		uint64(v[2])<<16 |
		uint64(v[1])<<8 |
		uint64(v[0])
}

// U48ToBytes makes bytearray from uint64 value (like as uint48)
func U48ToBytes(v uint64) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
		byte(v >> 24),
		byte(v >> 32),
		byte(v >> 40),
	}
}
