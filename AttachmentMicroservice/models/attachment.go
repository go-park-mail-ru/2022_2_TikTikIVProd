package models

const (
	ImageAtt string = "image"
	FileAtt  string = "file"
)

type Attachment struct {
	ID      uint64 `json:"id"`
	AttLink string `json:"-"`
	Type    string `json:"type"`
}
