package models

type Friends struct {
	Id1 int `json:"id_user" gorm:"column:id1"`
	Id2 int `json:"id_friend" gorm:"column:id2"`
}

