package models

import (
	"time"
)

type Post struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id" validate:"required"`
	UserFirstName string    `json:"user_first_name"`
	UserLastName  string    `json:"user_last_name"`
	Message       string    `json:"message" validate:"required"`
	CreateDate    time.Time `json:"create_date"`
	Images        []Image   `json:"images"`
}
