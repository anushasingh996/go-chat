package main

import (
	"bufio"
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

var done chan interface{}
var interrupt chan os.Signal

func receiveHandler(connection *websocket.Conn) {
	defer close(done)
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println("Error in receive:", err)
			return
		}
		log.Printf("Received: %s\n", msg)
	}
}

func main() {
	done = make(chan interface{})
	interrupt = make(chan os.Signal)

	signal.Notify(interrupt, os.Interrupt)

	socketUrl := "ws://localhost:8001" + "/ws"
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}
	defer conn.Close()
	go receiveHandler(conn)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		err := conn.WriteMessage(websocket.TextMessage, []byte(scanner.Text()))
		if err != nil {
			log.Println("Error during writing to websocket:", err)
			return
		}
		_, ok := <-interrupt
		if ok {
			log.Println("Received SIGINT interrupt signal. Closing all pending connections")

			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error during closing websocket:", err)
				return
			}
			return
		}
	}

}
