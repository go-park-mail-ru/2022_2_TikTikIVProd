package models

type Friends struct {
	Id1 uint64 `gorm:"column:id1"`
	Id2 uint64 `param:"friend_id" validate:"required" gorm:"column:id2"`
}

