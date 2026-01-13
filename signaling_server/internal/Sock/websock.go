package sock

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/omr50/P2P_Video_Calling_App/internal/api"
	"github.com/omr50/P2P_Video_Calling_App/internal/auth"
)

var (
	connections = make(map[string]*websocket.Conn)
	connMu      sync.RWMutex
)

func addConnection(email string, conn *websocket.Conn) {
	connMu.Lock()
	connections[email] = conn
	connMu.Unlock()
	fmt.Println("added client connection to map")
}

func removeConnection(email string) {
	connMu.Lock()
	delete(connections, email)
	connMu.Unlock()
	fmt.Println("removed client connection from map")
}

func getConnection(email string) *websocket.Conn {
	var conn *websocket.Conn
	connMu.Lock()
	conn = connections[email]
	connMu.Unlock()
	return conn
}

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Type           string
	SenderEmail    string
	RecipientEmail string
	Payload        json.RawMessage
}

func clientPing(conn *websocket.Conn) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if err := conn.WriteControl(
			websocket.PingMessage,
			nil,
			time.Now().Add(5*time.Second),
		); err != nil {
			return
		}
	}
}

func handleCallOffer(conn *websocket.Conn, msg Message) {
	// server finds other socket, sends call request to it
	fmt.Println("got call offer message:", msg)
	recipientConn := getConnection(msg.RecipientEmail)

	if recipientConn != nil {
		// just forward the message
		data, err := json.Marshal(msg)

		if err != nil {
			return
		}
		recipientConn.WriteMessage(websocket.TextMessage, data)

	} else {
		// write back to client that the call is declined
		// works in either case, no pickup or user not online
		emptyPayload := json.RawMessage(`{}`)
		declinedMessage := Message{
			Type:           "call_declined",
			SenderEmail:    msg.SenderEmail,
			RecipientEmail: msg.RecipientEmail,
			Payload:        emptyPayload,
		}

		jsonData, err := json.Marshal(declinedMessage)
		if err != nil {
			return
		}
		conn.WriteMessage(websocket.TextMessage, jsonData)
	}

}

func WebsockHandler(w http.ResponseWriter, r *http.Request) {
	authValue := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authValue, "Bearer ")

	claims, err := auth.VerifyToken(token)

	if err != nil {
		fmt.Println("Invalid token")
		return
	}

	conn, err := Upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("Failed to upgrade connection: %v\n", err)
		return
	}

	conn.SetReadDeadline(time.Now().Add(60 * time.Second))

	ctx := context.Background()

	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		api.RedisClient.Expire(ctx, claims.Email, 360*time.Second)
		return nil
	})

	addConnection(claims.Email, conn)

	api.RedisClient.HSet(ctx, claims.Email, map[string]interface{}{
		"server_id": "server-1",
		"last-seen": time.Now().Unix(),
	})
	err = api.MarkUserOnline(ctx, api.RedisClient, claims.Email)
	if err != nil {
		fmt.Println("Error marking user online in redis")
		return
	}

	// api.RedisClient.Expire(ctx, claims.Email, 5*time.Minute)
	err = api.SetClientExpiration(ctx, api.RedisClient, claims.Email, 360*time.Second)
	if err != nil {
		fmt.Println("Error setting client expiration in redis")
		return
	}
	defer conn.Close()

	go clientPing(conn)

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			api.MarkUserOffline(ctx, api.RedisClient, claims.Email)
			removeConnection(claims.Email)
			log.Printf("Websocket closed or read error: %v\n", err)
			return
		}
		fmt.Println("Message type: ", messageType)
		var msg Message
		json.Unmarshal(message, &msg)

		fmt.Println("message: ", msg)

		switch msg.Type {

		case "call_offer":
			handleCallOffer(conn, msg)
		case "call_accepted":
			// handleAnswer(msg)
		case "ice":
			// handleICE(msg)
		case "call_end":
			// handleCallEnd(msg)
		}
	}

}
