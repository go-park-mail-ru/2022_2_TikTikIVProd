package models

type Sticker struct {
	ID   uint64 `json:"id" gorm:"column:id"`
	Link string `json:"-" gorm:"column:link"`
}
