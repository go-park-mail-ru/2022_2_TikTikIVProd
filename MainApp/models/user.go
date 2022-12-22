package models

import "time"

type User struct {
	Id        uint64    `json:"id" readonly:"true" gorm:"column:id"`
	FirstName string    `json:"first_name" validate:"required" gorm:"column:first_name"`
	LastName  string    `json:"last_name" validate:"required" gorm:"column:last_name"`
	NickName  string    `json:"nick_name" validate:"required" gorm:"column:nick_name"`
	Avatar    uint64    `json:"avatar" gorm:"column:avatar_att_id"`
	Email     string    `json:"email" validate:"required" gorm:"column:email"`
	Password  string    `json:"password,omitempty" validate:"required" gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

type UserSignIn struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
