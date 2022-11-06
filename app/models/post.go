package models

import "time"

type Post struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	AvatarID      int       `json:"avatar_id" readonly:"true"`
	UserFirstName string    `json:"user_first_name" readonly:"true"`
	UserLastName  string    `json:"user_last_name" readonly:"true"`
	Message       string    `json:"message" validate:"required"`
	CreateDate    time.Time `json:"create_date" readonly:"true"`
	Images        []Image   `json:"images"`
}
