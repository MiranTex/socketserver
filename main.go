package main

import (
	"log"
	"net/http"

	facades "socket_project/cluster_module/Facades"
	r "socket_project/router"
	"socket_project/socket"
)

func main() {

	clusters := facades.GetAllClusters()

	socket.InitClusters(clusters)

	router := r.InitRouter()

	go router.Run(":8081")

	http.HandleFunc("/socket", socket.Handler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
