package delivery

import (
	"net/http"
	"strconv"

	chatUsecase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/internal/chat/usecase"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/models"
	ws "github.com/go-park-mail-ru/2022_2_TikTikIVProd/models/ws"
	"github.com/go-park-mail-ru/2022_2_TikTikIVProd/pkg"
	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type delivery struct {
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
func (delivery *delivery) GetDialog(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
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
func (delivery *delivery) GetAllDialogs(c echo.Context) error {
	userId, ok := c.Get("user_id").(int)
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
func (delivery *delivery) SendMessage(c echo.Context) error {
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

	requestSanitize(&message)

	userId, ok := c.Get("user_id").(int)
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

func (delivery *delivery) WsChatHandler(c echo.Context) error {
	roomID, err := strconv.Atoi(c.Param("roomId"))

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, "Chat not found") //TODO переделать на ошибки в файле
	}

	ws.ServeWs(c, roomID, delivery.hub, delivery.ChatUC)

	c.Response().Header().Set(echo.HeaderUpgrade, "websocket")
	return c.NoContent(http.StatusSwitchingProtocols)
}

func isRequestValid(message interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(message)
	if err != nil {
		return false, err
	}
	return true, nil
}

func requestSanitize(message *models.Message) {
	sanitizer := bluemonday.UGCPolicy()

	message.Body = sanitizer.Sanitize(message.Body)
}

func NewDelivery(e *echo.Echo, cu chatUsecase.UseCaseI) {
	hub := ws.NewHub()
	go hub.Run()
	handler := &delivery{
		ChatUC: cu,
		hub:    hub,
	}

	e.GET("/chat/:id", handler.GetDialog)
	e.GET("/chat", handler.GetAllDialogs)
	e.POST("/chat/send_message", handler.SendMessage)
	e.File("/room/:roomId", "app/cmd/index.html")
	e.GET("/ws/:roomId", handler.WsChatHandler)
}
