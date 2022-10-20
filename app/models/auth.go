package models

import (
	"time"
)

type Cookie struct {
	SessionToken string    `json:"session_token" gorm:"column:value"`
	UserId       int       `json:"user_id" gorm:"column:user_id"`
	Expires      time.Time `json:"expires" gorm:"column:expires"`
}
