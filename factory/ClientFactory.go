package factory

import (
	Classes "socket_project/classes"
	"socket_project/interfaces"

	"github.com/gorilla/websocket"
)

func CreateClient(id string, Conn *websocket.Conn, subscriptions []string) interfaces.ClientInterface {

	client := Classes.NewClient(id, Conn, subscriptions)

	return &client
}

func CreateHttpClient(id string, notifyUrl string, subscriptions []string) interfaces.ClientInterface {

	client := Classes.NewHttpClient(id, notifyUrl, subscriptions)

	return &client
}
