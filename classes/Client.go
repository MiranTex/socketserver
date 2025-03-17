package Classes

import (
	"log"
	"socket_project/utils"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	ClientBase
	mux *sync.Mutex
}

func NewClient(id string, conn *websocket.Conn, subscriptions []string) Client {

	mux := &sync.Mutex{}

	return Client{
		conn: conn,
		ClientBase: ClientBase{
			id:            id,
			Subscriptions: subscriptions,
		},
		mux: mux,
	}
}

func (c Client) SendEvent(event Event) (bool, error) {

	if !c.isSubscribedTo(event.GetEventType()) && event.GetEventType() != "connection" {
		return false, nil
	}

	if event.AmItheSender(c.id) {
		return false, nil
	}

	c.mux.Lock()
	defer c.mux.Unlock()

	if err := c.conn.WriteMessage(websocket.TextMessage, event.ToByteArray()); err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil
}

func (c Client) isSubscribedTo(eventType string) bool {
	return utils.Contains(c.Subscriptions, eventType)
}

func (c Client) GetConnection() *websocket.Conn {
	return c.conn
}

func (c Client) GetConnectionSatus() bool {

	// c.mux.Lock()
	// defer c.mux.Unlock()
	_, _, err := c.conn.ReadMessage()

	return err == nil
}
