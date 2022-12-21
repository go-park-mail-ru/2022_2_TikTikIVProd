package microservice

import (
	"context"

	userRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/user/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	user "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type microService struct {
	client user.UsersClient
}

func New(client user.UsersClient) userRep.RepositoryI {
	return &microService{
		client: client,
	}
}

func (userMS *microService) SelectUserByNickName(name string) (*models.User, error) {
	ctx := context.Background()

	pbSelectUserByNickNameRequest := user.SelectUserByNickNameRequest {
		NickName: name,
	}

	pbUser, err := userMS.client.SelectUserByNickName(ctx, &pbSelectUserByNickNameRequest)
	if err != nil {
		return nil, err
	}

	usr := &models.User {
		Id: pbUser.Id,
		FirstName: pbUser.FirstName,
		LastName: pbUser.LastName,
		NickName: pbUser.NickName,
		Avatar: pbUser.Avatar,
		Email: pbUser.Email,
		Password: pbUser.Password,
		CreatedAt: pbUser.CreatedAt.AsTime(),
	}

	return usr, nil
}

func (userMS *microService) SelectUserByEmail(email string) (*models.User, error) {
	ctx := context.Background()

	pbSelectUserByEmailRequest := user.SelectUserByEmailRequest {
		Email: email,
	}

	pbUser, err := userMS.client.SelectUserByEmail(ctx, &pbSelectUserByEmailRequest)
	if err != nil {
		return nil, err
	}

	usr := &models.User {
		Id: pbUser.Id,
		FirstName: pbUser.FirstName,
		LastName: pbUser.LastName,
		NickName: pbUser.NickName,
		Avatar: pbUser.Avatar,
		Email: pbUser.Email,
		Password: pbUser.Password,
		CreatedAt: pbUser.CreatedAt.AsTime(),
	}

	return usr, nil
}

func (userMS *microService) SelectUserById(id uint64) (*models.User, error) {
	ctx := context.Background()

	pbUserId := user.UserId {
		Id: id,
	}

	pbUser, err := userMS.client.SelectUserById(ctx, &pbUserId)
	if err != nil {
		return nil, err
	}

	usr := &models.User {
		Id: pbUser.Id,
		FirstName: pbUser.FirstName,
		LastName: pbUser.LastName,
		NickName: pbUser.NickName,
		Avatar: pbUser.Avatar,
		Email: pbUser.Email,
		Password: pbUser.Password,
		CreatedAt: pbUser.CreatedAt.AsTime(),
	}

	return usr, nil
}

func (userMS *microService) SelectAllUsers() ([]models.User, error) {
	ctx := context.Background()

	pbNothing := user.Nothing {
		Dummy: true,
	}

	pbUsers, err := userMS.client.SelectAllUsers(ctx, &pbNothing)
	if err != nil {
		return nil, err
	}

	users := make([]models.User, 0)

	for idx := range pbUsers.Users {
		usr := models.User {
			Id: pbUsers.Users[idx].Id,
			FirstName: pbUsers.Users[idx].FirstName,
			LastName: pbUsers.Users[idx].LastName,
			NickName: pbUsers.Users[idx].NickName,
			Avatar: pbUsers.Users[idx].Avatar,
			Email: pbUsers.Users[idx].Email,
			Password: pbUsers.Users[idx].Password,
			CreatedAt: pbUsers.Users[idx].CreatedAt.AsTime(),
		}

		users = append(users, usr)
	}
	
	return users, nil
}

func (userMS *microService) SearchUsers(name string) ([]models.User, error) {
	ctx := context.Background()

	pbSearchUsers := user.SearchUsersRequest {
		Name: name,
	}

	pbUsers, err := userMS.client.SearchUsers(ctx, &pbSearchUsers)
	if err != nil {
		return nil, err
	}

	users := make([]models.User, 0)

	for idx := range pbUsers.Users {
		usr := models.User {
			Id: pbUsers.Users[idx].Id,
			FirstName: pbUsers.Users[idx].FirstName,
			LastName: pbUsers.Users[idx].LastName,
			NickName: pbUsers.Users[idx].NickName,
			Avatar: pbUsers.Users[idx].Avatar,
			Email: pbUsers.Users[idx].Email,
			Password: pbUsers.Users[idx].Password,
			CreatedAt: pbUsers.Users[idx].CreatedAt.AsTime(),
		}

		users = append(users, usr)
	}
	
	return users, nil
}

func (userMS *microService) CreateUser(u *models.User) (error) {
	ctx := context.Background()

	ts := timestamppb.New(u.CreatedAt)
	pbUser := user.User {
		Id: u.Id,
		FirstName: u.FirstName,
		LastName: u.LastName,
		NickName: u.NickName,
		Avatar: u.Avatar,
		Email: u.Email,
		Password: u.Password,
		CreatedAt: ts,
	}

	userId, err := userMS.client.CreateUser(ctx, &pbUser)
	if err != nil {
		return err
	}

	u.Id = userId.Id

	return nil
}

func (userMS *microService) UpdateUser(usr models.User) (error) {
	ctx := context.Background()

	ts := timestamppb.New(usr.CreatedAt)
	pbUser := user.User {
		Id: usr.Id,
		FirstName: usr.FirstName,
		LastName: usr.LastName,
		NickName: usr.NickName,
		Avatar: usr.Avatar,
		Email: usr.Email,
		Password: usr.Password,
		CreatedAt: ts,
	}

	_, err := userMS.client.UpdateUser(ctx, &pbUser)
	if err != nil {
		return err
	}

	return nil
}

