package delivery

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	userUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/user/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/pkg"
	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
)

type Delivery struct {
	UserUC userUsecase.UseCaseI
}

// GetProfile godoc
// @Summary      GetProfile
// @Description  get user's profile
// @Tags     users
// @Produce  application/json
// @Param id path int true "User ID"
// @Success  200 {object} pkg.Response{body=models.User} "success get profile"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 404 {object} echo.HTTPError "can't find user with such id"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /users/{id} [get]
func (del *Delivery) GetProfile(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}
	user, err := del.UserUC.SelectUserById(id)
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
	return c.JSON(http.StatusOK, pkg.Response{Body: user})
}

// GetUsers godoc
// @Summary      GetUsers
// @Description  get all users
// @Tags     users
// @Produce  application/json
// @Success  200 {object} pkg.Response{body=[]models.User} "success get users"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /users [get]
func (del *Delivery) GetUsers(c echo.Context) error {
	users, err := del.UserUC.SelectUsers()
	if err != nil {
		causeErr := errors.Cause(err)
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, causeErr.Error())
	}
	return c.JSON(http.StatusOK, pkg.Response{Body: users})
}

// UpdateUser godoc
// @Summary      UpdateUser
// @Description  update user's profile
// @Tags     users
// @Accept	 application/json
// @Produce  application/json
// @Param user body models.User true "user data"
// @Success  204 "success update"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 404 {object} echo.HTTPError "can't find user with such id"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /users/update [put]
func (del *Delivery) UpdateUser(c echo.Context) error {
	var user models.User
	err := c.Bind(&user); if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	requestSanitizeUpdateUser(&user)

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	user.Id = userId
	
	err = del.UserUC.UpdateUser(user)
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
	return c.NoContent(http.StatusNoContent)
}

// SearchUsers godoc
// @Summary      SearchUsers
// @Description  search users by name
// @Tags     users
// @Produce  application/json
// @Param name path string true "User name"
// @Success  200 {object} pkg.Response{body=[]models.User} "success search users"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /users/search/{name} [get]
func (del *Delivery) SearchUsers(c echo.Context) error {
	name := c.Param("name")
	users, err := del.UserUC.SearchUsers(name)
	if err != nil {
		causeErr := errors.Cause(err)
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, causeErr.Error())
	}
	return c.JSON(http.StatusOK, pkg.Response{Body: users})
}

func requestSanitizeUpdateUser(user *models.User) {
	sanitizer := bluemonday.UGCPolicy()

	user.FirstName = sanitizer.Sanitize(user.FirstName)
	user.LastName = sanitizer.Sanitize(user.LastName)
	user.NickName = sanitizer.Sanitize(user.NickName)
	user.Email = sanitizer.Sanitize(user.Email)
	user.Password = sanitizer.Sanitize(user.Password)
}

func NewDelivery(e *echo.Echo, uc userUsecase.UseCaseI) {
	handler := &Delivery{
		UserUC: uc,
	}

	e.GET("/users/:id", handler.GetProfile)
	e.GET("/users", handler.GetUsers)
	e.GET("/users/search/:name", handler.SearchUsers)
	e.PUT("/users/update", handler.UpdateUser)
}
