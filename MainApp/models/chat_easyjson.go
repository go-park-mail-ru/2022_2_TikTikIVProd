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

func easyjson9b8f5552DecodeGithubComGoParkMailRu20222TikTikIVProdMainAppModels(in *jlexer.Lexer, out *Message) {
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
		case "dialog_id":
			out.DialogID = uint64(in.Uint64())
		case "sender_id":
			out.SenderID = uint64(in.Uint64())
		case "receiver_id":
			out.ReceiverID = uint64(in.Uint64())
		case "body":
			out.Body = string(in.String())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
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
func easyjson9b8f5552EncodeGithubComGoParkMailRu20222TikTikIVProdMainAppModels(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"dialog_id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.DialogID))
	}
	{
		const prefix string = ",\"sender_id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.SenderID))
	}
	{
		const prefix string = ",\"receiver_id\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.ReceiverID))
	}
	{
		const prefix string = ",\"body\":"
		out.RawString(prefix)
		out.String(string(in.Body))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9b8f5552EncodeGithubComGoParkMailRu20222TikTikIVProdMainAppModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9b8f5552EncodeGithubComGoParkMailRu20222TikTikIVProdMainAppModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9b8f5552DecodeGithubComGoParkMailRu20222TikTikIVProdMainAppModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9b8f5552DecodeGithubComGoParkMailRu20222TikTikIVProdMainAppModels(l, v)
}
func easyjson9b8f5552DecodeGithubComGoParkMailRu20222TikTikIVProdMainAppModels1(in *jlexer.Lexer, out *Dialog) {
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
		case "dialog_id":
			out.Id = uint64(in.Uint64())
		case "UserId1":
			out.UserId1 = uint64(in.Uint64())
		case "UserId2":
			out.UserId2 = uint64(in.Uint64())
		case "messages":
			if in.IsNull() {
				in.Skip()
				out.Messages = nil
			} else {
				in.Delim('[')
				if out.Messages == nil {
					if !in.IsDelim(']') {
						out.Messages = make([]Message, 0, 0)
					} else {
						out.Messages = []Message{}
					}
				} else {
					out.Messages = (out.Messages)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Message
					(v1).UnmarshalEasyJSON(in)
					out.Messages = append(out.Messages, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson9b8f5552EncodeGithubComGoParkMailRu20222TikTikIVProdMainAppModels1(out *jwriter.Writer, in Dialog) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"dialog_id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"UserId1\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.UserId1))
	}
	{
		const prefix string = ",\"UserId2\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.UserId2))
	}
	if len(in.Messages) != 0 {
		const prefix string = ",\"messages\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v2, v3 := range in.Messages {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Dialog) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9b8f5552EncodeGithubComGoParkMailRu20222TikTikIVProdMainAppModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Dialog) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9b8f5552EncodeGithubComGoParkMailRu20222TikTikIVProdMainAppModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Dialog) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9b8f5552DecodeGithubComGoParkMailRu20222TikTikIVProdMainAppModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Dialog) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9b8f5552DecodeGithubComGoParkMailRu20222TikTikIVProdMainAppModels1(l, v)
}