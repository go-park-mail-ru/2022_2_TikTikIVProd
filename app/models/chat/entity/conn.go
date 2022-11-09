package entity

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

const (
	// Time allowed to write a Message_ to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong Message_ from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum Message_ size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// connection is an middleman between the websocket connection and the Hub.
type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan Message
}

// readPump pumps messages from the websocket connection to the Hub.
func (s Subscription) readPump(hub *Hub) {
	c := s.conn
	defer func() {
		hub.unregister <- s
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		msg := Message{
			ChatID:    s.room,
			CreatedAt: time.Now(),
		}

		err := c.ws.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}

		//TODO сохранение в базе
		hub.broadcast <- msg
	}
}

// write writes a Message_ with the given Message_ type and payload.
func (c *connection) write(msg Message) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteJSON(msg)
}

func (c *connection) writeType(mt int) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, []byte{})
}

// writePump pumps messages from the Hub to the websocket connection.
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

// serveWs handles websocket requests from the peer.
func ServeWs(c echo.Context, roomId int, hub *Hub) {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	conn := &connection{send: make(chan Message, 256), ws: ws}
	sub := Subscription{conn, roomId}
	hub.register <- sub
	go sub.writePump(hub)
	go sub.readPump(hub)
}
