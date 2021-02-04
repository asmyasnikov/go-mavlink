package parser

import (
	"github.com/asmyasnikov/go-mavlink/mavlink/mock"
	"github.com/asmyasnikov/go-mavlink/mavlink/packet"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMarshallUnmarshal(t *testing.T) {
	ping := &mock.Ping{
		Seq: 9999,
	}
	payload, err := ping.Marshal()
	require.NoError(t, err)
	cases := []struct{
		packet packet.Packet
		bytes  []byte
	}{
		{
			&packet1{msgID: ping.MsgID(), payload: payload},
			[]byte{0xfe, 0x4, 0x0, 0x0, 0x0, 0xde, 0xf, 0x27, 0x0, 0x0, 0xf4, 0xe2},
		},
		{
			&packet2{msgID: ping.MsgID(), payload: payload},
			[]byte{0xfd, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0xde, 0x0, 0x0, 0xf, 0x27, 0x7, 0xc5},
		},
		{
			&packet2{
				msgID: ping.MsgID(),
				payload: payload,
				incompatFlags: 0x01,
				signature: []byte{0x0e, 0x47, 0x04, 0x0c, 0xef, 0x3b, 0x0e, 0x47, 0x04, 0x0c, 0xef, 0xef, 0x9b},
			},
			[]byte{0xfd, 0x2, 0x1, 0x0, 0x0, 0x0, 0x0, 0xde, 0x0, 0x0, 0xf, 0x27, 0x96, 0x90, 0xe, 0x47, 0x4, 0xc, 0xef, 0x3b, 0xe, 0x47, 0x4, 0xc, 0xef, 0xef, 0x9b},
		},
	}
	for _, c := range cases {
		b, err := c.packet.Marshal()
		require.NoError(t, err)
		require.Equal(t, c.bytes, b)
	}
	for _, c := range cases {
		err := c.packet.Unmarshal(c.bytes)
		require.NoError(t, err)
		msg, err := c.packet.Message()
		require.NoError(t, err)
		p, ok := msg.(*mock.Ping)
		require.True(t, ok)
		require.Equal(t, ping.Seq, p.Seq)
	}
}
