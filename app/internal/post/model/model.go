package model

import "time"

type Post struct {
	ID         int       	`json:"id"`
	UserID     int       	`json:"user_id"`
	Message    string    	`json:"message"`
	CreateDate time.Time 	`json:"create_date"`
	ImageLinks 	[]string  	`json:"image_links" gorm:"-"`
	UserFirstName  string	`json:"user_first_name"`
	UserLastName   string	`json:"user_last_name"`
}

