package delivery

import (
	"context"
	authUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AuthMicroservice/internal/auth/usecase"
	auth "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AuthMicroservice/proto"
)

type AuthManager struct {
	auth.UnimplementedAuthServer
	AuthUC authUsecase.UseCaseI
}

func New(uc authUsecase.UseCaseI) auth.AuthServer {
	return AuthManager{AuthUC: uc}
}

func (am AuthManager) GetCookie(ctx context.Context, value *auth.ValueCookieRequest) (*auth.GetCookieResponse, error) {
	resp, err := am.AuthUC.GetCookie(value)
	return resp, err
}

func (am AuthManager) DeleteCookie(ctx context.Context, value *auth.ValueCookieRequest) (*auth.Nothing, error) {
	resp, err := am.AuthUC.DeleteCookie(value)
	return resp, err
}

func (am AuthManager) CreateCookie(ctx context.Context, cookie *auth.Cookie) (*auth.Nothing, error) {
	resp, err := am.AuthUC.CreateCookie(cookie)
	return resp, err
}
