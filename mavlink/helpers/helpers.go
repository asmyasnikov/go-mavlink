/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package helpers

import (
	"math/rand"
)

// U16ToBytes creates bytearray representation
func U16ToBytes(v uint16) []byte {
	return []byte{byte(v & 0xff), byte(v >> 8)}
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
