package mock

import (
	"encoding/binary"
	"fmt"
	"github.com/asmyasnikov/go-mavlink/mavlink/message"
	"github.com/asmyasnikov/go-mavlink/mavlink/register"
)

func init() {
	register.Register(MSG_ID_PING_MOCK, "MSG_ID_PING_MOCK", 4, 0, func(bytes []byte) (message.Message, error) {
		msg := new(Ping)
		if err := msg.Unmarshal(bytes); err != nil {
			return nil, err
		}
		return msg, nil
	})
}

const (
	// MSG_ID_PING_MOCK is a define of id of ping message
	MSG_ID_PING_MOCK message.MessageID = 222
)

// Ping struct
type Ping struct {
	Seq uint32 // Ping sequence
}

// MsgID func
func (m *Ping) MsgID() message.MessageID {
	return MSG_ID_PING_MOCK
}

// Dict func
func (m *Ping) Dict() map[string]interface{} {
	return map[string]interface{}{
		"Ping.Seq": m.Seq,
	}
}

// String func
func (m *Ping) String() string {
	return fmt.Sprintf(
		"&Ping{ Seq: %+v }",
		m.Seq,
	)
}

// Len func
func (m *Ping) Len() int {
	return 4
}

// Marshal func
func (m *Ping) Marshal() ([]byte, error) {
	payload := make([]byte, 4)
	binary.LittleEndian.PutUint32(payload[0:], uint32(m.Seq))
	return payload, nil
}

// Unmarshal func
func (m *Ping) Unmarshal(data []byte) error {
	payload := make([]byte, 4)
	copy(payload[0:], data)
	m.Seq = uint32(binary.LittleEndian.Uint32(payload[0:]))
	return nil
}
