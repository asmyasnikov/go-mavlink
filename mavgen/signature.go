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
		"    \"crypto/sha256\"\n" +
		"    \"../helpers\"\n" +
		")\n" +
		"\n" +
		"// Signature type\n" +
		"type Signature struct {\n" +
		"    linkID    byte\n" +
		"    timestamp time.Time\n" +
		"    crc       [6]byte\n" +
		"}\n" +
		"\n" +
		"const (\n" +
		"    // SIGNATURE_LEN constant\n" +
		"    SIGNATURE_LEN = 13\n" +
		")\n" +
		"\n" +
		"// New returns new Signature struct instance\n" +
		"func New(linkID byte, timestamp time.Time, secret [32]byte, packet []byte) (*Signature, error) {\n" +
		"\tif timestamp.Sub(signatureReferenceDate) < 0 {\n" +
		"\t\treturn nil, fmt.Errorf(\"bad timestamp %+v (must be newer than %+v)\", timestamp, signatureReferenceDate)\n" +
		"\t}\n" +
		"\ttimestamp = signatureReferenceDate.Add((timestamp.Sub(signatureReferenceDate) / (time.Microsecond * 10)) * (time.Microsecond * 10))\n" +
		"\ts := &Signature{\n" +
		"\t\tlinkID:    linkID,\n" +
		"\t\ttimestamp: timestamp,\n" +
		"\t}\n" +
		"\tt := uint64(s.timestamp.Sub(signatureReferenceDate) / (time.Microsecond * 10))\n" +
		"\th := sha256.New()\n" +
		"\th.Write(secret[:])\n" +
		"\th.Write(packet[:])\n" +
		"\th.Write([]byte{s.linkID})\n" +
		"\th.Write(helpers.U48ToBytes(t))\n" +
		"\tcopy(s.crc[:], h.Sum(nil)[:6])\n" +
		"\treturn s, nil\n" +
		"}\n" +
		"\n" +
		"// Copy returns copy of Signature struct instance\n" +
		"func (s *Signature) Copy() *Signature {\n" +
		"\tif s == nil {\n" +
		"\t\treturn nil\n" +
		"\t}\n" +
		"\tc := &Signature{\n" +
		"\t\tlinkID:    s.linkID,\n" +
		"\t\ttimestamp: s.timestamp,\n" +
		"\t}\n" +
		"\tcopy(c.crc[:], s.crc[:])\n" +
		"\treturn c\n" +
		"}\n" +
		"\n" +
		"// Unmarshal returns de-serialized bytes to Signature struct\n" +
		"func (s *Signature) Unmarshal(buffer []byte) error {\n" +
		"    s.linkID = buffer[0]\n" +
		"    s.timestamp = signatureReferenceDate.Add(time.Duration(helpers.BytesToU48(buffer[1:7])) * (time.Microsecond * 10))\n" +
		"    copy(s.crc[:], buffer[7:13])\n" +
		"    return nil\n" +
		"}\n" +
		"\n" +
		"// Marshal returns serialized bytes from Signature struct\n" +
		"func (s *Signature) Marshal() ([]byte, error) {\n" +
		"    bytes := make([]byte, 0, SIGNATURE_LEN)\n" +
		"    bytes = append(bytes, byte(s.linkID))\n" +
		"    bytes = append(bytes, helpers.U48ToBytes(uint64(s.timestamp.Sub(signatureReferenceDate) / (time.Microsecond * 10)))...)\n" +
		"    bytes = append(bytes, s.crc[:]...)\n" +
		"    return bytes, nil\n" +
		"}\n" +
		"\n" +
		"// LinkID returns link id\n" +
		"func (s *Signature) LinkID() byte {\n" +
		"    return s.linkID\n" +
		"}\n" +
		"\n" +
		"// 1st January 2015 GMT https://mavlink.io/en/guide/message_signing.html#timestamps\n" +
		"var signatureReferenceDate = time.Date(2015, 01, 01, 0, 0, 0, 0, time.UTC)\n" +
		"\n" +
		"// Timestamp returns timestamp of signature\n" +
		"func (s *Signature) Timestamp() time.Time {\n" +
		"    return s.timestamp\n" +
		"}\n" +
		"\n" +
		"// Signature returns signature slice\n" +
		"func (s *Signature) CRC() (crc [6]byte) {\n" +
		"    copy(crc[:], s.crc[:])\n" +
		"    return crc\n" +
		"}\n" +
		"\n" +
		"// String function return string view of Packet struct\n" +
		"func (s Signature) String() string {\n" +
		"\treturn fmt.Sprintf(\n" +
		"\t\t\"&Signature{ linkID: 0x%02X, timestamp: \\\"%+v\\\", crc: \\\"%06X\\\" }\",\n" +
		"    \ts.LinkID(),\n" +
		"    \ts.Timestamp().UTC(),\n" +
		"    \ts.CRC(),\n" +
		"    )\n" +
		"}\n" +
		"\n" +
		""
	return tmpl
}
