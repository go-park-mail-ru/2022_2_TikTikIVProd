package delivery

import (
	imgUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/middleware"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"strconv"
)

type DeliveryI interface {
	GetImageByID(c echo.Context) error
	UploadImage(c echo.Context) error
}

type delivery struct {
	imgUsecase imgUsecase.ImageUseCaseI
}

// GetImageByID godoc
// @Summary      Get image by id
// @Description  Get image by id
// @Tags     	 image
// @Param id path int  true  "image ID"
// @Produce  image/png
// @Success  200 "success get image"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Router   /image/{id} [get]
func (delivery *delivery) GetImageByID(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	img, err := delivery.imgUsecase.GetImageById(id)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	f, err := os.Open("images/" + img.ImgLink)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.Stream(http.StatusOK, "image/png", f)
}

// UploadImage godoc
// @Summary      Upload image
// @Description  Upload image
// @Tags     	 image
// @Param image formData file  true  "image file"
// @Accept multipart/form-data
// @Produce  application/json
// @Success  200 "success upload image"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Router   /image/upload [post]
func (delivery *delivery) UploadImage(c echo.Context) error {
	// Source
	file, err := c.FormFile("image")
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Not image in form")
	}
	src, err := file.Open()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("images/" + file.Filename)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")

	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	image := models.Image{ImgLink: file.Filename}
	_ = delivery.imgUsecase.CreateImage(&image)

	return c.JSON(http.StatusOK, pkg.Response{Body: image})
}

func NewDelivery(e *echo.Echo, iu imgUsecase.ImageUseCaseI, authMid *middleware.Middleware) {
	handler := &delivery{
		imgUsecase: iu,
	}

	e.POST("/image/upload", handler.UploadImage, authMid.Auth)
	e.GET("/image/:id", handler.GetImageByID, authMid.Auth)
}
