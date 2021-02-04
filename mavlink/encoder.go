/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package mavlink

import (
	"fmt"
	"github.com/asmyasnikov/go-mavlink/mavlink/message"
	"github.com/asmyasnikov/go-mavlink/mavlink/packet"
	"github.com/asmyasnikov/go-mavlink/mavlink/parser"
	"github.com/asmyasnikov/go-mavlink/mavlink/version"
	"io"
)

// Encoder struct provide decoding processor
type Encoder struct {
	writer io.Writer
}

var currentSeqNum uint8

func nextSeqNum() byte {
	currentSeqNum++
	return currentSeqNum
}

// Encode encode packet to output stream. Method return error or nil on success
func (e *Encoder) Encode(p packet.Packet) error {
	b, err := p.Marshal()
	if err != nil {
		return err
	}
	n, err := e.writer.Write(b)
	if len(b) != n {
		return fmt.Errorf("writed %d bytes, but need to write %d bytes", n, len(b))
	}
	return err
}

// NewEncoder function creates encoder instance
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		writer: w,
	}
}

// NewPacket function creates packet
func NewPacket(v version.MAVLINK_VERSION, sysID uint8, compID uint8, message message.Message) (packet.Packet, error) {
	switch v {
	case version.MAVLINK_V1:
		return parser.NewPacketV1(sysID, compID, nextSeqNum(), message)
	case version.MAVLINK_V2:
		return parser.NewPacketV2(sysID, compID, nextSeqNum(), message)
	default:
		return nil, fmt.Errorf("Unknown mavlink version %d", v)
	}
}
