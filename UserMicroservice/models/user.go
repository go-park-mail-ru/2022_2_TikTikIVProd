package models

import "time"

type User struct {
	Id        uint64    `gorm:"column:id"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	NickName  string    `gorm:"column:nick_name"`
	Avatar    uint64    `gorm:"column:avatar_att_id"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

type Friends struct {
	Id1 uint64 `gorm:"column:user_id1"`
	Id2 uint64 `gorm:"column:user_id2"`
}
