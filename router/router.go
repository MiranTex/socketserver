package router

import (
	"html/template"
	entity "socket_project/cluster_module/Entity"
	facades "socket_project/cluster_module/Facades"
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

	router.GET("/socket/:event", func(c *gin.Context) {

		// event := Classes.CreateEvent("1", c.Param("event"), nil)
		// socket.SendUpdateSign(event)
	})

	router.GET("/socket/Connections", func(c *gin.Context) {

		cons := socket.GetConnections()

		var cons_ []string

		for _, con := range cons {
			cons_ = append(cons_, con.String())
		}

		c.JSON(200, gin.H{
			"connections": cons_,
		})

	})

	router.GET("/socket", func(c *gin.Context) {
		socket.Handler(c.Writer, c.Request)
	})

	return router
}
