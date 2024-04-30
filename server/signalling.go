package server

import (
	"fmt"
	"net/http"
)

//AllRooms is the global hashmap for the server
var AllRooms RoomMap

// CreateRoomRequest Create a Room and returns roomID
func CreateRoomRequest(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w,"Hello Create")
}

// JoinRoomRequest will join client in a particular room
func JoinRoomRequest(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w,"Hello Room")
}