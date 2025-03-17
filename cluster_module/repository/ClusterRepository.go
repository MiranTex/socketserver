package repository

import (
	models "socket_project/cluster_module/Models"

	"gorm.io/gorm"
)

type ToModelCovertable[T any] interface {
	ToModel() T
}

func SaveCluster(dbConnection *gorm.DB, cluster ToModelCovertable[models.Cluster]) int {

	clusterModel := cluster.ToModel()

	result := dbConnection.Create(&clusterModel)

	if result.Error != nil {
		return -1
	}

	return clusterModel.ID

}
