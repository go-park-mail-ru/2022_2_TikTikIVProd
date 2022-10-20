package models

type User struct {
	Id        int    `json:"id" gorm:"column:id"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	NickName  string `json:"nick_name" gorm:"column:nick_name"`
	Avatar    int    `json:"avatar" gorm:"column:avatar_img_id"`
	Email     string `json:"email" gorm:"column:email"`
	Password  string `json:"password,omitempty" gorm:"column:passhash"`
}

type UserSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

