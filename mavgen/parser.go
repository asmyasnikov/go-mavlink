/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

// parserTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func parserTemplate() string {
	var tmpl = "package parser\n" +
		"\n" +
		"import (\n" +
		"    \"{{.CommonPackageURL}}/packet\"\n" +
		")\n" +
		"\n" +
		"// Parser interface is abstract of parsers\n" +
		"type Parser interface {\n" +
		"\tParseChar(c byte) (packet.Packet, error)\n" +
		"    Destroy()\n" +
		"}\n" +
		""
	return tmpl
}
