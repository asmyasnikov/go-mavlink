package mavlink

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
	"time"
)

func TestRoundTripChannels(t *testing.T) {

	AddDialect(DialectArdupilotmega)

	cases := []struct{ seq uint32 }{
		{12345},
		{54321},
		{65432},
		{23456},
		{65456},
		{23423},
		{65444},
		{23456},
		{76543},
	}

	dec := NewChannelDecoder()
	defer dec.Stop()

	for i, _ := range cases {
		seq := cases[i].seq
		p := CommonPing{
			Seq: seq,
		}

		packet := &Packet{}
		require.Nil(t, p.Pack(packet), "Pack fail")
		packet.fixChecksum(dialects)

		dec.PushData(packet.Bytes())
	}

	sort.Slice(cases, func(i, j int) bool {
		return cases[i].seq < cases[j].seq
	})

	for packet := dec.NextPacket(time.Millisecond); packet != nil && len(cases) > 0; packet = dec.NextPacket(time.Millisecond) {
		require.Equal(t, packet.MsgID, MSG_ID_PING, "MsgID fail")
		pingOut := &CommonPing{}
		require.Nil(t, pingOut.Unpack(packet), "Unpack fail")
		i := sort.Search(len(cases), func(i int) bool {
			return cases[i].seq >= pingOut.Seq
		})
		require.NotEqual(t, i, len(cases), "Search case by seq = %d fail. Cases: %+v", pingOut.Seq, cases)
		require.Equal(t, cases[i].seq, pingOut.Seq, "Seq fail")
		cases = append(cases[:i], cases[i+1:]...)
	}
	require.Equal(t, 0, len(cases), "Cases not cleared")
}

func TestDecode(t *testing.T) {
	// decode a known good byte stream
	dec := NewChannelDecoder()
	defer dec.Stop()
	dec.PushData([]byte{0xFD, 0x2B, 0x00, 0x00, 0xC5, 0x02, 0xFC, 0x64, 0x00, 0x00, 0x16, 0x89, 0xB6, 0x44, 0xA5, 0x47, 0x00, 0x00, 0xAB, 0x47, 0x00, 0x00, 0xAA, 0x47, 0x00, 0x00, 0xBF, 0x04, 0x00, 0x00, 0xA4, 0xBD, 0x02, 0x31, 0x2E, 0x30, 0x30, 0x00, 0x05, 0x11, 0x12, 0x0A, 0x01, 0x00, 0x64, 0x00, 0x00, 0x03, 0x07, 0x01, 0x00, 0x00, 0x02, 0x4A, 0x26})
	require.NotNil(t, dec.NextPacket(time.Millisecond), "Decode fail")
}

func TestDecodeTwoMessages(t *testing.T) {
	// decode a known good byte stream
	dec := NewChannelDecoder()
	defer dec.Stop()
	dec.PushData([]byte{0xFD, 0x0D, 0x00, 0x00, 0x8A, 0x02, 0xFC, 0x04, 0x00, 0x00, 0xE5, 0xF2, 0x41, 0x21, 0x09, 0x7E, 0x3B, 0x16, 0x00, 0x00, 0x01, 0x00, 0x6B, 0x4A, 0x83,
		0xFD, 0x2B, 0x00, 0x00, 0xC5, 0x02, 0xFC, 0x64, 0x00, 0x00, 0x16, 0x89, 0xB6, 0x44, 0xA5, 0x47, 0x00, 0x00, 0xAB, 0x47, 0x00, 0x00, 0xAA, 0x47, 0x00, 0x00, 0xBF, 0x04, 0x00, 0x00, 0xA4, 0xBD, 0x02, 0x31, 0x2E, 0x30, 0x30, 0x00, 0x05, 0x11, 0x12, 0x0A, 0x01, 0x00, 0x64, 0x00, 0x00, 0x03, 0x07, 0x01, 0x00, 0x00, 0x02, 0x4A, 0x26})
	msg1 := dec.NextPacket(time.Millisecond)
	require.NotNil(t, msg1, "Decode fail")
	msg2 := dec.NextPacket(time.Millisecond)
	require.NotNil(t, msg2, "Decode fail")
	require.NotEqual(t, *msg1, *msg2, "Messages should not match")
}

