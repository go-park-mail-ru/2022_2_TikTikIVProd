package models

import (
	"time"
)

type Cookie struct {
	SessionToken string    `gorm:"column:value"`
	UserId       int       `gorm:"column:user_id"`
	Expires      time.Time `gorm:"column:expires"`
}
