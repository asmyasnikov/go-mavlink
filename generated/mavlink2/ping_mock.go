package mavlink

import (
	"encoding/binary"
	"fmt"
)

const (
	MSG_ID_PING_MOCK MessageID = 222 // PING MSG ID
)

// pingMock struct
type pingMock struct {
	Seq uint32 // pingMock sequence
}

// MsgID
func (m *pingMock) MsgID() MessageID {
	return MSG_ID_PING_MOCK
}

// String
func (m *pingMock) String() string {
	return fmt.Sprintf(
		"&pingMock{ Seq: %+v }",
		m.Seq,
	)
}

// Pack
func (m *pingMock) Pack(p *Packet) error {
	payload := make([]byte, 4)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.Seq))
	payloadLen := len(payload)
	for payloadLen > 1 && payload[payloadLen-1] == 0 {
		payloadLen--
	}
	payload = payload[:payloadLen]
	p.MsgID = m.MsgID()
	p.Payload = payload
	return nil
}

// Unpack
func (m *pingMock) Unpack(p *Packet) error {
	payload := p.Payload[:]
	if len(p.Payload) < 4 {
		payload = append(payload, ZeroTail[:4-len(p.Payload)]...)
	}
	m.Seq = uint32(binary.LittleEndian.Uint32(payload[0:]))
	return nil
}
