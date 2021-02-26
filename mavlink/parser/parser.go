/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package parser

import (
	"../packet"
)

// Parser interface is abstract of parsers
type Parser interface {
	ParseChar(c byte) (packet.Packet, error)
	String() string
	Destroy()
}
