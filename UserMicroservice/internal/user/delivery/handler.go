package delivery

import (
	"context"
	userUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/UserMicroservice/internal/user/usecase"
	user "github.com/go-park-mail-ru/2022_2_TikTikIVProd/UserMicroservice/proto"
)

type UserManager struct {
	user.UnimplementedUsersServer
	UserUC userUsecase.UseCaseI
}

func New(uc userUsecase.UseCaseI) user.UsersServer {
	return UserManager{UserUC: uc}
}

func (um UserManager) SelectUserByNickName(ctx context.Context, nickName *user.SelectUserByNickNameRequest) (*user.User, error) {
	resp, err := um.UserUC.SelectUserByNickName(nickName)
	return resp, err
}

func (um UserManager) SelectUserByEmail(ctx context.Context, email *user.SelectUserByEmailRequest) (*user.User, error) {
	resp, err := um.UserUC.SelectUserByEmail(email)
	return resp, err
}

func (um UserManager) SelectUserById(ctx context.Context, id *user.UserId) (*user.User, error) {
	resp, err := um.UserUC.SelectUserById(id)
	return resp, err
}

func (um UserManager) SelectAllUsers(ctx context.Context, nothing *user.Nothing) (*user.UsersList, error) {
	resp, err := um.UserUC.SelectAllUsers(nothing)
	return resp, err
}

func (um UserManager) SearchUsers(ctx context.Context, name *user.SearchUsersRequest) (*user.UsersList, error) {
	resp, err := um.UserUC.SearchUsers(name)
	return resp, err
}

func (um UserManager) CreateUser(ctx context.Context, usr *user.User) (*user.UserId, error) {
	_, err := um.UserUC.CreateUser(usr)
	return &user.UserId{Id:usr.Id}, err
}

func (um UserManager) UpdateUser(ctx context.Context, usr *user.User) (*user.Nothing, error) {
	resp, err := um.UserUC.UpdateUser(usr)
	return resp, err
}

func (um UserManager) AddFriend(ctx context.Context, friends *user.Friends) (*user.Nothing, error) {
	resp, err := um.UserUC.AddFriend(friends)
	return resp, err
}

func (um UserManager) DeleteFriend(ctx context.Context, friends *user.Friends) (*user.Nothing, error) {
	resp, err := um.UserUC.DeleteFriend(friends)
	return resp, err
}

func (um UserManager) CheckFriends(ctx context.Context, friends *user.Friends) (*user.CheckFriendsResponse, error) {
	resp, err := um.UserUC.CheckFriends(friends)
	return resp, err
}

func (um UserManager) SelectFriends(ctx context.Context, userId *user.UserId) (*user.UsersList, error) {
	resp, err := um.UserUC.SelectFriends(userId)
	return resp, err
}
