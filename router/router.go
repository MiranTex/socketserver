package router

import (
	"fmt"
	"html/template"
	Classes "socket_project/classes"
	entity "socket_project/cluster_module/Entity"
	facades "socket_project/cluster_module/Facades"
	"socket_project/interfaces"
	"socket_project/socket"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {

		template := template.Must(template.ParseFiles("index.template.html"))

		template.Execute(c.Writer, nil)

	})

	router.POST("/cluster", func(c *gin.Context) {

		newId := facades.CreateCluster(c)

		c.JSON(200, gin.H{
			"newId": newId,
		})
	})

	router.GET("/clusters", func(c *gin.Context) {

		clusters := facades.GetAllClusters()

		clustersJosn := []entity.ClusterJSON{}

		for _, cluster := range clusters {
			clustersJosn = append(clustersJosn, cluster.JSON())
		}

		c.JSON(200, gin.H{
			"clusters": clustersJosn,
		})

	})

	router.POST("/socket/:clusterPublicID", func(c *gin.Context) {

		httpRequest := &Classes.HttpRequest{}

		if err := c.ShouldBindJSON(httpRequest); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid request",
			})
			return
		}

		clusterPublicId := c.Param("clusterPublicID")
		access_token := httpRequest.AccessToken
		eventType := httpRequest.EventType
		eventData := httpRequest.EventData
		clientID := httpRequest.ClientId

		c.Bind(httpRequest)

		cluster, err := socket.Clusters.GetClusterByPublickKey(clusterPublicId)

		fmt.Println(cluster)

		if err != nil {
			c.JSON(404, gin.H{
				"error": "Cluster not found",
			})
			return
		}

		clientHasAccess := cluster.AuthenticateClient(access_token)

		if !clientHasAccess {
			c.JSON(401, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		event := Classes.CreateEvent(clientID, eventType, eventData)

		cluster.NotifyClients(event)

		c.JSON(200, gin.H{
			"message": "Event sent",
		})
	})

	router.GET("/socket/connections/:clusterid", func(c *gin.Context) {

		clusterid := c.Param("clusterid")

		cons := socket.GetConnections(clusterid)

		var cons_ []string

		cons.Range(func(key, value interface{}) bool {
			cons_ = append(cons_, value.(interfaces.ClientInterface).String())
			return true
		})

		c.JSON(200, gin.H{
			"connections": cons_,
		})

	})

	router.GET("/socket", func(c *gin.Context) {
		socket.Handler(c.Writer, c.Request)
	})

	return router
}
