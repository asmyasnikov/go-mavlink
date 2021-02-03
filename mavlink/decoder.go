/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package mavlink

import (
	"bufio"
	"io"
)

// Packet is the interface implemented by frames of every supported version.
type Packet interface {
	// SysID returns system id
	SysID() uint8
	// CompID returns component id
	CompID() uint8
	// MsgID returns message id
	MsgID() MessageID
	// Checksum returns packet checksum
	Checksum() uint16
	// SeqID returns packet sequence number
	SeqID() uint8
	// Payload returns packet payload
	Payload() []byte
	// Assign assign internal fields from right hand side packet
	//Assign(rhs Packet) error
	// Copy returns deep copy of packet
	Copy() Packet
	// Message returns dialect message
	Message() (Message, error)
	// String returns string representation of packet
	String() string
	// Marshal encodes Packet to byte slice
	Marshal() ([]byte, error)
	// Unmarshal parses PAYLOAD and stores the result in Packet
	Unmarshal(payload []byte) error
}

// Parser interface is abstract of parsers
type Parser interface {
	ParseChar(c byte) (Packet, error)
	Destroy()
}

// Decoder struct provide decoding processor
type Decoder struct {
	reader  io.ByteReader
	parsers []Parser
}

func (d *Decoder) clearParser(parser Parser) {
	parser.Destroy()
}

func (d *Decoder) clearParsers() {
	for _, parser := range d.parsers {
		d.clearParser(parser)
	}
	d.parsers = d.parsers[:0]
}

// Decode decode input stream to packet. Method return error or nil
func (d *Decoder) Decode() (Packet, error) {
	for {
		c, err := d.reader.ReadByte()
		if err != nil {
			return nil, err
		}
		switch c {
		case 0xfe: // mavlink1
			d.parsers = append(d.parsers, NewParserV1())
		case 0xfd: // mavlink2
			d.parsers = append(d.parsers, NewParserV2())
		}
		parsers := make([]Parser, 0, len(d.parsers))
		for _, parser := range d.parsers {
			if p, err := parser.ParseChar(c); err != nil {
				d.clearParser(parser)
			} else if p != nil {
				d.clearParsers()
				return p, nil
			} else {
				parsers = append(parsers, parser)
			}
		}
		d.parsers = parsers
	}
}

func byteReader(r io.Reader) io.ByteReader {
	if rb, ok := r.(io.ByteReader); ok {
		return rb
	}
	return bufio.NewReader(r)
}

// NewDecoder function create decoder instance
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		reader:  byteReader(r),
		parsers: make([]Parser, 0),
	}
}
