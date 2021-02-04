/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

// versionTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func versionTemplate() string {
	var tmpl = "package version\n" +
		"\n" +
		"type MAVLINK_VERSION int\n" +
		"\n" +
		"const (\n" +
		"    MAVLINK_V1 MAVLINK_VERSION = 1\n" +
		"    MAVLINK_V2 MAVLINK_VERSION = 2\n" +
		")\n" +
		""
	return tmpl
}
