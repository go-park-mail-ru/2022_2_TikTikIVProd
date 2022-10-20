package delivery

import (
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg"
	"github.com/labstack/echo/v4"
)

type DeliveryI interface {
	GetProfile(c echo.Context) error
}

type delivery struct {
	uc usecase.UseCaseI
}

func New(uc usecase.UseCaseI) DeliveryI {
	return &delivery{
		uc: uc,
	}
}

func (del *delivery) GetProfile(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, echo.GET)
	if c.Request().Method != http.MethodGet {
		c.Logger().Error("invalid http method")
		return pkg.ErrorResponse(c.Response(), http.StatusMethodNotAllowed, "invalid http method")
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Logger().Error(err)
		return pkg.ErrorResponse(c.Response(), http.StatusBadRequest, "bad request")
	}
	user, err := del.uc.SelectUserById(id)
	if err.Error() == "can't find user with such id" {
		c.Logger().Error(err)
		return pkg.ErrorResponse(c.Response(), http.StatusNotFound, err.Error())
	} else if err != nil {
		c.Logger().Error(err)
		return pkg.ErrorResponse(c.Response(), http.StatusInternalServerError, err.Error())
	}
	return pkg.JSONresponse(c.Response(), http.StatusOK, user)
}

