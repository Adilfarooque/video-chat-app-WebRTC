package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Adilfarooque/video-chat-app/internals/server"
)

const (
	port = ":8000"
)

func main() {
	server.AllRooms.Init()
	http.HandleFunc("/create", server.CreateRoomRequest)
	http.HandleFunc("/joint", server.JoinRoomRequest)

	fmt.Println("Starting server on Port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
