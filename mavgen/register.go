/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

// registerTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func registerTemplate() string {
	var tmpl = "package mavlink\n" +
		"\n" +
		"import \"strconv\"\n" +
		"\n" +
		"type message struct {\n" +
		"    Name string\n" +
		"    Size int\n" +
		"    Extra uint8\n" +
		"    Constructor func(p Packet) (Message, error)\n" +
		"}\n" +
		"\n" +
		"var supported = make(map[MessageID]*message)\n" +
		"\n" +
		"// Register method provide register dialect message on decoder knowledge\n" +
		"func Register(msgID MessageID, msgName string, msgSize int, crcExtra uint8, msgConstructor func(p Packet) (Message, error)) {\n" +
		"\tif msg, ok := supported[msgID]; ok {\n" +
		"\t\tpanic(\"Message with ID = \" + strconv.Itoa(int(msgID)) + \" already exists. Fix collision '\" + msgName + \"' vs '\" + msg.Name + \"' and re-run mavgen\")\n" +
		"\t} else {\n" +
		"\t\tsupported[msgID] = &message{\n" +
		"\t\t    Name:        msgName,\n" +
		"\t\t    Size:        msgSize,\n" +
		"\t\t    Extra:       crcExtra,\n" +
		"\t\t    Constructor: msgConstructor,\n" +
		"\t\t}\n" +
		"\t}\n" +
		"}\n" +
		""
	return tmpl
}
