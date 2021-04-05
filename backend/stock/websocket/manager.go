package websocket

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
)

var manager *ClientManager

// ClientManager ClientManager
type ClientManager struct {
	sync.Mutex
	Clients map[*Client]bool
}

// Init Init
func Init() error {
	manager = &ClientManager{
		Clients: make(map[*Client]bool),
	}

	return nil
}

// GetManager GetManager
func GetManager() *ClientManager {
	return manager
}

// Register Register
func (m *ClientManager) Register(client *Client) {
	m.Lock()
	defer m.Unlock()
	m.Clients[client] = true
}

// Unregister Unregister
func (m *ClientManager) Unregister(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.Clients[client]; ok {
		log.Printf("Unregister: %v", client.ID)
		client.Lock()
		client.Cancel()
		close(client.ReadMessage)
		close(client.WriteMessage)
		delete(m.Clients, client)
		client = nil
	}
}

// Broadcast Broadcast
func (m *ClientManager) Broadcast(message []byte) {
	m.Lock()
	defer m.Unlock()
	for client := range m.Clients {
		select {
		case client.WriteMessage <- message:
		default:
			client.Lock()
			close(client.ReadMessage)
			close(client.WriteMessage)
			delete(m.Clients, client)
			client = nil
		}
	}
}

// GetClient GetClient
func (m *ClientManager) GetClient(id string) (client *Client, err error) {
	m.Lock()
	defer m.Unlock()
	for client := range m.Clients {
		if client.ID == id {
			return client, nil
		}
	}

	return nil, errors.New("client not found")
}

// SendMessage send ws message to ws client
func (m *ClientManager) SendMessage(clientID string, message interface{}) {
	m.Lock()
	defer m.Unlock()
	jsonMessage, _ := json.Marshal(message)
	for client := range m.Clients {
		if client.ID == clientID {
			client.Send(jsonMessage)
		}
	}
}
