package world

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const ZoneLength = 100

// World of the game
var World = make(map[int64]int32)

// Upgrader for gorilla
var Upgrader = websocket.Upgrader{} // use default options

// ActionMessage struct of json
type ActionMessage struct {
	ActionType string
	ZoneId int
	PosX int
	PosY int
	PosZ int
	Color int32
	Message string
}

func encodeXYZtoInt(x int, y int, z int) int64 {
	return int64((x << 42) + (y << 21) + z)
}

func AddBlock(jsonByte []byte) error {
	var actionMessage ActionMessage
	if err := json.Unmarshal(jsonByte, &actionMessage); err != nil {
		return err
	}
	World[encodeXYZtoInt(actionMessage.PosX, actionMessage.PosY, actionMessage.PosZ)] = actionMessage.Color
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
