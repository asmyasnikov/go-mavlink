/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

// packetTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func packetTemplate() string {
	var tmpl = "package packet\n" +
		"\n" +
		"import (\n" +
		"    \"time\"\n" +
		"    \"{{.CommonPackageURL}}/version\"\n" +
		"    \"{{.CommonPackageURL}}/message\"\n" +
		"    \"{{.CommonPackageURL}}/signature\"\n" +
		")\n" +
		"\n" +
		"// MAGIC_NUMBER type\n" +
		"type MAGIC_NUMBER uint8\n" +
		"\n" +
		"// Packet is the interface implemented by frames of every supported version.\n" +
		"type Packet interface {\n" +
		"\t// Version returns version of packet framing\n" +
		"\tVersion() version.MAVLINK_VERSION\n" +
		"\t// IsNil returns true if packet is nil\n" +
		"\tIsNil() bool\n" +
		"\t// IsSigned checks whether the frame contains a signature. It does not validate the signature\n" +
		"\tIsSigned() bool\n" +
		"\t// SysID returns system id\n" +
		"\tSysID() uint8\n" +
		"\t// CompID returns component id\n" +
		"\tCompID() uint8\n" +
		"\t// MsgID returns message id\n" +
		"\tMsgID() message.MessageID\n" +
		"\t// Checksum returns packet checksum\n" +
		"\tChecksum() uint16\n" +
		"\t// SeqID returns packet sequence number\n" +
		"\tSeqID() uint8\n" +
		"\t// Payload returns packet payload\n" +
		"\tPayload() []byte\n" +
		"\t// Signature returns packet signature\n" +
		"\tSignature() *signature.Signature\n" +
		"\t// Copy returns deep copy of packet\n" +
		"\tCopy() Packet\n" +
		"\t// Message returns dialect message\n" +
		"\tMessage() (message.Message, error)\n" +
		"\t// String returns string representation of packet\n" +
		"\tString() string\n" +
		"    // Marshal encodes Packet to byte slice\n" +
		"    Marshal() ([]byte, error)\n" +
		"    // MarshalWithSignature encodes Packet to byte slice including signature at appendix\n" +
		"    MarshalWithSignature(linkID byte, timestamp time.Time, secretKey [32]byte) ([]byte, error)\n" +
		"    // Unmarshal parses PAYLOAD and stores the result in Packet\n" +
		"    Unmarshal(payload []byte) error\n" +
		"}\n" +
		""
	return tmpl
}
