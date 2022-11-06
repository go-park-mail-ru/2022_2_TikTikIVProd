package delivery

import (
	chatUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/chat/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models/chat/dto"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DeliveryI interface {
	CreateDialog(c echo.Context) error
}

type delivery struct {
	cUsecase chatUsecase.ChatUsecaseI
}

func (delivery *delivery) CreateDialog(c echo.Context) error {
	request := new(dto.CreateDialogRequest)

	if err := c.Bind(&request); err != nil {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	userId, ok := c.Get("user_id").(int)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	request.UserID = userId
	response := new(dto.CreateDialogResponse)

	if err := delivery.cUsecase.CreateDialog(request, response); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: response})
}

func (delivery *delivery) GetDialog(c echo.Context) error {
	panic("")
}

func (delivery *delivery) GetDialogsInfo(c echo.Context) error {
	panic("")
}

func (delivery *delivery) WsChatHandler(c echo.Context) error {
	panic("")
}

func NewDelivery(e *echo.Echo, cu chatUsecase.ChatUsecaseI) {
	handler := &delivery{
		cUsecase: cu,
	}

	e.POST("/chat/create", handler.CreateDialog)
	//e.GET("/chat/dialog/:id", handler.GetDialog)
	//e.GET("/chat/user/:id/dialogs", handler.GetDialogsInfo)
	//e.GET("/chat/ws", handler.WsChatHandler)
}
