package main

import (
	"log"
	"net/http"

	r "socket_project/router"
	"socket_project/socket"
)

func main() {

	router := r.InitRouter()

	go router.Run(":8081")

	http.HandleFunc("/socket", socket.Handler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
