package microservice

import (
	"context"

	friendsRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/friends/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	user "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto/user"
	"github.com/pkg/errors"
)

type microService struct {
	client user.UsersClient
}

func New(client user.UsersClient) friendsRep.RepositoryI {
	return &microService{
		client: client,
	}
}

func (userMS *microService) SelectFriends(id uint64) ([]models.User, error) {
	ctx := context.Background()

	pbUserId := user.UserId {
		Id: id,
	}

	pbUsers, err := userMS.client.SelectFriends(ctx, &pbUserId)
	if err != nil {
		return nil, errors.Wrap(err, "user microservice error")
	}

	var users []models.User

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

func (userMS *microService) AddFriend(friends models.Friends) (error) {
	ctx := context.Background()

	pbFriends := user.Friends {
		Id1: friends.Id1,
		Id2: friends.Id2,
	}

	_, err := userMS.client.AddFriend(ctx, &pbFriends)
	if err != nil {
		return errors.Wrap(err, "user microservice error")
	}

	return nil
}

func (userMS *microService) DeleteFriend(friends models.Friends) (error) {
	ctx := context.Background()

	pbFriends := user.Friends {
		Id1: friends.Id1,
		Id2: friends.Id2,
	}

	_, err := userMS.client.DeleteFriend(ctx, &pbFriends)
	if err != nil {
		return errors.Wrap(err, "user microservice error")
	}

	return nil
}

func (userMS *microService) CheckFriends(friends models.Friends) (bool, error) {
	ctx := context.Background()

	pbFriends := user.Friends {
		Id1: friends.Id1,
		Id2: friends.Id2,
	}

	pbCheckFriendsResponse, err := userMS.client.CheckFriends(ctx, &pbFriends)
	if err != nil {
		return false, errors.Wrap(err, "user microservice error")
	}

	return pbCheckFriendsResponse.IsExists, nil
}

