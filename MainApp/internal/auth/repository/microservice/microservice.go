package microservice

import (
	"context"

	authRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/auth/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	auth "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/proto"
	"github.com/pkg/errors"
)

type microService struct {
	client auth.AuthClient
}

func New(client auth.AuthClient) authRep.RepositoryI {
	return &microService{
		client: client,
	}
}

func (authMS *microService) CreateCookie(cookie *models.Cookie) error {
	ctx := context.Background()

	pbCookie := auth.Cookie{
		SessionToken: cookie.SessionToken,
		UserId:       cookie.UserId,
		MaxAge:       cookie.MaxAge,
	}

	_, err := authMS.client.CreateCookie(ctx, &pbCookie)
	if err != nil {
		return errors.Wrap(err, "auth microservice error")
	}

	return nil
}

func (authMS *microService) GetCookie(value string) (string, error) {
	ctx := context.Background()

	pbValueCookieRequest := auth.ValueCookieRequest{
		ValueCookie: value,
	}

	userId, err := authMS.client.GetCookie(ctx, &pbValueCookieRequest)
	if err != nil {
		return "", errors.Wrap(err, "auth microservice error")
	}

	strUserId := userId.UserId

	return strUserId, nil
}

func (authMS *microService) DeleteCookie(value string) error {
	ctx := context.Background()

	pbValueCookieRequest := auth.ValueCookieRequest{
		ValueCookie: value,
	}

	_, err := authMS.client.DeleteCookie(ctx, &pbValueCookieRequest)
	if err != nil {
		return errors.Wrap(err, "auth microservice error")
	}

	return nil
}
