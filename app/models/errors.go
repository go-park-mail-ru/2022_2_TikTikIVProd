package models

import "errors"

var (
	ErrNotFound = errors.New("item is not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrConflictNickname = errors.New("nickname already exists")
	ErrConflictEmail = errors.New("email already exists")
	ErrBadRequest = errors.New("bad request")
	ErrConflictFriend = errors.New("friend already exists")
	ErrUnauthorized = errors.New("no cookie")
	ErrInternalServerError = errors.New("internal server error")
)

