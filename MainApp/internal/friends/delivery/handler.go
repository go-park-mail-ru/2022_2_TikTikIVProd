package delivery

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	friendUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/friends/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/pkg"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type Delivery struct {
	FriendsUC friendUsecase.UseCaseI
}

// AddFriend godoc
// @Summary      AddFriend
// @Description  add friend
// @Tags     friends
// @Produce  application/json
// @Param friend_id path int true "Friend ID"
// @Success  201 "friend added"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 404 {object} echo.HTTPError "friend or user doesn't exist"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 409 {object} echo.HTTPError "friend already exists"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /friends/add/{friend_id} [post]
func (del *Delivery) AddFriend(c echo.Context) error {
	var friends models.Friends
	err := c.Bind(&friends); if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	if ok, err := isRequestValid(&friends); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	friends.Id1 = userId

	err = del.FriendsUC.AddFriend(friends)
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
			return echo.NewHTTPError(http.StatusInternalServerError, causeErr.Error())
		}
	}

	return c.NoContent(http.StatusCreated)
}

// DeleteFriend godoc
// @Summary      DeleteFriend
// @Description  delete friend
// @Tags     friends
// @Produce  application/json
// @Param friend_id path int true "Friend ID"
// @Success  204 "friend deleted, body is empty"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 404 {object} echo.HTTPError "friend/user/friendship doesn't exist"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /friends/delete/{friend_id} [delete]
func (del *Delivery) DeleteFriend(c echo.Context) error {
	var friends models.Friends
	err := c.Bind(&friends); if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	if ok, err := isRequestValid(&friends); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	friends.Id1 = userId

	err = del.FriendsUC.DeleteFriend(friends)
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
			return echo.NewHTTPError(http.StatusInternalServerError, causeErr.Error())
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

// SelectFriends godoc
// @Summary      SelectFriends
// @Description  get user's friends
// @Tags     friends
// @Produce  application/json
// @Param user_id path int true "User ID"
// @Success  200 {object} pkg.Response{body=[]models.User} "success get profile"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 404 {object} echo.HTTPError "user doesn't exist"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Router   /friends/{user_id} [get]
func (del *Delivery) SelectFriends(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	friends, err := del.FriendsUC.SelectFriends(id)
	if err != nil {
		causeErr := errors.Cause(err)
		switch {
		case errors.Is(causeErr, models.ErrNotFound):
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusNotFound, models.ErrNotFound.Error())
		default:
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, causeErr.Error())
		}
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: friends})
}

// CheckIsFriend godoc
// @Summary      CheckIsFriend
// @Description  check friend
// @Tags     friends
// @Produce  application/json
// @Param friend_id path int true "Friend ID"
// @Success  200 {object} pkg.Response{body=bool} "success get profile"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 404 {object} echo.HTTPError "item doesn't exist"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Router   /friends/check/{friend_id} [get]
func (del *Delivery) CheckIsFriend(c echo.Context) error {
	friendId, err := strconv.ParseUint(c.Param("friend_id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	friends := models.Friends {
		Id1: userId,
		Id2: friendId,
	}

	isFriend, err := del.FriendsUC.CheckIsFriend(friends)
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
			return echo.NewHTTPError(http.StatusInternalServerError, causeErr.Error())
		}
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: isFriend})
}


func NewDelivery(e *echo.Echo, uc friendUsecase.UseCaseI) {
	handler := &Delivery{
		FriendsUC: uc,
	}

	e.POST("/friends/add/:friend_id", handler.AddFriend)
	e.DELETE("/friends/delete/:friend_id", handler.DeleteFriend)
	e.GET("/friends/:user_id", handler.SelectFriends)
	e.GET("/friends/check/:friend_id", handler.CheckIsFriend)
}

