/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package minimal

import (
	"../../message"
	"../../register"
)

func init() {
	for msgID, info := range map[message.MessageID]register.MessageInfo{
		MSG_ID_HEARTBEAT: {
			"MSG_ID_HEARTBEAT",
			9,
			50,
			func(bytes []byte) (message.Message, error) {
				msg := new(Heartbeat)
				return msg, msg.Unmarshal(bytes)
			},
		},
	} {
		register.Register(msgID, info.Name, info.Size, info.Extra, info.Constructor)
	}
}
