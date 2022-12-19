package models

import (
	chatUseCase "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/internal/chat/usecase"
	models "github.com/go-park-mail-ru/2022_2_TikTikIVProd/MainApp/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type connection struct {
	ws   *websocket.Conn
	send chan models.Message
}

func (s Subscription) readPump(hub *Hub, cu chatUseCase.UseCaseI) {
	c := s.conn
	defer func() {
		hub.unregister <- s
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		msg := models.Message{
			DialogID:  s.room,
			CreatedAt: time.Now(),
		}

		err := c.ws.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		go cu.SendMessage(&msg)
		hub.broadcast <- msg
	}
}

func (c *connection) write(msg models.Message) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteJSON(msg)
}

func (c *connection) writeType(mt int) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, []byte{})
}

func (s *Subscription) writePump(hub *Hub) {
	c := s.conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.writeType(websocket.CloseMessage)
				return
			}
			if err := c.write(message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.writeType(websocket.PingMessage); err != nil {
				return
			}
		}
	}
}

func ServeWs(c echo.Context, roomId uint64, hub *Hub, cu chatUseCase.UseCaseI) {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	conn := &connection{send: make(chan models.Message, 256), ws: ws}
	sub := Subscription{conn, roomId}
	hub.register <- sub
	go sub.writePump(hub)
	go sub.readPump(hub, cu)
}
