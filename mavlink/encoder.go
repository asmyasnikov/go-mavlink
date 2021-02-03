/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package mavlink

import (
	"fmt"
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

func (e *Encoder) makePacket(version MAVLINK_VERSION, sysID uint8, compID uint8, message Message) (Packet, error) {
	payload, err := message.Marshal()
	if err != nil {
		return nil, err
	}
	switch version {
	case MAVLINK_V1:
		return &packet1{
			seqID:   e.nextSeqNum(),
			sysID:   sysID,
			compID:  compID,
			msgID:   message.MsgID(),
			payload: payload,
		}, nil
	case MAVLINK_V2:
		return &packet2{
			seqID:   e.nextSeqNum(),
			sysID:   sysID,
			compID:  compID,
			msgID:   message.MsgID(),
			payload: payload,
		}, nil
	default:
		return nil, fmt.Errorf("Undefined mavlink version %d", version)
	}
}

// Encode encode packet to output stream. Method return error or nil
func (e *Encoder) Encode(version MAVLINK_VERSION, sysID uint8, compID uint8, message Message) error {
	packet, err := e.makePacket(version, sysID, compID, message)
	if err != nil {
		return err
	}
	b, err := packet.Marshal()
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
