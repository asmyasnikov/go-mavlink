/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

// messageTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func messageTemplate() string {
	var tmpl = "package message\n" +
		"\n" +
		"// MessageID typedef\n" +
		"type MessageID uint32\n" +
		"\n" +
		"// Message is a basic type for encoding/decoding mavlink messages.\n" +
		"// use the Pack() and Unpack() routines on specific message\n" +
		"// types to convert them to/from the Packet type.\n" +
		"type Message interface {\n" +
		"    // MsgID returns message id\n" +
		"\tMsgID() MessageID\n" +
		"    // String returns human-readable representation on Message\n" +
		"\tString() string\n" +
		"\t// Dict returns key-value dictionary of Message internal fields\n" +
		"\tDict() map[string]interface{}\n" +
		"    // Marshal encodes Packet to byte slice\n" +
		"    Marshal() ([]byte, error)\n" +
		"    // Unmarshal parses PAYLOAD and stores the result in Packet\n" +
		"    Unmarshal(payload []byte) error\n" +
		"}\n" +
		""
	return tmpl
}
