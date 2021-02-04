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
		"func U16ToBytes(v uint16) []byte {\n" +
		"\treturn []byte{byte(v & 0xff), byte(v >> 8)}\n" +
		"}\n" +
		""
	return tmpl
}
