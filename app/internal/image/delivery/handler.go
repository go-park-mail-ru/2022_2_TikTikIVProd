package delivery

import (
	imgUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/image/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strconv"
)

type DeliveryI interface {
	GetImageByID(c echo.Context) error
	//GetPost(w http.ResponseWriter, r *http.Request)
	//CreatePost(w http.ResponseWriter, r *http.Request)
	//UpdatePost(w http.ResponseWriter, r *http.Request)
	//DeletePost(w http.ResponseWriter, r *http.Request)
}

type delivery struct {
	imgUsecase imgUsecase.ImageUseCaseI
}

func NewDelivery(imgUsecase imgUsecase.ImageUseCaseI) DeliveryI {
	return &delivery{
		imgUsecase: imgUsecase,
	}
}

func (delivery *delivery) GetImageByID(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	img, err := delivery.imgUsecase.GetImage(id)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	f, err := os.Open(img.ImgLink)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return c.Stream(http.StatusOK, "image/png", f)
}
