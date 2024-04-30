package server

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Participants struct {
	Host bool
	Conn *websocket.Conn
}


type RoomMap struct{
	Mutex sync.RWMutex
	Map map[string][]Participants
}
