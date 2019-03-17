package main

import (
	"github.com/Thomas-Webber/webcraft-backend/chat"
	"github.com/Thomas-Webber/webcraft-backend/world"
	"log"
	"net/http"
	"os"
)

func getEnv(name string, defaultValue string) string {
	val := os.Getenv(name)
	if len(val) == 0 {
		return defaultValue
	}
	return val
}

var port = getEnv("PORT", "8888")
var env = getEnv("ENV", "dev")

func main() {
	log.SetFlags(0)
	if env == "dev" {
		world.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		chat.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	}

	hub := chat.NewHub()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(hub, w, r)
	})

	world.World[world.EncodeXYZtoInt(15, 5, 5)] = 0xffffff
	world.World[world.EncodeXYZtoInt(5, 5, 5)] = 0xff0000
	http.HandleFunc("/get", world.GetHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Hello")); err != nil {
			log.Fatal(err)
		}
	})
	http.HandleFunc("/action", world.ActionHandler)
	http.HandleFunc("/reset", world.ResetHandler)

	addr := "localhost:" + *port
	log.Println("Starting server at: " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
