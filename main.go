package main

import (
	"fmt"
	"net/http"

	"github.com/Adilfarooque/video-chat-app/server"
)

const (
	port = ":8080"
)

func main() {
	http.HandleFunc("/create", server.CreateRoomRequest)
	http.HandleFunc("/joint", server.JoinRoomRequest)

	fmt.Println("Starting the server ", port)
	http.ListenAndServe(port, nil)
}
