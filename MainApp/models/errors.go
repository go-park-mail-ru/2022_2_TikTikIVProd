package models

import (
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

var (
	ErrNotFound            = errors.New("item is not found")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrConflictNickname    = errors.New("nickname already exists")
	ErrConflictEmail       = errors.New("email already exists")
	ErrBadRequest          = errors.New("bad request")
	ErrConflictFriend      = errors.New("friend already exists")
	ErrUnauthorized        = errors.New("no cookie")
	ErrInternalServerError = errors.New("internal server error")
	ErrEmptyCsrf           = errors.New("empty csrf token")
	ErrInvalidCsrf         = errors.New("invalid csrf")
	ErrPermissionDenied    = errors.New("permission denied")
)

func ErrEq(err error, target error) bool {
	log.Info("err1: ", err.Error(), "target:", target.Error())
	log.Info("cause: ", errors.Cause(err).Error(), "target:", target.Error())
	return errors.Cause(err).Error() == target.Error()
}
