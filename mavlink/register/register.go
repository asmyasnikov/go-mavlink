/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package register

import (
	"github.com/asmyasnikov/go-mavlink/mavlink/errors"
	"github.com/asmyasnikov/go-mavlink/mavlink/message"
	"github.com/asmyasnikov/go-mavlink/mavlink/packet"
	"strconv"
)

// MessageInfo type
type MessageInfo struct {
	Name        string
	Size        int
	Extra       uint8
	Constructor func(p packet.Packet) (message.Message, error)
}

var supported = make(map[message.MessageID]*MessageInfo)

// Register method provide register dialect message on decoder knowledge
func Register(msgID message.MessageID, msgName string, msgSize int, crcExtra uint8, msgConstructor func(p packet.Packet) (message.Message, error)) {
	if info, ok := supported[msgID]; ok {
		panic("Message with ID = " + strconv.Itoa(int(msgID)) + " already exists. Fix collision '" + msgName + "' vs '" + info.Name + "' and re-run mavgen")
	} else {
		supported[msgID] = &MessageInfo{
			Name:        msgName,
			Size:        msgSize,
			Extra:       crcExtra,
			Constructor: msgConstructor,
		}
	}
}

// Info return info about message
func Info(msgID message.MessageID) (*MessageInfo, error) {
	if info, ok := supported[msgID]; ok {
		return info, nil
	} else {
		return nil, errors.ErrUnknownMsgID
	}
}
