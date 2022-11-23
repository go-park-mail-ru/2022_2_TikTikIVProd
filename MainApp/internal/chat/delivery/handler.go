package delivery

import (
	"net/http"
	"strconv"

	chatUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/chat/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	ws "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models/ws"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/pkg"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type Delivery struct {
	ChatUC chatUsecase.UseCaseI
	hub    *ws.Hub
}

// GetDialog godoc
// @Summary      GetDialog
// @Description  get dialog
// @Tags     chat
// @Produce  application/json
// @Param id path int true "Chat ID"
// @Success  200 {object} pkg.Response{body=models.Dialog} "success get dialog"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 404 {object} echo.HTTPError "can't find chat with such id"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /chat/{id} [get]
func (delivery *Delivery) GetDialog(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	dialog, err := delivery.ChatUC.SelectDialog(id)
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
	return c.JSON(http.StatusOK, pkg.Response{Body: dialog})
}

// GetDialogByUsers godoc
// @Summary      GetDialogByUsers
// @Description  get dialog
// @Tags     chat
// @Produce  application/json
// @Param id path int true "Friend ID"
// @Success  200 {object} pkg.Response{body=models.Dialog} "success get dialog"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 404 {object} echo.HTTPError "can't find chat with such id"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /chat/user/{id} [get]
func (delivery *Delivery) GetDialogByUsers(c echo.Context) error {
	friendId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	dialog, err := delivery.ChatUC.SelectDialogByUsers(userId, friendId)
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
	return c.JSON(http.StatusOK, pkg.Response{Body: dialog})
}

// GetAllDialogs godoc
// @Summary      GetAllDialogs
// @Description  get all dialogs
// @Tags     chat
// @Produce  application/json
// @Success  200 {object} pkg.Response{body=[]models.Dialog} "success get dialogs"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /chat [get]
func (delivery *Delivery) GetAllDialogs(c echo.Context) error {
	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	dialogs, err := delivery.ChatUC.SelectAllDialogs(userId)
	if err != nil {
		causeErr := errors.Cause(err)
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, causeErr.Error())
	}
	return c.JSON(http.StatusOK, pkg.Response{Body: dialogs})
}

// SendMessage godoc
// @Summary      SendMessage
// @Description  send message
// @Tags     chat
// @Accept	 application/json
// @Produce  application/json
// @Param    message body models.Message true "message data"
// @Success  200 {object} pkg.Response{body=models.Message} "success send message"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 404 {object} echo.HTTPError "can't find item with such id"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /chat/send_message [post]
func (delivery *Delivery) SendMessage(c echo.Context) error {
	var message models.Message
	err := c.Bind(&message)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	if ok, err := isRequestValid(&message); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error(models.ErrInternalServerError)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	message.SenderID = userId

	err = delivery.ChatUC.SendMessage(&message)
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
	return c.JSON(http.StatusOK, pkg.Response{Body: message})
}

func (delivery *Delivery) WsChatHandler(c echo.Context) error {
	roomID, err := strconv.ParseUint(c.Param("roomId"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error()) //TODO переделать на ошибки в файле
	}

	ws.ServeWs(c, roomID, delivery.hub, delivery.ChatUC)

	return c.NoContent(http.StatusOK)
}

func isRequestValid(message interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(message)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewDelivery(e *echo.Echo, cu chatUsecase.UseCaseI) {
	hub := ws.NewHub()
	go hub.Run()
	handler := &Delivery{
		ChatUC: cu,
		hub:    hub,
	}

	e.GET("/chat/:id", handler.GetDialog)
	e.GET("/chat/user/:id", handler.GetDialogByUsers)
	e.GET("/chat", handler.GetAllDialogs)
	e.POST("/chat/send_message", handler.SendMessage)
	e.File("/room/:roomId", "app/cmd/index.html")
	e.GET("/ws/:roomId", handler.WsChatHandler)
}
