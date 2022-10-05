package model

import "time"

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	NickName  string `json:"nick_name"`
	Avatar    int    `json:"avatar"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserSignIn struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Cookie struct {
	SessionToken string    `json:"session_token"`
	UserId       int       `json:"user_id"`
	Expires      time.Time `json:"expires"`
}
