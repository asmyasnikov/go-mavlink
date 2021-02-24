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
		"\t\"math/rand\"\n" +
		")\n" +
		"\n" +
		"// U16ToBytes makes bytearray representation\n" +
		"func U16ToBytes(v uint16) []byte {\n" +
		"\treturn []byte{\n" +
		"\t\tbyte(v),\n" +
		"\t\tbyte(v >> 8),\n" +
		"\t}\n" +
		"}\n" +
		"\n" +
		"// RandomByte generate random byte\n" +
		"func RandomByte() byte {\n" +
		"\treturn RandomBytes(1)[0]\n" +
		"}\n" +
		"\n" +
		"// RandomBytes generate random byte\n" +
		"func RandomBytes(size int) []byte {\n" +
		"\tbuffer := make([]byte, size)\n" +
		"\t_, _ = rand.Read(buffer[:])\n" +
		"\treturn buffer\n" +
		"}\n" +
		"\n" +
		"// BytesToU24 makes uint32 (like as uint24) from bytearray\n" +
		"func BytesToU24(v []byte) uint32 {\n" +
		"\treturn uint32(v[2])<<16 |\n" +
		"\t\tuint32(v[1])<<8 |\n" +
		"\t\tuint32(v[0])\n" +
		"}\n" +
		"\n" +
		"// U24ToBytes makes bytearray from uint32 value (like as uin24)\n" +
		"func U24ToBytes(v uint32) []byte {\n" +
		"\treturn []byte{\n" +
		"\t\tbyte(v),\n" +
		"\t\tbyte(v >> 8),\n" +
		"\t\tbyte(v >> 16),\n" +
		"\t}\n" +
		"}\n" +
		"\n" +
		"// BytesToU48 makes uint64 value (like as uint48)\n" +
		"func BytesToU48(v []byte) uint64 {\n" +
		"\treturn uint64(v[5])<<40 |\n" +
		"\t\tuint64(v[4])<<32 |\n" +
		"\t\tuint64(v[3])<<24 |\n" +
		"\t\tuint64(v[2])<<16 |\n" +
		"\t\tuint64(v[1])<<8 |\n" +
		"\t\tuint64(v[0])\n" +
		"}\n" +
		"\n" +
		"// U48ToBytes makes bytearray from uint64 value (like as uint48)\n" +
		"func U48ToBytes(v uint64) []byte {\n" +
		"\treturn []byte{\n" +
		"\t\tbyte(v),\n" +
		"\t\tbyte(v >> 8),\n" +
		"\t\tbyte(v >> 16),\n" +
		"\t\tbyte(v >> 24),\n" +
		"\t\tbyte(v >> 32),\n" +
		"\t\tbyte(v >> 40),\n" +
		"\t}\n" +
		"}\n" +
		""
	return tmpl
}
