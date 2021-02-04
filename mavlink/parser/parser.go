/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package parser

import (
	"github.com/asmyasnikov/go-mavlink/mavlink/packet"
)

// Parser interface is abstract of parsers
type Parser interface {
	ParseChar(c byte) (packet.Packet, error)
	Destroy()
}
