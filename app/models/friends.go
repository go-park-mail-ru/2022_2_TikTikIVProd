package models

type Friends struct {
	Id1 int `json:"id_user" param:"id_user" validate:"required" gorm:"column:id1"`
	Id2 int `json:"id_friend" param:"id_friend" validate:"required" gorm:"column:id2"`
}

