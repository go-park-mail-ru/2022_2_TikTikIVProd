package delivery

import (
	"net/http"
	"os"
	"strconv"

	stickersUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/stickers/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/pkg"
	"github.com/labstack/echo/v4"
)

type DeliveryI interface {
	GetStickerByID(c echo.Context) error
	GetAllStickers(c echo.Context) error
}

type delivery struct {
	stickerUsecase stickersUsecase.StickerUseCaseI
}

// GetStickerByID godoc
// @Summary      Get sticker by id
// @Description  Get sticker by id
// @Tags     	 stickers
// @Param id path int  true  "sticker ID"
// @Produce  Attachment/png
// @Success  200 "success get sticker"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Router   /sticker/{id} [get]
func (delivery *delivery) GetStickerByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	sticker, err := delivery.stickerUsecase.GetStickerByID(id)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	f, err := os.Open("stickers/" + sticker.Link)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	return c.Stream(http.StatusOK, "Attachment/png", f)
}

// GetAllStickers godoc
// @Summary      Get all stickers
// @Description  Get all stickers
// @Tags     	 stickers
// @Produce  application/json
// @Success  200 {object} pkg.Response{body=[]models.Sticker} "success get sticker"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Router   /stickers [get]
func (delivery *delivery) GetAllStickers(c echo.Context) error {
	stickers, err := delivery.stickerUsecase.GetAllStickers()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: stickers})
}

func NewDelivery(e *echo.Echo, su stickersUsecase.StickerUseCaseI) {
	handler := &delivery{
		stickerUsecase: su,
	}

	e.GET("/sticker/:id", handler.GetStickerByID)
	e.GET("/stickers", handler.GetAllStickers)
}
