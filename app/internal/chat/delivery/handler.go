package delivery

import (
	chatUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/chat/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/middleware"
	"github.com/labstack/echo/v4"
)

type delivery struct {
	cUsecase chatUsecase.CharUsecaseI
}

func (delivery *delivery) CreateDialog(c echo.Context) error {
	panic("")
}

func (delivery *delivery) GetDialog(c echo.Context) error {
	panic("")
}

func (delivery *delivery) GetUserDialogsInfo(c echo.Context) error {
	panic("")
}

func (delivery *delivery) WsChatHandler(c echo.Context) error {
	panic("")
}

func NewDelivery(e *echo.Echo, cu chatUsecase.CharUsecaseI, authMid *middleware.Middleware) {
	handler := &delivery{
		cUsecase: cu,
	}

	e.POST("chat/create", handler.CreateDialog, authMid.Auth)
	e.GET("chat/dialog/:id", handler.GetDialog, authMid.Auth)
	e.GET("chat/user/:id/dialogs", handler.GetUserDialogsInfo, authMid.Auth)
	e.GET("chat/ws", handler.WsChatHandler, authMid.Auth)

}
