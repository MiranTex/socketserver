package entity

import (
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

func NewCluster(id int, name string, clients []interfaces.ClusterClientInterface, access_token string, status bool, owner string) *Cluster {
	return &Cluster{
		id:           id,
		name:         name,
		clients:      clients,
		access_token: access_token,
		status:       status,
		owner:        owner,
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
