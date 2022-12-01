package models

type Cookie struct {
	SessionToken string
	UserId       uint64
	MaxAge      int64
}
