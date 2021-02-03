package mavlink

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestU16ToBytes(t *testing.T) {
	expected := []byte{0x25, 0x93}
	bytes := u16ToBytes(37669)
	require.Equal(t, len(bytes), 2, "Len of bytes should be equal 2")
	require.Equal(t, expected[0], bytes[0], "First byte unexpected")
	require.Equal(t, expected[1], bytes[1], "Second byte unexpected")
}

func TestMarshallUnmarshal(t *testing.T) {
	ping := &pingMock{
		Seq: 9999,
	}
	payload, err := ping.Marshal()
	require.NoError(t, err)
	cases := []struct{
		packet Packet
		bytes []byte
	}{
		{&packet1{msgID: ping.MsgID(), payload: payload}, []byte{0xfe, 0x4, 0x0, 0x0, 0x0, 0xde, 0xf, 0x27, 0x0, 0x0, 0xf4, 0xe2} },
		{&packet2{msgID: ping.MsgID(), payload: payload}, []byte{0xfd, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0xde, 0x0, 0x0, 0xf, 0x27, 0x7, 0xc5} },
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
		p, ok := msg.(*pingMock)
		require.True(t, ok)
		require.Equal(t, ping.Seq, p.Seq)
	}
}
