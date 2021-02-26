package parser

import (
	"../helpers"
	"../mock"
	"../packet"
	"../signature"
	"fmt"
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
			&packet1{seqID: 22, compID: 14, sysID: 122, msgID: msg.MsgID(), payload: payload},
			[]byte{0xfe, 0x4, 0x16, 0x7a, 0xe, 0xde, 0xf, 0x27, 0x0, 0x0, 0xf4, 0x91},
		},
		{
			&packet2{seqID: 14, compID: 22, sysID: 22, msgID: msg.MsgID(), payload: payload},
			[]byte{0xfd, 0x2, 0x0, 0x0, 0xe, 0x16, 0x16, 0xde, 0x0, 0x0, 0xf, 0x27, 0x1f, 0xb5},
		},
		{
			&packet2{seqID: 122, compID: 122, sysID: 14, msgID: msg.MsgID(), payload: payload},
			[]byte{0xfd, 0x2, 0x0, 0x0, 0x7a, 0xe, 0x7a, 0xde, 0x0, 0x0, 0xf, 0x27, 0xef, 0x7},
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
	linkId := byte(0x0e)
	timestamp := time.Date(2019, 12, 17, 9, 49, 24, 147900, time.UTC)
	msg := &mock.Ping{
		Seq: 9999,
	}
	payload, err := msg.Marshal()
	require.NoError(t, err)
	cases := []struct {
		packet    *packet2
		bytes     []byte
		signature *signature.Signature
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
				signature: func() *signature.Signature {
					s := &signature.Signature{}
					require.NoError(t, s.Unmarshal([]byte{0xe, 0x8e, 0xfe, 0xb, 0xef, 0x3b, 0xe, 0xd3, 0x6c, 0x3b, 0xcb, 0x7, 0xdd}));
					return s
				}(),
			},
			[]byte{0xfd, 0x2, 0x1, 0x0, 0x0, 0x0, 0x0, 0xde, 0x0, 0x0, 0xf, 0x27, 0x96, 0x90, 0xe, 0x8e, 0xfe, 0xb, 0xef, 0x3b, 0xe, 0xd3, 0x6c, 0x3b, 0xcb, 0x7, 0xdd},
			func() *signature.Signature {
				s := &signature.Signature{}
				require.NoError(t, s.Unmarshal([]byte{0xe, 0x8e, 0xfe, 0xb, 0xef, 0x3b, 0xe, 0xd3, 0x6c, 0x3b, 0xcb, 0x7, 0xdd}));
				return s
			}(),
		},
	}
	for i, c := range cases {
		b, err := c.packet.Copy().MarshalWithSignature(linkId, timestamp, [32]byte{})
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
			signature:     func() *signature.Signature {
				s := &signature.Signature{}
				require.NoError(t, s.Unmarshal([]byte{0xe, 0x8e, 0xfe, 0xb, 0xef, 0x3b, 0xe, 0xd3, 0x6c, 0x3b, 0xcb, 0x7, 0xdd}));
				return s
			}(),
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

func TestPacketMarshalUnmarshalSignature(t *testing.T) {
	linkId := helpers.RandomByte()
	timestamp := time.Now()
	randBytes := helpers.RandomBytes(32)
	secretKey := [32]byte{}
	copy(secretKey[:], randBytes)
	msg := &mock.Ping{
		Seq: 9999,
	}
	payload, err := msg.Marshal()
	require.NoError(t, err)
	for i := 0; i < 256; i++ {
		u := &packet2{
			incompatFlags: 0x01,
			sysID: helpers.RandomByte(),
			compID: helpers.RandomByte(),
			seqID: helpers.RandomByte(),
			msgID: msg.MsgID(),
			payload: payload,
		}
		b, err := u.MarshalWithSignature(linkId, timestamp, secretKey)
		require.NoError(t, err, fmt.Sprintf("case %d: { packet: %+v, bytes: %+v }", i, u, b))
		v := &packet2{}
		err = v.Unmarshal(b)
		require.NoError(t, err, fmt.Sprintf("case %d: { packet: %+v, bytes: %+v }", i, v, b))
		fmt.Println(u)
		fmt.Println(v)
		require.Equal(t, u, v, fmt.Sprintf("case %d: { %+v != %+v }", i, v, u))
	}
}
