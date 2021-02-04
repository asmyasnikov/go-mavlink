/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

// errorsTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func errorsTemplate() string {
	var tmpl = "package errors\n" +
		"\n" +
		"import (\n" +
		"\t\"errors\"\n" +
		")\n" +
		"\n" +
		"var (\n" +
		"\t// ErrUnknownMsgID define\n" +
		"\tErrUnknownMsgID = errors.New(\"unknown msg id\")\n" +
		"\t// ErrCrcFail define\n" +
		"\tErrCrcFail = errors.New(\"checksum did not match\")\n" +
		"\t// ErrNoNewData define\n" +
		"\tErrNoNewData = errors.New(\"no new data\")\n" +
		"\t// ErrNilPointerReference define\n" +
		"\tErrNilPointerReference = errors.New(\"nil pointer reference\")\n" +
		"    // ErrPayloadTooSmall define\n" +
		"    ErrPayloadTooSmall = errors.New(\"payload too small\")\n" +
		")\n" +
		""
	return tmpl
}
