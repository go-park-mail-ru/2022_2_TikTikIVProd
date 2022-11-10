package models

type User struct {
	Id        int    `json:"id" readonly:"true" gorm:"column:id"`
	FirstName string `json:"first_name" validate:"required" gorm:"column:first_name"`
	LastName  string `json:"last_name" validate:"required" gorm:"column:last_name"`
	NickName  string `json:"nick_name" validate:"required" gorm:"column:nick_name"`
	Avatar    int    `json:"avatar" gorm:"column:avatar_img_id"`
	Email     string `json:"email" validate:"required" gorm:"column:email"`
	Password  string `json:"password,omitempty" validate:"required" gorm:"column:password"`
}

type UserSignIn struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

