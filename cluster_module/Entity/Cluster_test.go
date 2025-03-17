package entity_test

import (
	entity "socket_project/cluster_module/Entity"
	"testing"
)

func TestClusterCreation(t *testing.T) {

	cluster := entity.NewCluster(1, "Cluster 1", nil, "token", false, "owner")

	if cluster.ID() != 1 {
		t.Errorf("Expected cluster ID to be 1, got %d", cluster.ID())
	}

}
