/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

// signatureTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func signatureTemplate() string {
	var tmpl = "package signature\n" +
		"\n" +
		"import (\n" +
		"    \"fmt\"\n" +
		"    \"time\"\n" +
		")\n" +
		"\n" +
		"// Signature type\n" +
		"type Signature []byte\n" +
		"\n" +
		"const (\n" +
		"    // SIGNATURE_LEN constant\n" +
		"    SIGNATURE_LEN = 13\n" +
		")\n" +
		"\n" +
		"// NewSignature creates signature\n" +
		"func NewSignature(linkID byte, timestamp time.Time, signature [6]byte) Signature {\n" +
		"    s := make([]byte, SIGNATURE_LEN)\n" +
		"    s[0] = linkID\n" +
		"    t := timestamp.Sub(signatureReferenceDate) / (time.Microsecond * 10)\n" +
		"\ts[1] = byte(t)\n" +
		"\ts[2] = byte(t >> 8)\n" +
		"\ts[3] = byte(t >> 16)\n" +
		"\ts[4] = byte(t >> 24)\n" +
		"\ts[5] = byte(t >> 32)\n" +
		"\ts[6] = byte(t >> 40)\n" +
		"    copy(s[7:], signature[:])\n" +
		"    return s\n" +
		"}\n" +
		"\n" +
		"// LinkID returns link id\n" +
		"func (s Signature) LinkID() byte {\n" +
		"    return s[0]\n" +
		"}\n" +
		"\n" +
		"// 1st January 2015 GMT https://mavlink.io/en/guide/message_signing.html#timestamps\n" +
		"var signatureReferenceDate = time.Date(2015, 01, 01, 0, 0, 0, 0, time.UTC)\n" +
		"\n" +
		"// Timestamp returns timestamp of signature\n" +
		"func (s Signature) Timestamp() time.Time {\n" +
		"    return signatureReferenceDate.Add(time.Duration(uint64(s[6])<<40 | uint64(s[5])<<32 | uint64(s[4])<<24 | uint64(s[3])<<16 | uint64(s[2])<<8 | uint64(s[1])) * (10 * time.Microsecond))\n" +
		"}\n" +
		"\n" +
		"// Signature returns signature slice\n" +
		"func (s Signature) Signature() (signature [6]byte) {\n" +
		"    copy(signature[:], s[7:])\n" +
		"    return signature\n" +
		"}\n" +
		"\n" +
		"// String function return string view of Packet struct\n" +
		"func (s Signature) String() string {\n" +
		"\treturn fmt.Sprintf(\n" +
		"\t\t\"&Signature{ linkID: 0x%02X, timestamp: \\\"%+v\\\", signature: \\\"%06X\\\" }\",\n" +
		"    \ts.LinkID(),\n" +
		"    \ts.Timestamp(),\n" +
		"    \ts.Signature(),\n" +
		"    )\n" +
		"}\n" +
		"\n" +
		""
	return tmpl
}
