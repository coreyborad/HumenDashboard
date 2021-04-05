package websocket

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Error constants
var (
	ErrUserNotExist = errors.New("User Not Exist")
	newline         = []byte{'\n'}
	space           = []byte{' '}
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

// Message Message
type Message struct {
	Type    string      `json:"type"`
	Message interface{} `json:"message"`
}

// Client Client
type Client struct {
	sync.Mutex
	ctx          context.Context
	Cancel       context.CancelFunc
	ID           string
	ws           *websocket.Conn
	WriteMessage chan []byte
	ReadMessage  chan []byte
}

// Clients Clients
type Clients []*Client

// NewClient NewClient
func NewClient(ctx *gin.Context) (client *Client, err error) {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return nil, err
	}

	c, cancel := context.WithCancel(ctx)
	client = &Client{
		ctx:          c,
		Cancel:       cancel,
		ID:           uuid.New().String(),
		ws:           ws,
		WriteMessage: make(chan []byte),
		ReadMessage:  make(chan []byte),
	}

	manager.Register(client)

	go client.ReadServe()
	go client.WriteServe()

	return
}

// Context Context
func (c *Client) Context() context.Context {
	if c.ctx == nil {
		ctx, cancel := context.WithCancel(context.Background())
		c.Lock()
		defer c.Unlock()

		c.Cancel = cancel
		c.ctx = ctx
	}

	return c.ctx
}

// SetUser SetUser
// func (c *Client) SetUser(user *models.User) *Client {
// 	c.Lock()
// 	defer c.Unlock()

// 	c.User = user

// 	return c
// }

// Send Send
func (c *Client) Send(message []byte) *Client {
	c.Lock()
	defer c.Unlock()

	c.WriteMessage <- message

	return c
}

// ReadServe ReadServe
func (c *Client) ReadServe() {
	defer func() {
		c.ws.Close()
		manager.Unregister(c)
	}()

	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error {
		c.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}

		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.ReadMessage <- message
	}
}

// WriteServe WriteServe
func (c *Client) WriteServe() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.WriteMessage:
			c.ws.SetWriteDeadline(time.Now().Add(writeWait))

			if !ok {
				c.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.ws.WriteMessage(websocket.TextMessage, message)

		case <-ticker.C:
			c.ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.ws.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
