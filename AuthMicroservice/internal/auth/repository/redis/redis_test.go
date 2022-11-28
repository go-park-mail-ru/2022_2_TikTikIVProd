package redis_test

import (
	"strconv"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/bxcodec/faker"
	authRep "github.com/go-park-mail-ru/2022_2_TikTikIVProd/AuthMicroservice/internal/auth/repository/redis"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/AuthMicroservice/models"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCaseGetCookie struct {
	ArgData string
	ExpectedRes string
	Error error
}

func TestCreateCookie(t *testing.T) {
	s := miniredis.RunT(t)

	redisClient := redis.NewClient(&redis.Options{Addr: s.Addr()})

	repository := authRep.New(redisClient)

	mockCookie := models.Cookie {
		SessionToken: "session_token",
		UserId: 1,
		MaxAge: 3600 * 24 * 365,
	}

	err := repository.CreateCookie(&mockCookie)
	require.NoError(t, err)

	s.CheckGet(t, mockCookie.SessionToken, strconv.Itoa(mockCookie.UserId))
}

func TestGetCookie(t *testing.T) {
	s := miniredis.RunT(t)

	redisClient := redis.NewClient(&redis.Options{Addr: s.Addr()})

	repository := authRep.New(redisClient)

	var mockCookie models.Cookie
	err := faker.FakeData(&mockCookie)
	assert.NoError(t, err)

	cases := map[string]TestCaseGetCookie {
		"success": {
			ArgData:   mockCookie.SessionToken,
			ExpectedRes: strconv.Itoa(mockCookie.UserId),
			Error: nil,
		},
		"not_found": {
			ArgData:   mockCookie.SessionToken,
			Error: models.ErrNotFound,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			if name == "success" {
				s.Set(test.ArgData, test.ExpectedRes)
			} else if name == "not_found" {
				s.Del(test.ArgData)
			}
			userId, err := repository.GetCookie(test.ArgData)
			require.Equal(t, test.Error, err)

			if err == nil {
				assert.Equal(t, test.ExpectedRes, userId)
			}
		})
	}
}

func TestDeleteCookie(t *testing.T) {
	s := miniredis.RunT(t)

	redisClient := redis.NewClient(&redis.Options{Addr: s.Addr()})

	repository := authRep.New(redisClient)

	var mockCookie models.Cookie
	err := faker.FakeData(&mockCookie)
	assert.NoError(t, err)

	err = repository.DeleteCookie(mockCookie.SessionToken)
	assert.NoError(t, err)
}
