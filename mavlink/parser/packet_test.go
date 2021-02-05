package parser

import (
	"fmt"
	"github.com/asmyasnikov/go-mavlink/mavlink/mock"
	"github.com/asmyasnikov/go-mavlink/mavlink/packet"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestMarshallUnmarshal(t *testing.T) {
	msg := &mock.Ping{
		Seq: 9999,
	}
	payload, err := msg.Marshal()
	require.NoError(t, err)
	cases := []struct {
		packet packet.Packet
		bytes  []byte
	}{
		{
			&packet1{msgID: msg.MsgID(), payload: payload},
			[]byte{0xfe, 0x4, 0x0, 0x0, 0x0, 0xde, 0xf, 0x27, 0x0, 0x0, 0xf4, 0xe2},
		},
		{
			&packet2{msgID: msg.MsgID(), payload: payload},
			[]byte{0xfd, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0xde, 0x0, 0x0, 0xf, 0x27, 0x7, 0xc5},
		},
		{
			&packet2{
				msgID:         msg.MsgID(),
				payload:       payload,
				incompatFlags: 0x01,
				signature:     []byte{0x0e, 0x47, 0x04, 0x0c, 0xef, 0x3b, 0x0e, 0x47, 0x04, 0x0c, 0xef, 0xef, 0x9b},
			},
			[]byte{0xfd, 0x2, 0x1, 0x0, 0x0, 0x0, 0x0, 0xde, 0x0, 0x0, 0xf, 0x27, 0x96, 0x90, 0xe, 0x47, 0x4, 0xc, 0xef, 0x3b, 0xe, 0x47, 0x4, 0xc, 0xef, 0xef, 0x9b},
		},
	}
	for i, c := range cases {
		b, err := c.packet.Marshal()
		require.NoError(t, err, fmt.Sprintf("case %d: { packet: %+v, bytes: %+v }", i, c.packet, c.bytes))
		require.Equal(t, c.bytes, b, fmt.Sprintf("case %d: { packet: %+v, bytes: %+v }", i, c.packet, c.bytes))
	}
	for i, c := range cases {
		p := c.packet.Copy()
		err := p.Unmarshal(c.bytes)
		require.NoError(t, err, fmt.Sprintf("case %d: { packet: %+v, bytes: %+v }", i, c.packet, c.bytes))
		m, err := c.packet.Message()
		require.NoError(t, err, fmt.Sprintf("case %d: { packet: %+v, bytes: %+v }", i, c.packet, c.bytes))
		ping, ok := m.(*mock.Ping)
		require.True(t, ok, fmt.Sprintf("case %d: { packet: %+v, bytes: %+v }", i, c.packet, c.bytes))
		require.Equal(t, msg.Seq, ping.Seq, fmt.Sprintf("case %d: { packet: %+v, bytes: %+v }", i, c.packet, c.bytes))
	}
}

func TestPacketV2Signature(t *testing.T) {
	msg := &mock.Ping{
		Seq: 9999,
	}
	payload, err := msg.Marshal()
	require.NoError(t, err)
	cases := []struct {
		packet    *packet2
		bytes     []byte
		signature Signature
	}{
		{
			&packet2{msgID: msg.MsgID(), payload: payload},
			[]byte{0xfd, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0xde, 0x0, 0x0, 0xf, 0x27, 0x7, 0xc5},
			nil,
		},
		{
			&packet2{
				msgID:         msg.MsgID(),
				payload:       payload,
				incompatFlags: 0x01,
				signature:     []byte{0x0e, 0x47, 0x04, 0x0c, 0xef, 0x3b, 0x0e, 0x47, 0x04, 0x0c, 0xef, 0xef, 0x9b},
			},
			[]byte{0xfd, 0x2, 0x1, 0x0, 0x0, 0x0, 0x0, 0xde, 0x0, 0x0, 0xf, 0x27, 0x96, 0x90, 0xe, 0x47, 0x4, 0xc, 0xef, 0x3b, 0xe, 0x47, 0x4, 0xc, 0xef, 0xef, 0x9b},
			Signature{0xe, 0x47, 0x4, 0xc, 0xef, 0x3b, 0xe, 0x47, 0x4, 0xc, 0xef, 0xef, 0x9b},
		},
	}
	for i, c := range cases {
		b, err := c.packet.Copy().Marshal()
		require.NoError(t, err, fmt.Sprintf("case %d: { packet: %+v, bytes: %+v }", i, c.packet, c.bytes))
		require.Equal(t, c.bytes, b, fmt.Sprintf("case %d: { packet: %+v, bytes: %+v }", i, c.packet, c.bytes))
		require.Equal(t, c.signature, c.packet.signature, fmt.Sprintf("case %d: { packet: %+v, bytes: %+v }", i, c.packet, c.bytes))
	}
	for i, c := range cases {
		p := packet2{}
		err := p.Unmarshal(c.bytes)
		require.NoError(t, err, fmt.Sprintf("case %d: { packet: %+v, bytes: %+v }", i, c.packet, c.bytes))
		require.Equal(t, c.signature, c.packet.signature, fmt.Sprintf("case %d: { packet: %+v, bytes: %+v }", i, c.packet, c.bytes))
		m, err := p.Message()
		require.NoError(t, err, fmt.Sprintf("case %d: { packet: %+v, bytes: %+v }", i, c.packet, c.bytes))
		ping, ok := m.(*mock.Ping)
		require.True(t, ok, fmt.Sprintf("case %d: { packet: %+v, bytes: %+v }", i, c.packet, c.bytes))
		require.Equal(t, msg.Seq, ping.Seq, fmt.Sprintf("case %d: { packet: %+v, bytes: %+v }", i, c.packet, c.bytes))
	}
}

func TestPacketCopy(t *testing.T) {
	msg := &mock.Ping{
		Seq: 9999,
	}
	payload, err := msg.Marshal()
	require.NoError(t, err)
	cases := []packet.Packet{
		&packet1{msgID: msg.MsgID(), payload: payload},
		&packet2{msgID: msg.MsgID(), payload: payload},
		&packet2{
			msgID:         msg.MsgID(),
			payload:       payload,
			incompatFlags: 0x01,
			signature:     NewSignature(0xAA, time.Now(), [6]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06}),
		},
	}
	for i, packet := range cases {
		copy := packet.Copy()
		require.Equal(t, packet, copy, fmt.Sprintf("case %d: %+v", i, packet))
		b1, err := packet.Marshal()
		require.NoError(t, err, fmt.Sprintf("case %d: %+v", i, packet))
		b2, err := copy.Marshal()
		require.NoError(t, err, fmt.Sprintf("case %d: %+v", i, packet))
		require.Equal(t, b1, b2, fmt.Sprintf("case %d: %+v", i, packet))
	}
}
