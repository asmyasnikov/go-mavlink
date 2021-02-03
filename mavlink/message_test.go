package mavlink

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func init() {
	Register(MSG_ID_PING_MOCK, "MSG_ID_PING_MOCK", 4, 0, func(p Packet) (Message, error) {
		msg := new(pingMock)
		if err := msg.Unmarshal(p.Payload()); err != nil {
			return nil, err
		}
		return msg, nil
	})
}

func TestRoundTripChannels(t *testing.T) {
	rand.Seed(43)
	var buffer bytes.Buffer
	dec := NewDecoder(&buffer)
	enc := NewEncoder(&buffer)

	pings := make([]*pingMock, 0, 255)
	processed := make([]*pingMock, 0, 255)

	for i := 0; i < 255; i++ {
		ping := &pingMock{
			Seq: rand.Uint32(),
		}
		pings = append(pings, ping)
		require.NoError(t, enc.Encode(MAVLINK_VERSION(i%2+1), uint8(rand.Uint32()%uint32(^uint8(0))), uint8(rand.Uint32()%uint32(^uint8(0))), ping))
	}
	for {
		packet, err := dec.Decode()
		if err != nil {
			break
		}
		require.Equal(t, packet.MsgID(), MSG_ID_PING_MOCK, "MsgID fail")
		msg, err := packet.Message()
		require.NoError(t, err)
		ping, ok := msg.(*pingMock)
		require.True(t, ok)
		processed = append(processed, ping)
		if len(processed) == 255 {
			break
		}
	}
	require.Equal(t, len(pings), len(processed), "Pings not processed")
	for i, v := range pings {
		require.Equal(t, processed[i].Seq, v.Seq, "Order failed")
	}
}

func TestDecode(t *testing.T) {
	// decode a known good byte stream
	dec := NewDecoder(bytes.NewBuffer([]byte{0xfe, 0x4, 0x0, 0x0, 0x0, 0xde, 0xe8, 0x3, 0x0, 0x0, 0x71, 0x4}))
	packet, err := dec.Decode()
	require.NoError(t, err)
	require.NotNil(t, packet)
}

func TestDecodeTwoMessages(t *testing.T) {
	// decode a known good byte stream
	dec := NewDecoder(bytes.NewBuffer([]byte{0xfe, 0x4, 0x0, 0x0, 0x0, 0xde, 0xe8, 0x3, 0x0, 0x0, 0x71, 0x4,
		0xfe, 0x4, 0x0, 0x0, 0x0, 0xde, 0xf, 0x27, 0x0, 0x0, 0xf4, 0xe2}))
	a, err := dec.Decode()
	require.NoError(t, err)
	require.NotNil(t, a)
	b, err := dec.Decode()
	require.NoError(t, err)
	require.NotNil(t, b)
	require.NotEqual(t, a, b)
}

func TestDecodeFalseStart(t *testing.T) {
	// a false start followed by a known good byte stream
	dec := NewDecoder(bytes.NewBuffer([]byte{0xfe, 0x00, 0xfe, 0xfe, 0x4, 0x0, 0x0, 0x0, 0xde, 0xe8, 0x3, 0x0, 0x0, 0x71, 0x4}))
	packet, err := dec.Decode()
	require.NoError(t, err)
	require.NotNil(t, packet)
}

func TestDecodeMultipleFalseStarts(t *testing.T) {
	// two false starts followed by a known good byte stream
	dec := NewDecoder(bytes.NewBuffer([]byte{0xfe, 0x00, 0xfd, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0xde, 0x0, 0x0, 0xf, 0x27, 0x0, 0x0, 0xc4, 0xd2,
		0xfe, 0xfe, 0x4, 0x0, 0x0, 0x0, 0xde, 0xf, 0x27, 0x0, 0x0, 0xf4, 0xe2}))
	a, err := dec.Decode()
	require.NoError(t, err)
	require.NotNil(t, a)
	b, err := dec.Decode()
	require.NoError(t, err)
	require.NotNil(t, b)
	require.NotEqual(t, a, b)
}
