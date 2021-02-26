/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package mavlink

import (
	"./packet"
	"./parser"
	"bufio"
	"io"
)

// Decoder struct provide decoding processor
type Decoder struct {
	reader  io.ByteReader
	parsers []parser.Parser
}

func (d *Decoder) clearParser(parser parser.Parser) {
	parser.Destroy()
}

func (d *Decoder) clearParsers() {
	for _, parser := range d.parsers {
		d.clearParser(parser)
	}
	d.parsers = d.parsers[:0]
}

// Decode decode input stream to packet. Method return error or nil
func (d *Decoder) Decode() (packet.Packet, error) {
	for {
		c, err := d.reader.ReadByte()
		if err != nil {
			return nil, err
		}
		switch packet.MAGIC_NUMBER(c) {
		case parser.MAGIC_NUMBER_V1: // mavlink1
			d.parsers = append(d.parsers, parser.NewParserV1())
		case parser.MAGIC_NUMBER_V2: // mavlink2
			d.parsers = append(d.parsers, parser.NewParserV2())
		}
		parsers := make([]parser.Parser, 0, len(d.parsers))
		for _, parser := range d.parsers {
			if p, err := parser.ParseChar(c); err != nil {
				d.clearParser(parser)
			} else if !p.IsNil() {
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
		parsers: make([]parser.Parser, 0),
	}
}
