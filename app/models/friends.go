package models

type Friends struct {
	Id1 int `gorm:"column:id1"`
	Id2 int `param:"friend_id" validate:"required" gorm:"column:id2"`
}

