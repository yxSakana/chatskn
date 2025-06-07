package webs

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type ServeInput struct {
	HubCore *Hub
	Writer  http.ResponseWriter
	Req     *http.Request
	UserId  int64
	Handler wsClientHandler
}

func ServeWs(in ServeInput) {
	w := in.Writer
	r := in.Req
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	client := NewClient(in.HubCore, in.UserId, conn)
	client.hub.register <- client

	go client.writePump()
	go client.readPump(in.Handler)
}
