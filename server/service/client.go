package service

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

var once sync.Once

// Client represents user client
type Client struct {
	ws        *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte
	isClosed  bool
	mu        sync.Mutex
}

func NewClient(ws *websocket.Conn) *Client {
	client := &Client{
		ws:        ws,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
	}

	go client.readLoop()
	go client.writeLoop()

	return client
}

func (c *Client) ReadMessage() (msg []byte, err error) {
	select {
	case msg = <-c.inChan:
	case <-c.closeChan:
		err = errors.New("Client disconnection")
	}
	return
}

func (c *Client) WriteMessage(msg []byte) (err error) {
	select {
	case c.outChan <- msg:
	case <-c.closeChan:
		err = errors.New("Client disconnected")
	}
	return
}

func (c *Client) Close() {
	c.ws.Close()
	c.mu.Lock()
	if !c.isClosed {
		close(c.closeChan)
		c.isClosed = true
	}
	c.mu.Unlock()
}

func (c *Client) readLoop() {
	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			goto ERR
		}

		select {
		case c.inChan <- msg:
		case <-c.closeChan:
			goto ERR
		}
	}
ERR:
	c.Close()
}

func (c *Client) writeLoop() {
	var msg []byte
	for {
		select {
		case msg = <-c.outChan:
		case <-c.closeChan:
			goto ERR
		}
		if err := c.ws.WriteMessage(websocket.TextMessage, msg); err != nil {
			goto ERR
		}
	}
ERR:
	c.Close()
}