func TestDecodeFalseStart(t *testing.T) {
	// a false start followed by a known good byte stream
	dec := NewChannelDecoder()
	defer dec.Stop()
	dec.PushData([]byte{0xFD, 0x0D, 0xFD, 0x0D, 0x00, 0x00, 0x8A, 0x02, 0xFC, 0x04, 0x00, 0x00, 0xE5, 0xF2, 0x41, 0x21, 0x09, 0x7E, 0x3B, 0x16, 0x00, 0x00, 0x01, 0x00, 0x6B, 0x4A, 0x83})
	require.NotNil(t, dec.NextPacket(time.Millisecond), "Decode fail")
}

func TestDecodeMultipleFalseStarts(t *testing.T) {
	// two false starts followed by a known good byte stream
	dec := NewChannelDecoder()
	defer dec.Stop()
	dec.PushData([]byte{0xFD, 0x0D, 0xFD, 0x0D, 0xFD, 0x0D, 0xFD, 0x0D, 0xFD, 0x0D, 0x00, 0x00, 0x8A, 0x02, 0xFC, 0x04, 0x00, 0x00, 0xE5, 0xF2, 0x41, 0x21, 0x09, 0x7E, 0x3B, 0x16, 0x00, 0x00, 0x01, 0x00, 0x6B, 0x4A, 0x83,
		0x00, 0xFD, 0x2B, 0xFD, 0x2B, 0x00, 0x00, 0xC5, 0x02, 0xFC, 0x64, 0x00, 0x00, 0x16, 0x89, 0xB6, 0x44, 0xA5, 0x47, 0x00, 0x00, 0xAB, 0x47, 0x00, 0x00, 0xAA, 0x47, 0x00, 0x00, 0xBF, 0x04, 0x00, 0x00, 0xA4, 0xBD, 0x02, 0x31, 0x2E, 0x30, 0x30, 0x00, 0x05, 0x11, 0x12, 0x0A, 0x01, 0x00, 0x64, 0x00, 0x00, 0x03, 0x07, 0x01, 0x00, 0x00, 0x02, 0x4A, 0x26})
	msg1 := dec.NextPacket(time.Millisecond)
	require.NotNil(t, msg1, "Decode fail")
	msg2 := dec.NextPacket(time.Millisecond)
	require.NotNil(t, msg2, "Decode fail")
	require.NotEqual(t, *msg1, *msg2, "Messages should not match")
}

func TestDialects(t *testing.T) {
	dec := NewChannelDecoder()
	defer dec.Stop()

	// try to encode an ardupilot msg before we've added that dialect,
	// ensure it fails as expected
	mi := &ArdupilotmegaMeminfo{
		Brkval:  1000,
		Freemem: 10,
	}

	RemoveDialect(DialectArdupilotmega)

	packet := &Packet{}
	require.Equal(t, packet.Encode(0x1, 0x1, mi), ErrUnknownMsgID, "encode expected ErrUnknownMsgID")

	// add the dialect, and ensure it succeeds
	AddDialect(DialectArdupilotmega)
	require.Nil(t, packet.Encode(0x1, 0x1, mi), "encode unexpected err")

	dec.PushData(packet.Bytes())
	packet = dec.NextPacket(time.Millisecond)
	require.NotNil(t, packet, "Decode fail")
	var miOut ArdupilotmegaMeminfo
	require.Nil(t, miOut.Unpack(packet), "Unpack fail")
	require.Equal(t, miOut.Brkval, mi.Brkval, "Round trip fail")
	require.Equal(t, miOut.Freemem, mi.Freemem, "Round trip fail")
}
