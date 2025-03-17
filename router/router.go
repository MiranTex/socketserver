package router

import (
	"socket_project/socket"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
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
