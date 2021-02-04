/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package minimal

import (
	"github.com/asmyasnikov/go-mavlink/mavlink/message"
	"github.com/asmyasnikov/go-mavlink/mavlink/packet"
	"github.com/asmyasnikov/go-mavlink/mavlink/register"
)

func init() {
	register.Register(MSG_ID_HEARTBEAT, "MSG_ID_HEARTBEAT", 9, 50, func(p packet.Packet) (message.Message, error) {
		msg := new(Heartbeat)
		if err := msg.Unmarshal(p.Payload()); err != nil {
			return nil, err
		}
		return msg, nil
	})
}
