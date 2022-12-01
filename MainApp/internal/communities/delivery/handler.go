package delivery

import (
	communitiesUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/communities/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/pkg"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type Delivery struct {
	CommUC communitiesUsecase.UseCaseI
}

// CreateCommunity godoc
// @Summary      Create a community
// @Description  Create a community
// @Tags     	 communities
// @Accept	 application/json
// @Produce  application/json
// @Param    community body models.ReqCommunityCreate true "community info"
// @Success  200 {object} pkg.Response{body=models.Community} "success create community"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 422 {object} echo.HTTPError "unprocessable entity"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /communities/create [post]
func (delivery *Delivery) CreateCommunity(c echo.Context) error {
	var reqComm models.ReqCommunityCreate
	err := c.Bind(&reqComm)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, models.ErrBadRequest.Error())
	}

	if ok, err := isRequestValid(&reqComm); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	//requestSanitizePost(&post)

	userId, ok := c.Get("user_id").(uint64)

	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	comm := models.ReqCreateToComm(reqComm)
	comm.OwnerID = userId
	err = delivery.CommUC.CreateCommunity(&comm)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: comm})
}

// UpdateCommunity godoc
// @Summary      Update a community
// @Description  Update a community
// @Tags     	 communities
// @Accept	 application/json
// @Produce  application/json
// @Param    community body models.Community true "community info"
// @Success  200 {object} pkg.Response{body=models.Community} "success update community"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 422 {object} echo.HTTPError "unprocessable entity"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf or permission denied"
// @Router   /communities/edit [post]
func (delivery *Delivery) UpdateCommunity(c echo.Context) error {
	var comm models.Community
	err := c.Bind(&comm)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, models.ErrBadRequest.Error())
	}

	if ok, err := isRequestValid(&comm); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	//requestSanitizePost(&post) TODO

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	comm.OwnerID = userId
	err = delivery.CommUC.UpdateCommunity(&comm)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: comm})
}

// GetCommunity godoc
// @Summary      Get a community
// @Description  Get a community info
// @Tags     communities
// @Produce  application/json
// @Param id path int true "community ID"
// @Success  200 {object} pkg.Response{body=models.Community} "success get community"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 404 {object} echo.HTTPError "can't find community with such id"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /communities/{id} [get]
func (delivery *Delivery) GetCommunity(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	community, err := delivery.CommUC.GetCommunity(id)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: community})
}

// GetAllCommunities godoc
// @Summary      Get all communities
// @Description  Get all communities
// @Tags     communities
// @Produce  application/json
// @Success  200 {object} pkg.Response{body=models.Community} "success get community"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /communities [get]
func (delivery *Delivery) GetAllCommunities(c echo.Context) error {
	communities, err := delivery.CommUC.GetAllCommunities()

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: communities})
}

// SearchCommunity godoc
// @Summary      Search a community
// @Description  Search a community
// @Tags     	 communities
// @Accept	 	 application/json
// @Param        q    query     string  false  "name search by q"
// @Success  200 {object} pkg.Response{body=models.Community} "success create community"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /communities/{id} [delete]
func (delivery *Delivery) SearchCommunity(c echo.Context) error {
	param := c.QueryParam("q")

	communities, err := delivery.CommUC.SearchCommunities(param)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: communities})
}

// DeleteCommunity godoc
// @Summary      Delete a community
// @Description  Delete a community
// @Tags     	 communities
// @Accept	 application/json
// @Param id path int  true  "Community ID"
// @Success  204
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 404 {object} echo.HTTPError "can't find community with such id"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /communities/{id} [delete]
func (delivery *Delivery) DeleteCommunity(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, models.ErrNotFound.Error())
	}

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	err = delivery.CommUC.DeleteCommunity(id, userId)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestValid(c interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return false, err
	}
	return true, nil
}

func handleError(err error) *echo.HTTPError {
	causeErr := errors.Cause(err)
	switch {
	case errors.Is(causeErr, models.ErrNotFound):
		return echo.NewHTTPError(http.StatusNotFound, models.ErrNotFound.Error())
	case errors.Is(causeErr, models.ErrBadRequest):
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	case errors.Is(causeErr, models.ErrPermissionDenied):
		return echo.NewHTTPError(http.StatusForbidden, models.ErrPermissionDenied.Error())
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, causeErr.Error())
	}
}

func NewDelivery(e *echo.Echo, cu communitiesUsecase.UseCaseI) {
	handler := &Delivery{
		CommUC: cu,
	}

	e.POST("/communities/create", handler.CreateCommunity)
	e.POST("/communities/edit", handler.UpdateCommunity)
	e.GET("/communities/:id", handler.GetCommunity)
	e.GET("/communities/search", handler.SearchCommunity)
	e.GET("/communities", handler.GetAllCommunities)
	e.DELETE("/communities/:id", handler.DeleteCommunity)
}
