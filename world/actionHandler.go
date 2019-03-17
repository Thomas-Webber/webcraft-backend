package world

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

//const ZoneLength = 1023

// World of the game
var World = make(map[int32]int32)

// Upgrader for gorilla
var Upgrader = websocket.Upgrader{} // use default options

// ActionMessage struct of json
type ActionMessage struct {
	ActionType string
	ZoneID     int
	PosX       int32
	PosY       int32
	PosZ       int32
	Color      int32
	Message    string
}

// EncodeXYZtoInt -
func EncodeXYZtoInt(x int32, y int32, z int32) int32 {
	return (x << 20) + (y << 10) + z
}

// AddBlock -
func AddBlock(jsonByte []byte) error {
	var actionMessage ActionMessage
	if err := json.Unmarshal(jsonByte, &actionMessage); err != nil {
		return err
	}
	World[EncodeXYZtoInt(actionMessage.PosX, actionMessage.PosY, actionMessage.PosZ)] = actionMessage.Color
	return nil
}

// ActionHandler -
func ActionHandler(w http.ResponseWriter, r *http.Request) {
	c, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
