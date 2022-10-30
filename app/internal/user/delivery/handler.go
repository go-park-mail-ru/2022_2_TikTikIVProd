package delivery

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	userUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/user/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/middleware"
	"github.com/labstack/echo/v4"
)

// type DeliveryI interface {
// 	GetProfile(c echo.Context) error
// }

type delivery struct {
	uc userUsecase.UseCaseI
}

// GetProfile godoc
// @Summary      GetProfile
// @Description  get user's profile
// @Tags     users
// @Produce  application/json
// @Param id path int true "User ID"
// @Success  200 {object} pkg.Response{body=models.User} "success"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 404 {object} echo.HTTPError "can't find user with such id"
// @Router   /users/{id} [get]
func (del *delivery) GetProfile(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, echo.GET)
	// c.Response().Header().Set("Access-Control-Allow-Credentials", "true")

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}
	user, err := del.uc.SelectUserById(id)
	if err != nil {
		switch {
		case errors.Is(errors.Cause(err), models.ErrNotFound):
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusNotFound, models.ErrNotFound.Error())
		default:
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(http.StatusOK, pkg.Response{Body: user})
}

func NewDelivery(e *echo.Echo, uc userUsecase.UseCaseI, authMid *middleware.Middleware) {
	handler := &delivery{
		uc: uc,
	}

	e.GET("/users/:id", handler.GetProfile, authMid.Auth)
}
