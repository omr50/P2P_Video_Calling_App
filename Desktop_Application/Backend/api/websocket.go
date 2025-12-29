package Api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type MessageType string

var GlobalClient *Client

const (
	MsgCreateSock   MessageType = "create_socket"
	MsgCallRequest  MessageType = "call_request"
	MsgCallAccepted MessageType = "call_accepted"
	MsgCallDeclined MessageType = "call_declined"
	MsgKeyExchange  MessageType = "key_exchange"
)

type Message struct {
	Type    string
	From    string
	To      string
	Payload json.RawMessage
}

func NewClient(serverURL string) (*Client, error) {

	header := http.Header{}
	header.Set("Authorization", "Bearer "+UserJWT)

	conn, _, err := websocket.DefaultDialer.Dial(serverURL, header)

	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

func WebsockClient() {

	// conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8090/ws", nil)
	var err error
	GlobalClient, err = NewClient("ws://localhost:8090/ws")
	if err != nil {
		log.Fatal("dial error:", err)
	}

	payload := map[string]interface{}{
		"text": "Hello World",
	}

	payloadBytes, _ := json.Marshal(payload)

	msg := Message{
		Type:    "call_offer",
		From:    GlobalClient.email,
		To:      "a@b",
		Payload: payloadBytes,
	}

	// defer GlobalClient.conn.Close()

	data, err := json.Marshal(msg)

	if err != nil {
		log.Fatal("Marshal error:", err)
	}

	go func() {
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			err = GlobalClient.conn.WriteMessage(websocket.TextMessage, data)
		}
	}()
	if err != nil {
		log.Fatal("write error:", err)
	}

	// read incoming msgs
	for {
		msgType, msg, err := GlobalClient.conn.ReadMessage()
		if err != nil {
			log.Println("read errro:", err)
			return
		}

		if msgType == websocket.TextMessage {
			log.Println("received:", string(msg))
		}
	}

}
