package usecase

import (
	authRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AuthMicroservice/internal/auth/repository"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/AuthMicroservice/models"
	auth "github.com/go-park-mail-ru/2022_2_TikTikIVProd/proto"
)

type UseCaseI interface {
	GetCookie(*auth.ValueCookieRequest) (*auth.GetCookieResponse, error)
	DeleteCookie(*auth.ValueCookieRequest) (*auth.Nothing, error)
	CreateCookie(*auth.Cookie) (*auth.Nothing, error)
}

type useCase struct {
	authRepository authRep.RepositoryI
}

func New(authRepository authRep.RepositoryI) UseCaseI {
	return &useCase{
		authRepository: authRepository,
	}
}

func (uc *useCase) GetCookie(value *auth.ValueCookieRequest) (*auth.GetCookieResponse, error) {
	userId, err := uc.authRepository.GetCookie(value.ValueCookie)
	return &auth.GetCookieResponse{UserId: userId}, err
}

func (uc *useCase) DeleteCookie(value *auth.ValueCookieRequest) (*auth.Nothing, error) {
	err := uc.authRepository.DeleteCookie(value.ValueCookie)
	return &auth.Nothing{Dummy: true}, err
}

func (uc *useCase) CreateCookie(cookie *auth.Cookie) (*auth.Nothing, error) {
	modelCookie := models.Cookie {
		SessionToken: cookie.SessionToken,
		UserId: cookie.UserId,
		MaxAge: cookie.MaxAge,
	}
	err := uc.authRepository.CreateCookie(&modelCookie)
	return &auth.Nothing{Dummy: true}, err
}

