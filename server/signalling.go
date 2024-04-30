package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// AllRooms is the global hashmap for the server
var AllRooms RoomMap

// CreateRoomRequest Create a Room and returns roomID
func CreateRoomRequest(w http.ResponseWriter, r *http.Request) {
	roomID := AllRooms.CreateRoom()

	type resp struct {
		RoomID string `json:"room_id"`
	}

	log.Println(AllRooms.Map)
	json.NewEncoder(w).Encode(resp{RoomID: roomID})
}

// JoinRoomRequest will join client in a particular room
func JoinRoomRequest(w http.ResponseWriter, r *http.Request) {
	
}
