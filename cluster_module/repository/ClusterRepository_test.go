package repository_test

import (
	"fmt"
	entity "socket_project/cluster_module/Entity"
	models "socket_project/cluster_module/Models"
	"socket_project/cluster_module/database"
	"socket_project/cluster_module/repository"
	"testing"
)

func TestClusterRepository(t *testing.T) {

	t.Run("Test Save Cluster", func(t *testing.T) {

		cluster := entity.NewCluster(1, "Cluster 1", nil, "token", false, "owner")

		fmt.Println(cluster)

		db := database.GetSqliteConnection()

		err := db.AutoMigrate(&models.Cluster{})

		if err != nil {
			fmt.Println(err)
			t.Errorf("Expected no error, got %s", err.Error())
		}

		result := repository.SaveCluster(db, cluster)

		if result != 1 {
			t.Errorf("Expected cluster ID to be 1, got %d", result)
		}

		//delete database file
		db.Migrator().DropTable(&models.Cluster{})

	})

}
