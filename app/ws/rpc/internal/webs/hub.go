package webs

type Hub struct {
	Clients    map[int64]*Client
	channels   map[int64]map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	h := &Hub{
		Clients:    make(map[int64]*Client),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
	go h.run()
	return h
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.Clients[client.UserId] = client
		case client := <-h.unregister:
			if _, ok := h.Clients[client.UserId]; ok {
				delete(h.Clients, client.UserId)
			}
		}
	}
}
