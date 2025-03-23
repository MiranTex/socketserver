package repository

import (
	entity "socket_project/cluster_module/Entity"
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

func GetAllClusters(dbConnection *gorm.DB) []*entity.Cluster {

	var clustersModels []models.Cluster

	dbConnection.Find(&clustersModels)

	var clusters []*entity.Cluster

	for _, clusterModel := range clustersModels {
		cluster := entity.CreateFromModel(clusterModel)
		clusters = append(clusters, cluster)
	}

	return clusters

}
