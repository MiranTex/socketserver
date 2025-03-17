package socket

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	Classes "socket_project/classes"
	"socket_project/factory"
	"socket_project/interfaces"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	connections sync.Map
	eventChan   = make(chan Classes.Event, 100)
	ctx, cancel = context.WithCancel(context.Background())
	wg          sync.WaitGroup
)

func init() {
	go eventDispatcher()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var request Classes.SocketRequest
		if err := json.Unmarshal(p, &request); err != nil {
			log.Println(err)
			continue
		}

		switch request.RequestType {
		case "connection":
			client := factory.CreateClient(request.Id, conn, request.Subscriptions)
			connections.Store(request.Id, client)
			wg.Add(1)
			go handleClientEvents(ctx, client)
		case "event":
			event := Classes.CreateEvent(request.Id, request.EventType, request.EventMessage)
			eventChan <- event
		}
	}
}

func handleClientEvents(ctx context.Context, client interfaces.ClientInterface) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if !client.GetConnectionSatus() {
				connections.Delete(client.Id())
				return
			}
		}
	}
}

func eventDispatcher() {
	for {
		select {
		case <-ctx.Done():
			return
		case event := <-eventChan:
			connections.Range(func(key, value interface{}) bool {
				client := value.(interfaces.ClientInterface)
				if success, err := client.SendEvent(event); !success || err != nil {
					log.Printf("Error sending event to client %s\n", client.Id())
				}
				return true
			})
		}
	}
}

func GetConnections() []interfaces.ClientInterface {
	var clients []interfaces.ClientInterface
	connections.Range(func(_, value interface{}) bool {
		clients = append(clients, value.(interfaces.ClientInterface))
		return true
	})
	return clients
}

func Shutdown() {
	cancel()
	close(eventChan)
	wg.Wait()
}
