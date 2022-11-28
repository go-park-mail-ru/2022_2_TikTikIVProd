package usecase

import (
	userRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/UserMicroservice/internal/user/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/UserMicroservice/models"
	user "github.com/go-park-mail-ru/2022_2_TikTikIVProd/UserMicroservice/proto"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UseCaseI interface {
	SelectUserByNickName(*user.SelectUserByNickNameRequest) (*user.User, error)
	SelectUserByEmail(*user.SelectUserByEmailRequest) (*user.User, error)
	SelectUserById(*user.UserId) (*user.User, error)
	SelectAllUsers(*user.Nothing) (*user.UsersList, error)
	SearchUsers(*user.SearchUsersRequest) (*user.UsersList, error)
	CreateUser(*user.User) (*user.Nothing, error)
	UpdateUser(*user.User) (*user.Nothing, error)

	AddFriend(*user.Friends) (*user.Nothing, error)
	DeleteFriend(*user.Friends) (*user.Nothing, error)
	CheckFriends(*user.Friends) (*user.CheckFriendsResponse, error)
	SelectFriends(*user.UserId) (*user.UsersList, error)
}

type useCase struct {
	userRepository userRep.RepositoryI
}

func New(userRepository userRep.RepositoryI) UseCaseI {
	return &useCase{
		userRepository: userRepository,
	}
}

func (uc *useCase) SelectUserByNickName(nickName *user.SelectUserByNickNameRequest) (*user.User, error) {
	usr, err := uc.userRepository.SelectUserByNickName(nickName.NickName)
	if err != nil {
		return nil, errors.Wrap(err, "user repository postgres error")
	}

	ts := timestamppb.New(usr.CreatedAt)
	pbUser := &user.User {
		Id: usr.Id,
		FirstName: usr.FirstName,
		LastName: usr.LastName,
		NickName: usr.NickName,
		Avatar: usr.Avatar,
		Email: usr.Email,
		Password: usr.Password,
		CreatedAt: ts,
	}

	return pbUser, err
}

func (uc *useCase) SelectUserByEmail(email *user.SelectUserByEmailRequest) (*user.User, error) {
	usr, err := uc.userRepository.SelectUserByEmail(email.Email)
	if err != nil {
		return nil, errors.Wrap(err, "user repository postgres error")
	}

	ts := timestamppb.New(usr.CreatedAt)
	pbUser := &user.User {
		Id: usr.Id,
		FirstName: usr.FirstName,
		LastName: usr.LastName,
		NickName: usr.NickName,
		Avatar: usr.Avatar,
		Email: usr.Email,
		Password: usr.Password,
		CreatedAt: ts,
	}

	return pbUser, err
}

func (uc *useCase) SelectUserById(userId *user.UserId) (*user.User, error) {
	usr, err := uc.userRepository.SelectUserById(userId.Id)
	if err != nil {
		return nil, errors.Wrap(err, "user repository postgres error")
	}

	ts := timestamppb.New(usr.CreatedAt)
	pbUser := &user.User {
		Id: usr.Id,
		FirstName: usr.FirstName,
		LastName: usr.LastName,
		NickName: usr.NickName,
		Avatar: usr.Avatar,
		Email: usr.Email,
		Password: usr.Password,
		CreatedAt: ts,
	}

	return pbUser, err
}

func (uc *useCase) SelectAllUsers(nothing *user.Nothing) (*user.UsersList, error) {
	users, err := uc.userRepository.SelectAllUsers()
	if err != nil {
		return nil, errors.Wrap(err, "user repository postgres error")
	}

	var pbUsers *user.UsersList

	for idx := range users {
		ts := timestamppb.New(users[idx].CreatedAt)
		pbUser := &user.User {
			Id: users[idx].Id,
			FirstName: users[idx].FirstName,
			LastName: users[idx].LastName,
			NickName: users[idx].NickName,
			Avatar: users[idx].Avatar,
			Email: users[idx].Email,
			Password: users[idx].Password,
			CreatedAt: ts,
		}
		pbUsers.Users = append(pbUsers.Users, pbUser)
	}

	return pbUsers, err
}

func (uc *useCase) SearchUsers(name *user.SearchUsersRequest) (*user.UsersList, error) {
	users, err := uc.userRepository.SearchUsers(name.Name)
	if err != nil {
		return nil, errors.Wrap(err, "user repository postgres error")
	}

	var pbUsers *user.UsersList

	for idx := range users {
		ts := timestamppb.New(users[idx].CreatedAt)
		pbUser := &user.User {
			Id: users[idx].Id,
			FirstName: users[idx].FirstName,
			LastName: users[idx].LastName,
			NickName: users[idx].NickName,
			Avatar: users[idx].Avatar,
			Email: users[idx].Email,
			Password: users[idx].Password,
			CreatedAt: ts,
		}
		pbUsers.Users = append(pbUsers.Users, pbUser)
	}

	return pbUsers, err
}

func (uc *useCase) CreateUser(pbUser *user.User) (*user.Nothing, error) {
	usr := models.User {
		Id: pbUser.Id,
		FirstName: pbUser.FirstName,
		LastName: pbUser.LastName,
		NickName: pbUser.NickName,
		Avatar: pbUser.Avatar,
		Email: pbUser.Email,
		Password: pbUser.Password,
		CreatedAt: pbUser.CreatedAt.AsTime(),
	}

	err := uc.userRepository.CreateUser(&usr)
	if err != nil {
		return nil, errors.Wrap(err, "user repository postgres error")
	}

	pbUser.Id = usr.Id

	return &user.Nothing{Dummy: true}, err
}

func (uc *useCase) UpdateUser(pbUser *user.User) (*user.Nothing, error) {
	usr := models.User {
		Id: pbUser.Id,
		FirstName: pbUser.FirstName,
		LastName: pbUser.LastName,
		NickName: pbUser.NickName,
		Avatar: pbUser.Avatar,
		Email: pbUser.Email,
		Password: pbUser.Password,
		CreatedAt: pbUser.CreatedAt.AsTime(),
	}

	err := uc.userRepository.UpdateUser(usr)

	return &user.Nothing{Dummy: true}, err
}

func (uc *useCase) AddFriend(pbFriends *user.Friends) (*user.Nothing, error) {
	friends := models.Friends {
		Id1: pbFriends.Id1,
		Id2: pbFriends.Id2,
	}

	err := uc.userRepository.AddFriend(friends)
	if err != nil {
		return nil, errors.Wrap(err, "user repository postgres error")
	}

	return &user.Nothing{Dummy: true}, err
}

func (uc *useCase) DeleteFriend(pbFriends *user.Friends) (*user.Nothing, error) {
	friends := models.Friends {
		Id1: pbFriends.Id1,
		Id2: pbFriends.Id2,
	}

	err := uc.userRepository.DeleteFriend(friends)
	if err != nil {
		return nil, errors.Wrap(err, "user repository postgres error")
	}

	return &user.Nothing{Dummy: true}, err
}

func (uc *useCase) CheckFriends(pbFriends *user.Friends) (*user.CheckFriendsResponse, error) {
	friends := models.Friends {
		Id1: pbFriends.Id1,
		Id2: pbFriends.Id2,
	}

	isExists, err := uc.userRepository.CheckFriends(friends)
	if err != nil {
		return nil, errors.Wrap(err, "user repository postgres error")
	}

	return &user.CheckFriendsResponse{IsExists: isExists}, err
}

func (uc *useCase) SelectFriends(pbUserId *user.UserId) (*user.UsersList, error) {
	users, err := uc.userRepository.SelectFriends(pbUserId.Id)
	if err != nil {
		return nil, errors.Wrap(err, "user repository postgres error")
	}

	var pbUsers *user.UsersList

	for idx := range users {
		ts := timestamppb.New(users[idx].CreatedAt)
		pbUser := &user.User {
			Id: users[idx].Id,
			FirstName: users[idx].FirstName,
			LastName: users[idx].LastName,
			NickName: users[idx].NickName,
			Avatar: users[idx].Avatar,
			Email: users[idx].Email,
			Password: users[idx].Password,
			CreatedAt: ts,
		}
		pbUsers.Users = append(pbUsers.Users, pbUser)
	}

	return pbUsers, err
}

