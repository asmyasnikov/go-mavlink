/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

// helpersTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func helpersTemplate() string {
	var tmpl = "package helpers\n" +
		"\n" +
		"import (\n" +
		"    \"math/rand\"\n" +
		")\n" +
		"\n" +
		"// U16ToBytes creates bytearray representation\n" +
		"func U16ToBytes(v uint16) []byte {\n" +
		"\treturn []byte{byte(v & 0xff), byte(v >> 8)}\n" +
		"}\n" +
		"\n" +
		"// RandomByte generate random byte\n" +
		"func RandomByte() byte {\n" +
		"    return RandomBytes(1)[0]\n" +
		"}\n" +
		"\n" +
		"// RandomBytes generate random byte\n" +
		"func RandomBytes(size int) []byte {\n" +
		"\tbuffer := make([]byte, size)\n" +
		"\t_, _ = rand.Read(buffer[:])\n" +
		"\treturn buffer\n" +
		"}"
	return tmpl
}
