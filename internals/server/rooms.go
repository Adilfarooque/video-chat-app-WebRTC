package server

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Participant describes a single entity in the hashmap
type Participants struct {
	Host bool
	ID   string
	Conn *websocket.Conn
}

// RooMap is the main hashmap [roomID string -> [[]participant]
type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participants
}

// Initialize the RoomMap struct
func (r *RoomMap) Init() {
	r.Map = make(map[string][]Participants)
}

// Get will return the array of participants in the room
func (r *RoomMap) Get(roomID string) []Participants {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	return r.Map[roomID]
}

// CreateRoom generate a unique roomID and return it -> insert it in the hashmap
func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	rand.New(rand.NewSource(time.Now().UnixNano()))
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 8)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomID := string(b)
	r.Map[roomID] = []Participants{}

	return roomID
}

// InsertIntoRoom will create a participant and add it in the hashmap
func (r *RoomMap) InsertIntoRoom(roomID string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	ClientID := uuid.New().String()
	incomingParticipant := Participants{host, ClientID, conn}

	log.Println("Inserting into Room with RoomID: ", roomID)
	r.Map[roomID] = append(r.Map[roomID], incomingParticipant)
}

// DeleteRoom deletes the room with the roomID
func (r *RoomMap) RemoveRoom(roomID string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Map, roomID)
}
