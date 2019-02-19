package mavlink

import (
	"log"
	"reflect"
	"testing"
	"time"
)

func TestRoundTripChannels(t *testing.T) {

	var dialects = DialectSlice{DialectCommon, DialectArdupilotmega}

	var data = make(chan []byte, 256)

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

	dec := NewChannelDecoder(data)

	for i, _ := range cases {
		seq := cases[i].seq
		log.Println("Make packet for seq", seq)
		p := CommonPing{
			Seq: seq,
		}

		var packet Packet
		if err := p.Pack(&packet); err != nil {
			t.Errorf("Pack fail %q (%q)", packet, err)
		}
		packet.fixChecksum(dialects)

		data <- packet.Bytes()

		select {
		case packet := <- dec.decoded :
			log.Println("Receive packet ", packet.Bytes())
			if packet.MsgID != MSG_ID_PING {
				t.Errorf("MsgID fail, want %d, got %d", MSG_ID_PING, packet.MsgID)
			}

			var pingOut CommonPing
			if err := pingOut.Unpack(packet); err != nil {
				t.Errorf("Unpack fail %s", err)
			}

			if seq != pingOut.Seq {
				log.Printf("Erase cases item with seq = %d", pingOut.Seq)
				cases = append(cases[:i], cases[i+1:]...)
				break
			}
		case <-time.After(time.Millisecond) :
			t.Error("Timeout elapsed")
		}
	}
}

func TestDecode(t *testing.T) {
	// decode a known good byte stream
	var data = make(chan []byte)
	dec := NewChannelDecoder(data)
	data <- []byte{0xFD, 0x2B, 0x00, 0x00, 0xC5, 0x02, 0xFC, 0x64, 0x00, 0x00, 0x16, 0x89, 0xB6, 0x44, 0xA5, 0x47, 0x00, 0x00, 0xAB, 0x47, 0x00, 0x00, 0xAA, 0x47, 0x00, 0x00, 0xBF, 0x04, 0x00, 0x00, 0xA4, 0xBD, 0x02, 0x31, 0x2E, 0x30, 0x30, 0x00, 0x05, 0x11, 0x12, 0x0A, 0x01, 0x00, 0x64, 0x00, 0x00, 0x03, 0x07, 0x01, 0x00, 0x00, 0x02, 0x4A, 0x26}
	select {
	case packet := <- dec.decoded :
		log.Println("Decode packet", packet)
	case <-time.After(time.Millisecond) :
		t.Error("Decode fail")
	}
}

func TestDecodeTwoMessages(t *testing.T) {
	// decode a known good byte stream
	var data = make(chan []byte)
	dec := NewChannelDecoder(data)
	data <- []byte{0xFD, 0x0D, 0x00, 0x00, 0x8A, 0x02, 0xFC, 0x04, 0x00, 0x00, 0xE5, 0xF2, 0x41, 0x21, 0x09, 0x7E, 0x3B, 0x16, 0x00, 0x00, 0x01, 0x00, 0x6B, 0x4A, 0x83,
		0xFD, 0x2B, 0x00, 0x00, 0xC5, 0x02, 0xFC, 0x64, 0x00, 0x00, 0x16, 0x89, 0xB6, 0x44, 0xA5, 0x47, 0x00, 0x00, 0xAB, 0x47, 0x00, 0x00, 0xAA, 0x47, 0x00, 0x00, 0xBF, 0x04, 0x00, 0x00, 0xA4, 0xBD, 0x02, 0x31, 0x2E, 0x30, 0x30, 0x00, 0x05, 0x11, 0x12, 0x0A, 0x01, 0x00, 0x64, 0x00, 0x00, 0x03, 0x07, 0x01, 0x00, 0x00, 0x02, 0x4A, 0x26}
	var msg1 *Packet
	select {
	case packet := <- dec.decoded :
		log.Println("Decode packet", packet)
		msg1 = packet
	case <-time.After(time.Millisecond) :
		t.Error("Decode fail")
	}
	var msg2 *Packet
	select {
	case packet := <- dec.decoded :
		log.Println("Decode packet", packet)
		msg2 = packet
	case <-time.After(time.Millisecond) :
		t.Error("Decode fail")
	}
	if reflect.DeepEqual(*msg1, *msg2) {
		t.Error("Messages should not match")
	}
}

