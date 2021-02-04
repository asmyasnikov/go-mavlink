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
	writer        io.Writer
	currentSeqNum byte
}

func (e *Encoder) nextSeqNum() byte {
	e.currentSeqNum++
	return e.currentSeqNum
}

// Encode encode packet to output stream. Method return error or nil
func (e *Encoder) Encode(v version.MAVLINK_VERSION, sysID uint8, compID uint8, message message.Message) error {
	p, err := MakePacket(v, sysID, compID, e.nextSeqNum(), message)
	if err != nil {
		return err
	}
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

// NewEncoder function create encoder instance
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		writer:        w,
		currentSeqNum: 0,
	}
}

func MakePacket(v version.MAVLINK_VERSION, sysID uint8, compID uint8, seqID uint8, message message.Message) (packet.Packet, error) {
	switch v {
	case version.MAVLINK_V1:
		return parser.MakePacketV1(sysID, compID, seqID, message)
	case version.MAVLINK_V2:
		return parser.MakePacketV2(sysID, compID, seqID, message)
	default:
		return nil, fmt.Errorf("Unknown mavlink version %d", v)
	}
}
