package Api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type MessageType string

var GlobalUserEmail string
var GlobalClient *Client

const (
	MsgCreateSock   MessageType = "create_socket"
	MsgCallRequest  MessageType = "call_offer"
	MsgCallAccepted MessageType = "call_accepted"
	MsgCallDeclined MessageType = "call_declined"
	MsgKeyExchange  MessageType = "key_exchange"
)

type Message struct {
	Type           string
	SenderEmail    string
	RecipientEmail string
	Payload        json.RawMessage
}

func NewClient(serverURL string) (*Client, error) {

	header := http.Header{}
	header.Set("Authorization", "Bearer "+UserJWT)

	conn, _, err := websocket.DefaultDialer.Dial(serverURL, header)

	if err != nil {
		return nil, err
	}

	return &Client{
		email: GlobalUserEmail,
		conn:  conn,
	}, nil
}

func handleCallOffer(msg Message) {
	// send caller through channel
	fmt.Println("handle call offer message", msg)
	SSEChannel <- msg.SenderEmail
}

func WebsockClient() {

	// conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8090/ws", nil)
	var err error
	GlobalClient, err = NewClient("ws://localhost:8090/ws")
	if err != nil {
		log.Fatal("dial error:", err)
	}

	if err != nil {
		log.Fatal("write error:", err)
	}

	// read incoming msgs
	for {
		msgType, message, err := GlobalClient.conn.ReadMessage()
		if err != nil {
			log.Println("read errro:", err)
			return
		}

		if msgType == websocket.TextMessage {
			log.Println("received:", string(message))
			var msg Message
			json.Unmarshal(message, &msg)
			switch msg.Type {

			case "call_offer":
				handleCallOffer(msg)
			case "call_accepted":
				// handleAnswer(msg)
			case "ice":
				// handleICE(msg)
			case "call_end":
				// handleCallEnd(msg)
			}
		}
	}

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	GlobalClient.mu.Lock()
	err := GlobalClient.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "client exit"),
	)

	GlobalClient.mu.Unlock()
	if err != nil {
		log.Println("Write close error: ", err)
	}
	time.Sleep(100 * time.Millisecond)
	GlobalClient.conn.Close()
}