func TestDecodeFalseStart(t *testing.T) {
	// a false start followed by a known good byte stream
	var data = make(chan []byte)
	dec := NewChannelDecoder(data)
	data <- []byte{0xFD, 0x0D, 0xFD, 0x0D, 0x00, 0x00, 0x8A, 0x02, 0xFC, 0x04, 0x00, 0x00, 0xE5, 0xF2, 0x41, 0x21, 0x09, 0x7E, 0x3B, 0x16, 0x00, 0x00, 0x01, 0x00, 0x6B, 0x4A, 0x83}
	select {
	case packet := <- dec.decoded :
		log.Println("Decode packet", packet)
	case <-time.After(time.Millisecond) :
		t.Error("Decode fail")
	}
}

func TestDecodeMultipleFalseStarts(t *testing.T) {
	// two false starts followed by a known good byte stream
	var data = make(chan []byte)
	dec := NewChannelDecoder(data)
	data <- []byte{0xFD, 0x0D, 0xFD, 0x0D, 0xFD, 0x0D, 0xFD, 0x0D, 0xFD, 0x0D, 0x00, 0x00, 0x8A, 0x02, 0xFC, 0x04, 0x00, 0x00, 0xE5, 0xF2, 0x41, 0x21, 0x09, 0x7E, 0x3B, 0x16, 0x00, 0x00, 0x01, 0x00, 0x6B, 0x4A, 0x83,
		0x00, 0xFD, 0x2B, 0xFD, 0x2B, 0x00, 0x00, 0xC5, 0x02, 0xFC, 0x64, 0x00, 0x00, 0x16, 0x89, 0xB6, 0x44, 0xA5, 0x47, 0x00, 0x00, 0xAB, 0x47, 0x00, 0x00, 0xAA, 0x47, 0x00, 0x00, 0xBF, 0x04, 0x00, 0x00, 0xA4, 0xBD, 0x02, 0x31, 0x2E, 0x30, 0x30, 0x00, 0x05, 0x11, 0x12, 0x0A, 0x01, 0x00, 0x64, 0x00, 0x00, 0x03, 0x07, 0x01, 0x00, 0x00, 0x02, 0x4A, 0x26}
	var msg1 *Packet
	select {
	case packet := <- dec.decoded :
		log.Println("Decode packet", packet)
		msg1 = packet
	case <-time.After(time.Millisecond) :
		t.Error("Decode fail")
	}
	var msg2 *Packet
	select {
	case packet := <- dec.decoded :
		log.Println("Decode packet", packet)
		msg2 = packet
	case <-time.After(time.Millisecond) :
		t.Error("Decode fail")
	}
	if reflect.DeepEqual(*msg1, *msg2) {
		t.Error("Messages should not match")
	}
}

//func TestDialects(t *testing.T) {
//	var dialects = DialectSlice{DialectCommon, DialectArdupilotmega}
//
//	var data = make(chan []byte, 256)
//
//	dec := NewChannelDecoder(data)
//
//	// try to encode an ardupilot msg before we've added that dialect,
//	// ensure it fails as expected
//	mi := &ArdupilotmegaMeminfo{
//		Brkval:  1000,
//		Freemem: 10,
//	}
//
//	err := mi.Encode(0x1, 0x1, mi)
//	if err != ErrUnknownMsgID {
//		t.Errorf("encode expected ErrUnknownMsgID, got %q", err)
//	}
//
//	buf.Reset()
//
//	if err = enc.Encode(0x1, 0x1, mi); err != nil {
//		t.Errorf("Encode fail %q", err)
//	}
//
//	log.Println("TRY DECODE ARDUPILOT MESSAGE")
//	_, err = NewDecoder(&buf).Decode()
//	log.Println("ARDUPILOT MESSAGE DECODED")
//	if err != ErrUnknownMsgID {
//		t.Errorf("decode expected ErrUnknownMsgID, got %q", err)
//	}
//
//
//	// add the dialect, and ensure it succeeds
//	dec.Dialects.Add(DialectArdupilotmega)
//	// re-encode the msg, and decode it again after adding the required dialect
//	if err = enc.Encode(0x1, 0x1, mi); err != nil {
//		t.Errorf("Encode fail %q", err)
//	}
//
//	pktOut, err := dec.Decode()
//	if err != nil {
//		t.Errorf("Decode fail %q", err)
//	}
//
//	// make sure the output matches our original input for good measure
//	var miOut ArdupilotmegaMeminfo
//	if err := miOut.Unpack(pktOut); err != nil {
//		t.Errorf("Unpack fail %q", err)
//	}
//
//	if miOut.Brkval != mi.Brkval || miOut.Freemem != mi.Freemem {
//		t.Errorf("Round trip fail")
//	}
//}
