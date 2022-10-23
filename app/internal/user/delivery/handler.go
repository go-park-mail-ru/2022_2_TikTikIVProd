package delivery

import (
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/usecase"
	_ "github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
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

// GetProfile godoc
// @Summary      GetProfile
// @Description  get user's profile
// @Tags     users
// @Produce  application/json
// @Success  200 {object} pkg.Response{body=models.User} "success"
// @Failure 405 {object} pkg.Error "invalid http method"
// @Failure 400 {object} pkg.Error "bad request"
// @Failure 500 {object} pkg.Error "internal server error"
// @Failure 404 {object} pkg.Error "can't find user with such id"
// @Router   /users/{id} [get]
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
	if err != nil {
		switch err.Error() {
		case "can't find user with such id":
			c.Logger().Error(err)
			return pkg.ErrorResponse(c.Response(), http.StatusNotFound, err.Error())
		default:
			c.Logger().Error(err)
			return pkg.ErrorResponse(c.Response(), http.StatusInternalServerError, err.Error())
		}
	}
	return pkg.JSONresponse(c.Response(), http.StatusOK, user)
}
