// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjsonA68a6153DecodeGithubComGoParkMailRu20222TikTikIVProdMainAppModels(in *jlexer.Lexer, out *Sticker) {
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
		case "id":
			out.ID = uint64(in.Uint64())
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
func easyjsonA68a6153EncodeGithubComGoParkMailRu20222TikTikIVProdMainAppModels(out *jwriter.Writer, in Sticker) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Sticker) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA68a6153EncodeGithubComGoParkMailRu20222TikTikIVProdMainAppModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Sticker) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA68a6153EncodeGithubComGoParkMailRu20222TikTikIVProdMainAppModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Sticker) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA68a6153DecodeGithubComGoParkMailRu20222TikTikIVProdMainAppModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Sticker) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA68a6153DecodeGithubComGoParkMailRu20222TikTikIVProdMainAppModels(l, v)
}
