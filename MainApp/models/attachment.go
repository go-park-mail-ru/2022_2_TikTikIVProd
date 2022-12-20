package models

const (
	ImageAtt string = "image"
	FileAtt  string = "file"
)

type Attachment struct {
	ID      uint64 `json:"id"`
	Type    string `json:"type"`
	AttLink string `json:"-"`
}
