package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(mt, message)
	}
}

func main() {
	http.HandleFunc("/ws/", wsHandler)
	fmt.Println("WebSocket API listening on :8081")
	http.ListenAndServe(":8081", nil)
}
