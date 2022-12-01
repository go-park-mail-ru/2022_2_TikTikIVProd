package models

type Image struct {
	ID      uint64    `json:"id"`
	ImgLink string `json:"-"`
}
