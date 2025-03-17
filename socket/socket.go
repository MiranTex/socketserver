package socket

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	Classes "socket_project/classes"
	entity "socket_project/cluster_module/Entity"
	"socket_project/factory"
	"socket_project/interfaces"
	"socket_project/utils"
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

	Clusters []entity.Cluster

	connections sync.Map
	eventChan   = make(chan Classes.Event, 100)
	ctx, cancel = context.WithCancel(context.Background())
	wg          sync.WaitGroup
)

func init() {
	go eventDispatcher()
}

func InitClusters(clusters []entity.Cluster) {
	Clusters = clusters
}

func Handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("New connection")
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

		cluster, err := getClusterByPublickKey(request.ClusterPublicId)
		fmt.Println(cluster)

		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte("Cluster not found"))
			break
		}

		if !cluster.AuthenticateonCluster(request.AccessToken) {
			conn.WriteMessage(websocket.TextMessage, []byte("Unauthorized"))

			break
		}

		switch request.RequestType {
		case "connection":

			fmt.Println("Connection request")

			client := factory.CreateClient(utils.GenerateUUID(), conn, request.Subscriptions)
			connections.Store(request.Id, client)

			client.SendEvent(Classes.CreateEvent(utils.GenerateUUID(), "connection", map[string]interface{}{"id": client.Id()}))

			wg.Add(1)
			go handleClientEvents(ctx, client)
		case "event":
			fmt.Println("Event request")
			event := Classes.CreateEvent(request.Id, request.EventType, request.EventMessage)
			eventChan <- event
		}
	}
}

func getClusterByPublickKey(publicKey string) (entity.Cluster, error) {
	for _, cluster := range Clusters {
		if cluster.PublicId() == publicKey {
			return cluster, nil
		}
	}
	return entity.Cluster{}, fmt.Errorf("cluster not found")
}

func handleClientEvents(ctx context.Context, client interfaces.ClientInterface) {

	fmt.Println("Handling client events")

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
