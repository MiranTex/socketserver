package facades

import (
	"fmt"
	entity "socket_project/cluster_module/Entity"
	"socket_project/cluster_module/database"
	"socket_project/cluster_module/repository"
	"socket_project/utils"

	"github.com/gin-gonic/gin"
)

func CreateCluster(c *gin.Context) int {

	type request struct {
		Name        string `json:"name"`
		Owner       string `json:"owner"`
		AccessToken string `json:"access_token"`
	}

	var req request

	if err := c.ShouldBindJSON(&req); err != nil {

		return -1
	}

	name := req.Name
	owner := req.Owner
	access_token := req.AccessToken

	fmt.Println("name: "+name, "owner: "+owner, "access_token: "+access_token)

	cluster := entity.NewCluster(0, name, nil, access_token, true, owner, utils.GenerateUUID())

	db := database.GetSqliteConnection()

	newID := repository.SaveCluster(db, cluster)

	return newID

}

func GetAllClusters() []*entity.Cluster {

	db := database.GetSqliteConnection()

	clusters := repository.GetAllClusters(db)

	return clusters
}
