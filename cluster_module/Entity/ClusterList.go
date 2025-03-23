package entity

import (
	"fmt"
	"sync"
)

type ClusterList struct {
	Clusters sync.Map
}

func NewClusterList(clusters ...*Cluster) *ClusterList {

	if len(clusters) > 0 {
		clusterList := &ClusterList{}
		for _, cluster := range clusters {
			clusterList.Append(cluster)
		}
		return clusterList
	}

	return &ClusterList{}
}

func (c *ClusterList) Append(cluster *Cluster) {
	c.Clusters.Store(cluster.PublicId(), cluster)
}

func (c *ClusterList) GetClusterByPublickKey(publicKey string) (*Cluster, error) {
	cluster, ok := c.Clusters.Load(publicKey)

	if ok {
		return cluster.(*Cluster), nil
	}

	return &Cluster{}, fmt.Errorf("cluster not found")
}
