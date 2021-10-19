package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func setupRoutes() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Simple Server")
	})
	http.HandleFunc("/ws", serveWs)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	go reader(ws)
	go writer(ws)
}

func reader(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(message))

		// if err := conn.WriteMessage(messageType, message); err != nil {
		//     log.Println(err)
		//     return
		// }

	}
}

func writer(conn *websocket.Conn) {
	if err := conn.WriteMessage(1, ([]byte)("message")); err != nil {
		log.Println(err)
		return
	}

}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
