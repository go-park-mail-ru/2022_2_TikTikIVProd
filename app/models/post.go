package models

import (
	"time"
)

type Post struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	Message    string    `json:"message"`
	CreateDate time.Time `json:"create_date"`
	Images     []Image   `json:"images"`
	FirstName  string
	LastName   string
}
