package models

type File struct {
	ID       uint64 `json:"id"`
	FileLink string `json:"-"`
}
