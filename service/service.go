package service

import (
	"net/http"
	"sync"
)

// Service represents service
type Service interface {
	Run() error
}

// NewService will create service instance
func NewService(address string) Service {
	return &service{address: address, clients: make(map[*Client]bool, 0)}
}

type service struct {
	address string
	clients map[*Client]bool
	mu      sync.Mutex
}

func (s *service) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/comment", s.commentHandler)
	return http.ListenAndServe(s.address, mux)
}

func (s *service) addClient(client *Client) {
	s.mu.Lock()
	s.clients[client] = true
	s.mu.Unlock()
}

func (s *service) removeClient(client *Client) {
	s.mu.Lock()
	delete(s.clients, client)
	s.mu.Unlock()
}
