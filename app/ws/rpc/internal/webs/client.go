package webs

import (
	"bytes"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	_writeWait      = 10 * time.Second
	_pongWait       = 60 * time.Second
	_pingPeriod     = (_writeWait * 9) / 10
	_maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type wsClientHandler func([]byte) error

type Client struct {
	UserId        int64
	Send          chan []byte
	hub           *Hub
	conn          *websocket.Conn
	lastHeartbeat time.Time
}

func NewClient(hub *Hub, userId int64, conn *websocket.Conn) *Client {
	return &Client{
		UserId: userId,
		Send:   make(chan []byte, 256),
		hub:    hub,
		conn:   conn,
	}
}

func (c *Client) readPump(handler wsClientHandler) {
	defer c.close()

	c.conn.SetReadLimit(_maxMessageSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(_pongWait))
	c.conn.SetPongHandler(func(string) error {
		return c.conn.SetReadDeadline(time.Now().Add(_pongWait))
	})
	for {
		typ, msg, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		if typ == websocket.PongMessage {
			c.lastHeartbeat = time.Now()
			continue
		}
		msg = bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))
		if err := handler(msg); err != nil {
			logx.Error(err)
		}
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(_pingPeriod)
	defer ticker.Stop()
	defer c.close()

	var gerr error
	for {
		if gerr != nil {
			break
		}
		select {
		case msg, ok := <-c.Send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				break
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				gerr = err
				break
			}
			w.Write(msg)

			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Send)
			}
			if err := w.Close(); err != nil {
				gerr = err
				break
			}
			logx.Infof("sended to uesr %d", c.UserId)
		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(_writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				gerr = err
				break
			}
		}
	}
	logx.Errorf("user %d: %v", c.UserId, gerr)
}

func (c *Client) close() {
	c.hub.unregister <- c
	_ = c.conn.WriteMessage(websocket.CloseMessage, nil)
	_ = c.conn.Close()
	logx.Errorf("user %d websocket is disconnected", c.UserId)
}
