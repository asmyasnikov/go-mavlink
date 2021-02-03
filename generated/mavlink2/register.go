/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package mavlink

import "strconv"

type message struct {
	Name        string
	Size        int
	Extra       uint8
	Constructor func(p *Packet) Message
}

var supported = make(map[MessageID]*message)

// Register method provide register dialect message on decoder knowledge
func Register(msgID MessageID, msgName string, msgSize int, crcExtra uint8, msgConstructor func(p *Packet) Message) {
	if msg, ok := supported[msgID]; ok {
		panic("Message with ID = " + strconv.Itoa(int(msgID)) + " already exists. Fix collision '" + msgName + "' vs '" + msg.Name + "' and re-run mavgen")
	} else {
		supported[msgID] = &message{
			Name:        msgName,
			Size:        msgSize,
			Extra:       crcExtra,
			Constructor: msgConstructor,
		}
	}
}
