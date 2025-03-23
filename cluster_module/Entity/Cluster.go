package entity

import (
	"encoding/json"
	"fmt"
	Classes "socket_project/classes"
	models "socket_project/cluster_module/Models"
	"socket_project/cluster_module/interfaces"
	"sync"
)

type Cluster struct {
	id           int
	publicId     string
	name         string
	clients      sync.Map
	access_token string
	status       bool
	owner        string
	EventChan    chan Classes.Event // Add EventChan field
}

type ClusterJSON struct {
	ID       int    `json:"id"`
	PublicID string `json:"publicId"`
	Name     string `json:"name"`
	// AccessToken  string `json:"access_token"`
	Status bool   `json:"status"`
	Owner  string `json:"owner"`
}

func NewCluster(id int, name string, clients []interfaces.ClusterClientInterface, access_token string, status bool, owner string, publicId string) *Cluster {

	cluster := &Cluster{
		id:           id,
		name:         name,
		clients:      sync.Map{},
		access_token: access_token,
		status:       status,
		owner:        owner,
		publicId:     publicId,
		EventChan:    make(chan Classes.Event, 1000), // Initialize EventChan
	}

	if clients != nil {
		for _, client := range clients {
			cluster.clients.Store(client.Id(), client)
		}
	}

	return cluster
}

func (c *Cluster) InitEventChannel() {
	c.EventChan = make(chan Classes.Event, 1000)
}

func (c *Cluster) ID() int {
	return c.id
}

func (c *Cluster) PublicId() string {
	return c.publicId
}

func (c *Cluster) AddClient(client interfaces.ClusterClientInterface) {
	c.clients.Store(client.Id(), client)
}

func (c *Cluster) RemoveClient(client interfaces.ClusterClientInterface) {
	c.clients.Delete(client.Id())
}

func (c *Cluster) StartCluster() {
	c.status = true
}

func (c *Cluster) StopCluster() {
	c.status = false
}

func (c *Cluster) ToModel() models.Cluster {
	return models.Cluster{
		PublicID:    c.publicId,
		Name:        c.name,
		AccessToken: c.access_token,
		IsPublic:    c.status,
		Owner:       c.owner,
	}
}

func CreateFromModel(clusterModel models.Cluster) *Cluster {
	return NewCluster(clusterModel.ID, clusterModel.Name, nil, clusterModel.AccessToken, clusterModel.IsPublic, clusterModel.Owner, clusterModel.PublicID)
}

func (c *Cluster) JSON() ClusterJSON {
	return ClusterJSON{
		ID:       c.id,
		PublicID: c.publicId,
		Name:     c.name,
		Status:   c.status,
		Owner:    c.owner,
	}
}

func (c *Cluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(ClusterJSON{
		ID:       c.id,
		PublicID: c.publicId,
		Name:     c.name,
		Status:   c.status,
		Owner:    c.owner,
	})
}

func (c *Cluster) AuthenticateonCluster(access_token string) bool {
	return c.access_token == access_token
}

func (c *Cluster) GetConnections() *sync.Map {
	return &c.clients
}

func (c *Cluster) AuthenticateClient(access_token string) bool {

	fmt.Println("Access token: ", access_token, "Cluster access token: ", c.access_token)
	return c.access_token == access_token
}

func (c *Cluster) NotifyClients(event Classes.Event) {
	c.clients.Range(func(key, value interface{}) bool {
		client := value.(interfaces.ClusterClientInterface)
		client.SendEvent(event)
		return true
	})
}
