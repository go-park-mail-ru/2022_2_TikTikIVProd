package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/friends/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg"
	"github.com/labstack/echo/v4"
)

type DeliveryI interface {
	AddFriend(c echo.Context) error
	DeleteFriend(c echo.Context) error
}

type delivery struct {
	uc usecase.UseCaseI
}

func New(uc usecase.UseCaseI) DeliveryI {
	return &delivery{
		uc: uc,
	}
}

// AddFriend godoc
// @Summary      AddFriend
// @Description  add friend
// @Tags     friends
// @Accept	 application/json
// @Produce  application/json
// @Param    friends body models.Friends true "friends info"
// @Success  201 "friend added"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 400 {object} pkg.Error "bad request"
// @Failure 500 {object} pkg.Error "internal server error"
// @Failure 409 {object} pkg.Error "friendship already exists"
// @Router   /friends/add [post]
func (del *delivery) AddFriend(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, echo.POST)
	if c.Request().Method != http.MethodPost {
		c.Logger().Error("invalid http method")
		return pkg.ErrorResponse(c.Response(), http.StatusMethodNotAllowed, "invalid http method")
	}

	friends := models.Friends{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&friends)
	if err != nil {
		c.Logger().Error(err.Error())
		return pkg.ErrorResponse(c.Response(), http.StatusBadRequest, "bad request")
	}

	err = del.uc.AddFriend(friends)
	switch {
	case err.Error() == "friendship already exists":
		c.Logger().Error(err)
		return pkg.ErrorResponse(c.Response(), http.StatusConflict, err.Error())
	case err != nil:
		c.Logger().Error(err)
		return pkg.ErrorResponse(c.Response(), http.StatusInternalServerError, err.Error())
	}

	c.Response().WriteHeader(http.StatusCreated)
	return nil
}

// DeleteFriend godoc
// @Summary      DeleteFriend
// @Description  delete friend
// @Tags     friends
// @Accept	 application/json
// @Produce  application/json
// @Param    friends body models.Friends true "friends info"
// @Success  204 "friend deleted, body is empty"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 400 {object} pkg.Error "bad request"
// @Failure 500 {object} pkg.Error "internal server error"
// @Failure 404 {object} pkg.Error "friend or user doesn't exist"
// @Router   /friends/delete [delete]
func (del *delivery) DeleteFriend(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, echo.DELETE)
	if c.Request().Method != http.MethodDelete {
		c.Logger().Error("invalid http method")
		return pkg.ErrorResponse(c.Response(), http.StatusMethodNotAllowed, "invalid http method")
	}

	friends := models.Friends{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&friends)
	if err != nil {
		c.Logger().Error(err.Error())
		return pkg.ErrorResponse(c.Response(), http.StatusBadRequest, "bad request")
	}

	err = del.uc.DeleteFriend(friends)
	switch {
	case err.Error() == "friend or user doesn't exist":
		c.Logger().Error(err)
		return pkg.ErrorResponse(c.Response(), http.StatusNotFound, err.Error())
	case err != nil:
		c.Logger().Error(err)
		return pkg.ErrorResponse(c.Response(), http.StatusInternalServerError, err.Error())
	}

	c.Response().WriteHeader(http.StatusNoContent)
	return nil
}

