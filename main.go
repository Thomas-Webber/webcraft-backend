package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var port = flag.String("port", "8888", "service port")
var env = flag.String("env", "dev", "service port")
var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
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

func main() {
	flag.Parse()
	log.SetFlags(0)

	if *env == "dev" {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	}

	http.HandleFunc("/echo", echo)
	addr := "localhost:"+*port
	log.Println("Starting server at: " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
