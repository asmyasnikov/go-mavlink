// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package minimal

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonFa90ddaeDecodeGithubComAsmyasnikovGoMavlinkMavlinkDialectsMinimal(in *jlexer.Lexer, out *Heartbeat) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "CustomMode":
			out.CustomMode = uint32(in.Uint32())
		case "Type":
			out.Type = MAV_TYPE(in.Int())
		case "Autopilot":
			out.Autopilot = MAV_AUTOPILOT(in.Int())
		case "BaseMode":
			out.BaseMode = MAV_MODE_FLAG(in.Int())
		case "SystemStatus":
			out.SystemStatus = MAV_STATE(in.Int())
		case "MavlinkVersion":
			out.MavlinkVersion = uint8(in.Uint8())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonFa90ddaeEncodeGithubComAsmyasnikovGoMavlinkMavlinkDialectsMinimal(out *jwriter.Writer, in Heartbeat) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"CustomMode\":"
		out.RawString(prefix[1:])
		out.Uint32(uint32(in.CustomMode))
	}
	{
		const prefix string = ",\"Type\":"
		out.RawString(prefix)
		out.Int(int(in.Type))
	}
	{
		const prefix string = ",\"Autopilot\":"
		out.RawString(prefix)
		out.Int(int(in.Autopilot))
	}
	{
		const prefix string = ",\"BaseMode\":"
		out.RawString(prefix)
		out.Int(int(in.BaseMode))
	}
	{
		const prefix string = ",\"SystemStatus\":"
		out.RawString(prefix)
		out.Int(int(in.SystemStatus))
	}
	{
		const prefix string = ",\"MavlinkVersion\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.MavlinkVersion))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Heartbeat) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFa90ddaeEncodeGithubComAsmyasnikovGoMavlinkMavlinkDialectsMinimal(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Heartbeat) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFa90ddaeEncodeGithubComAsmyasnikovGoMavlinkMavlinkDialectsMinimal(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Heartbeat) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFa90ddaeDecodeGithubComAsmyasnikovGoMavlinkMavlinkDialectsMinimal(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Heartbeat) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFa90ddaeDecodeGithubComAsmyasnikovGoMavlinkMavlinkDialectsMinimal(l, v)
}