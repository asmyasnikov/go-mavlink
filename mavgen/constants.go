/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

// constantsTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func constantsTemplate() string {
	var tmpl = "package mavlink\n" +
		"\n" +
		"import (\n" +
		"\t\"errors\"\n" +
		")\n" +
		"\n" +
		"type MAVLINK_VERSION int\n" +
		"\n" +
		"const (\n" +
		"    MAVLINK_V1 MAVLINK_VERSION = 1\n" +
		"    MAVLINK_V2 MAVLINK_VERSION = 2\n" +
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
		"    // zeroTail is a cache of zero slice for auto append tail to\n" +
		"    // payload in Mavlink2 messages with trimmed payload (variable length)\n" +
		"\tzeroTail           = make([]byte, 256)\n" +
		")\n" +
		""
	return tmpl
}
