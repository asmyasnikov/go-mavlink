/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

// registerTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func registerTemplate() string {
	var tmpl = "package register\n" +
		"\n" +
		"import (\n" +
		"    \"strconv\"\n" +
		"    \"{{.CommonPackageURL}}/errors\"\n" +
		"    \"{{.CommonPackageURL}}/message\"\n" +
		")\n" +
		"\n" +
		"// MessageInfo type\n" +
		"type MessageInfo struct {\n" +
		"    Name string\n" +
		"    Size int\n" +
		"    Extra uint8\n" +
		"    Constructor func(bytes []byte) (message.Message, error)\n" +
		"}\n" +
		"\n" +
		"var supported = make(map[message.MessageID]*MessageInfo)\n" +
		"\n" +
		"// Register method provide register dialect message on decoder knowledge\n" +
		"func Register(msgID message.MessageID, msgName string, msgSize int, crcExtra uint8, msgConstructor func(bytes []byte) (message.Message, error)) {\n" +
		"\tif info, ok := supported[msgID]; ok {\n" +
		"\t\tpanic(\"Message with ID = \" + strconv.Itoa(int(msgID)) + \" already exists. Fix collision '\" + msgName + \"' vs '\" + info.Name + \"' and re-run mavgen\")\n" +
		"\t} else {\n" +
		"\t\tsupported[msgID] = &MessageInfo{\n" +
		"\t\t    Name:        msgName,\n" +
		"\t\t    Size:        msgSize,\n" +
		"\t\t    Extra:       crcExtra,\n" +
		"\t\t    Constructor: msgConstructor,\n" +
		"\t\t}\n" +
		"\t}\n" +
		"}\n" +
		"\n" +
		"// Info return info about message\n" +
		"func Info(msgID message.MessageID) (*MessageInfo, error) {\n" +
		"\tif info, ok := supported[msgID]; ok {\n" +
		"\t    return info, nil\n" +
		"\t}\n" +
		"    return nil, errors.ErrUnknownMsgID\n" +
		"}"
	return tmpl
}
