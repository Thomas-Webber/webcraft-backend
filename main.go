package main

import (
	"flag"
	"github.com/Thomas-Webber/webcraft-backend/chat"
	"github.com/Thomas-Webber/webcraft-backend/world"
	"log"
	"net/http"
)

var port = flag.String("port", "8888", "service port")
var env = flag.String("env", "dev", "service port")

func main() {
	flag.Parse()
	log.SetFlags(0)
	if *env == "dev" {
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
	http.HandleFunc("/action", world.ActionHandler)
	http.HandleFunc("/reset", world.ResetHandler)

	addr := "localhost:" + *port
	log.Println("Starting server at: " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
