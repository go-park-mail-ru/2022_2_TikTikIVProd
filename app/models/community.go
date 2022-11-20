package models

import "time"

type Community struct {
	ID          int       `json:"id"`
	OwnerID     int       `json:"owner_id"`
	AvatarID    int       `json:"avatar_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreateDate  time.Time `json:"create_date" readonly:"true"`
}

type ReqCommunityCreate struct {
	AvatarID    int    `json:"avatar_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func ReqCreateToComm(req ReqCommunityCreate) Community {
	return Community{
		AvatarID:    req.AvatarID,
		Name:        req.Name,
		Description: req.Description,
	}
}
