package delivery

import (
	"net/http"
	"github.com/pkg/errors"

	friendUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/middleware"
)

// type DeliveryI interface {
// 	AddFriend(c echo.Context) error
// 	DeleteFriend(c echo.Context) error
// }

type delivery struct {
	uc friendUsecase.UseCaseI
}

// AddFriend godoc
// @Summary      AddFriend
// @Description  add friend
// @Tags     friends
// @Accept	 application/json
// @Produce  application/json
// @Param    friends body models.Friends true "friends info"
// @Success  201 "friend added"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 404 {object} echo.HTTPError "friend or user doesn't exist"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 409 {object} echo.HTTPError "friend already exists"
// @Router   /friends/add [post]
func (del *delivery) AddFriend(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, echo.POST)

	var friends models.Friends
	err := c.Bind(&friends); if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	if ok, err := isRequestValid(&friends); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	err = del.uc.AddFriend(friends)
	if err != nil {
		causeErr := errors.Cause(err)
		switch {
		case errors.Is(causeErr, models.ErrNotFound):
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusNotFound, models.ErrNotFound.Error())
		case errors.Is(causeErr, models.ErrConflictFriend):
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusConflict, models.ErrConflictFriend.Error())
		case errors.Is(causeErr, models.ErrBadRequest):
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
		default:
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.NoContent(http.StatusCreated)
}

// DeleteFriend godoc
// @Summary      DeleteFriend
// @Description  delete friend
// @Tags     friends
// @Produce  application/json
// @Param id_user path int true "User ID"
// @Param id_friend path int true "Friend ID"
// @Success  204 "friend deleted, body is empty"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 404 {object} echo.HTTPError "friend/user/friendship doesn't exist"
// @Router   /friends/delete/{id_user}/{id_friend} [delete]
func (del *delivery) DeleteFriend(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, echo.DELETE)

	var friends models.Friends
	err := c.Bind(&friends); if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	if ok, err := isRequestValid(&friends); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	err = del.uc.DeleteFriend(friends)
	if err != nil {
		causeErr := errors.Cause(err)
		switch {
		case errors.Is(causeErr, models.ErrNotFound):
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusNotFound, models.ErrNotFound.Error())
		case errors.Is(causeErr, models.ErrBadRequest):
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
		default:
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestValid(user interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewDelivery(e *echo.Echo, uc friendUsecase.UseCaseI, authMid *middleware.Middleware) {
	handler := &delivery{
		uc: uc,
	}

	e.POST("/friends/add", handler.AddFriend, authMid.Auth)
	e.DELETE("/friends/delete/:id_user/:id_friend", handler.DeleteFriend, authMid.Auth)
}
