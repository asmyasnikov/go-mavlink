package mavlink

import (
	"encoding/binary"
	"fmt"
)

const (
	// MSG_ID_PING_MOCK is a define of id of ping message
	MSG_ID_PING_MOCK MessageID = 222
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

// Len
func (m *pingMock) Len() int {
	return 4
}

// Marshal
func (m *pingMock) Marshal() ([]byte, error) {
	payload := make([]byte, 4)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.Seq))
	return payload, nil
}

// Unmarshal
func (m *pingMock) Unmarshal(payload []byte) error {
	m.Seq = uint32(binary.LittleEndian.Uint32(payload[0:]))
	return nil
}
