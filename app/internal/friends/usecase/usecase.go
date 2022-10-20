package usecase

import (
	"errors"

	friendsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
)

//go:generate mockgen -source=usecase.go -destination=mocks/mock.go

type UseCaseI interface {
	AddFriend(friends models.Friends) (error)
	DeleteFriend(friends models.Friends) (error)
}

type useCase struct {
	repository friendsRep.RepositoryI
}

func New(rep friendsRep.RepositoryI) UseCaseI {
	return &useCase{
		repository: rep,
	}
}

func (uc *useCase) AddFriend(friends models.Friends) (error) {
	err := uc.repository.AddFriend(friends)
	if err != nil {
		return errors.New("friendship already exists")
	}
	return err
}

func (uc *useCase) DeleteFriend(friends models.Friends) (error) {
	err := uc.repository.DeleteFriend(friends)
	if err != nil {
		return errors.New("friend or user doesn't exist")
	}
	return nil
}
