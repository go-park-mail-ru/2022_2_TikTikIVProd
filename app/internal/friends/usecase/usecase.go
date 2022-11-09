package usecase

import (
	"github.com/pkg/errors"

	friendsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/repository"
	usersRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
)

type UseCaseI interface {
	AddFriend(friends models.Friends) (error)
	DeleteFriend(friends models.Friends) (error)
	SelectFriends(id int) ([]models.User, error)
	CheckIsFriend(friends models.Friends) (bool, error)
}

type useCase struct {
	friendsRepository friendsRep.RepositoryI
	userRepository usersRep.RepositoryI
}

func New(fRep friendsRep.RepositoryI, uRep usersRep.RepositoryI) UseCaseI {
	return &useCase{
		friendsRepository: fRep,
		userRepository: uRep,
	}
}

func (uc *useCase) AddFriend(friends models.Friends) error {
	if friends.Id1 == friends.Id2 {
		return models.ErrBadRequest
	}

	_, err := uc.userRepository.SelectUserById(friends.Id2)
	if err != nil {
		return errors.Wrap(err, "user repository error")
	}

	friendExists, err := uc.friendsRepository.CheckFriends(friends)
	if err != nil {
		return errors.Wrap(err, "friends repository error")
	}
	if friendExists {
		return models.ErrConflictFriend
	}

	err = uc.friendsRepository.AddFriend(friends)
	if err != nil {
		return errors.Wrap(err, "friends repository error")
	}

	// dialog := models.Dialog {
	// 	UserId1: friends.Id1,
	// 	UserId2: friends.Id2,
	// }

	// checkFriends := models.Friends {
	// 	Id1: friends.Id2,
	// 	Id2: friends.Id1,
	// }

	// friendExists, err = uc.friendsRepository.CheckFriends(checkFriends)
	// if err != nil {
	// 	return errors.Wrap(err, "friends repository error")
	// }
	// if friendExists {
	// 	err = uc.chatRepository.CreateDialog(&dialog)
	// 	if err != nil {
	// 		return errors.Wrap(err, "friends repository error")
	// 	}
	// }

	return err
}

func (uc *useCase) DeleteFriend(friends models.Friends) error {
	if friends.Id1 == friends.Id2 {
		return models.ErrBadRequest
	}
	
	_, err := uc.userRepository.SelectUserById(friends.Id2)
	if err != nil {
		return errors.Wrap(err, "user repository error")
	}

	friendExists, err := uc.friendsRepository.CheckFriends(friends)
	if err != nil {
		return errors.Wrap(err, "friends repository error")
	}
	if !friendExists {
		return models.ErrNotFound
	}

	err = uc.friendsRepository.DeleteFriend(friends)
	if err != nil {
		return errors.Wrap(err, "friends repository error")
	}

	return nil
}

func (uc *useCase) SelectFriends(id int) ([]models.User, error) {
	_, err := uc.userRepository.SelectUserById(id)
	if err != nil {
		return nil, errors.Wrap(err, "user repository error")
	}

	friends, err := uc.friendsRepository.SelectFriends(id)
	if err != nil {
		return nil, errors.Wrap(err, "friends repository error")
	}

	return friends, nil
}

func (uc *useCase) CheckIsFriend(friends models.Friends) (bool, error) {
	if friends.Id1 == friends.Id2 {
		return false, models.ErrBadRequest
	}

	_, err := uc.userRepository.SelectUserById(friends.Id2)
	if err != nil {
		return false, errors.Wrap(err, "user repository error")
	}

	friendExists, err := uc.friendsRepository.CheckFriends(friends)
	if err != nil {
		return false, errors.Wrap(err, "friends repository error")
	}

	return friendExists, err
}

