package models

type Image struct {
	ID      int    `json:"id"`
	ImgLink string `json:"-"`
}
