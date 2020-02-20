package service

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (s *service) commentHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	client := NewClient(ws)
	defer func() {
		client.Close()
		s.removeClient(client)
	}()
	s.addClient(client)
	for {
		msg, err := client.ReadMessage()
		if err != nil {
			break
		}
		for c := range s.clients {
			if err := c.WriteMessage(msg); err != nil {
				s.removeClient(c)
			}
		}
	}
}
