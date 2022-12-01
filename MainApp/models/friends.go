package models

type Friends struct {
	Id1 uint64 `gorm:"column:user_id1"`
	Id2 uint64 `param:"friend_id" validate:"required" gorm:"column:user_id2"`
}

