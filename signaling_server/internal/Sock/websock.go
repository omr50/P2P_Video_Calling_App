package sock

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Type    string
	From    string
	To      string
	Payload json.RawMessage
}

func handleCallOffer(conn *websocket.Conn, msg Message) {
	// server finds other socket, sends call request to it
	fmt.Println("got call offer message:", msg)

}

func WebsockHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("Failed to upgrade connection: %v\n", err)
		return
	}

	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Failed to upgrade connection: %v\n", err)
			return
		}
		fmt.Println("Message type: ", messageType)
		var msg Message
		json.Unmarshal(message, &msg)

		fmt.Println("message: ", msg)

		switch msg.Type {

		case "call_offer":
			handleCallOffer(conn, msg)
		case "answer":
			// handleAnswer(msg)
		case "ice":
			// handleICE(msg)
		case "call_end":
			// handleCallEnd(msg)
		}
	}

}
