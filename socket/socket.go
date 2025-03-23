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

	Clusters *entity.ClusterList

	ctx, cancel = context.WithCancel(context.Background())
	wg          sync.WaitGroup
)

func init() {
	// Removed global event dispatcher
}

func InitClusters(clusters []*entity.Cluster) {
	Clusters = entity.NewClusterList(clusters...)
	for _, cluster := range clusters {
		cluster.InitEventChannel()
		go clusterEventDispatcher(cluster)
	}
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

		cluster, err := AuthClient(request.ClusterPublicId, request.AccessToken)

		if err != nil || cluster == nil {
			conn.WriteMessage(websocket.TextMessage, []byte("Unauthorized"))

			continue
		}

		switch request.RequestType {
		case "connection":

			client := factory.CreateClient(utils.GenerateUUID(), conn, request.Subscriptions)

			cluster.AddClient(client)

			client.SendEvent(Classes.CreateEvent(utils.GenerateUUID(), "connection", map[string]interface{}{"id": client.Id()}))

			wg.Add(1)
			go handleClientEvents(ctx, client, cluster)
		case "event":
			event := Classes.CreateEvent(request.Id, request.EventType, request.EventMessage)
			cluster.EventChan <- event
		}
	}
}

func handleClientEvents(ctx context.Context, client interfaces.ClientInterface, cluster *entity.Cluster) {

	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if !client.GetConnectionSatus() {
				cluster.RemoveClient(client)
				return
			}
		}
	}
}

func clusterEventDispatcher(cluster *entity.Cluster) {
	for {
		select {
		case <-ctx.Done():
			return
		case event := <-cluster.EventChan:

			cluster.GetConnections().Range(func(key, value interface{}) bool {
				client := value.(interfaces.ClientInterface)
				if success, err := client.SendEvent(event); !success || err != nil {
					// log.Printf("Error sending event to client %s\n", client.Id())
				}
				return true
			})
		}
	}
}

func GetConnections(clusterId string) *sync.Map {

	cluster, error := Clusters.GetClusterByPublickKey(clusterId)

	if error != nil {
		return nil
	}

	connections := cluster.GetConnections()

	return connections
}

func AuthClient(clusterPId string, access_token string) (*entity.Cluster, error) {

	cluster, err := Clusters.GetClusterByPublickKey(clusterPId)

	if err != nil {
		return nil, err
	}

	if !cluster.AuthenticateonCluster(access_token) {
		return nil, fmt.Errorf("Unauthorized")
	}

	return cluster, nil

}

func Shutdown() {
	cancel()
	wg.Wait()
}
