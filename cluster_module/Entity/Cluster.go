package entity

import (
	"encoding/json"
	models "socket_project/cluster_module/Models"
	"socket_project/cluster_module/interfaces"
)

type Cluster struct {
	id           int
	publicId     string
	name         string
	clients      []interfaces.ClusterClientInterface
	access_token string
	status       bool
	owner        string
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

	return &Cluster{
		id:           id,
		name:         name,
		clients:      clients,
		access_token: access_token,
		status:       status,
		owner:        owner,
		publicId:     publicId,
	}
}

func (c *Cluster) ID() int {
	return c.id
}

func (c *Cluster) PublicId() string {
	return c.publicId
}

func (c *Cluster) AddClient(client interfaces.ClusterClientInterface) {
	c.clients = append(c.clients, client)
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
