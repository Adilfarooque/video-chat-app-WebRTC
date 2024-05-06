package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// AllRooms is the global hashmap for the server
var AllRooms RoomMap

type resp struct {
	RoomID string `json:"room_id"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// CreateRoomRequest Create a Room and returns roomID
func CreateRoomRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	roomID := AllRooms.CreateRoom()

	log.Println(AllRooms.Map)
	json.NewEncoder(w).Encode(resp{RoomID: roomID})
}

type broadcastMsg struct {
	Message map[string]interface{}
	RoomID  string
	Client  *websocket.Conn
}

var broadcast = make(chan broadcastMsg)

func broadcaster() {
	for {
		msg := <-broadcast

		for _, client := range AllRooms.Map[msg.RoomID] {
			if client.Conn != msg.Client {
				err := client.Conn.WriteJSON(msg.Message)
				if err != nil {
					log.Fatal(err)
					client.Conn.Close()
				}
			}
		}
	}
}

// JoinRoomRequest will join client in a particular room
func JoinRoomRequest(w http.ResponseWriter, r *http.Request) {
	query_roomID, ok := r.URL.Query()["roomID"]

	if !ok {
		log.Println("roomID missing in unable to join the call")
		return
	}

	roomID := query_roomID[0]

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal("Web socket upgrade Error", err)
	}

	AllRooms.InsertIntoRoom(roomID, false, ws)

	go broadcaster()

	for {
		var msg broadcastMsg
		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Fatal("Read Error", err)
		}
		msg.Client = ws
		msg.RoomID = roomID

		broadcast <- msg
	}

}
