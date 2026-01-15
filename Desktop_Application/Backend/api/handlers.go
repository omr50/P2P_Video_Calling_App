package Api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

var SSEChannel = make(chan string)

type SearchedUser struct {
	Email      string
	Username   string
	Created_at string
	Online     bool
}

func CallHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Call handler!")
	email := r.URL.Query().Get("email")

	fmt.Println("call handler query: ", email)

	if email == "" {
		http.Error(w, "query email is empty", http.StatusBadRequest)
		return
	}

	JsonPayload := json.RawMessage(`{}`)
	message := Message{Type: string(MsgCallRequest), SenderEmail: GlobalClient.email, RecipientEmail: email, Payload: JsonPayload}
	fmt.Println("Call handler message", message)

	jsonMsg, err := json.Marshal(message)

	if err != nil {
		http.Error(w, "query email is empty", http.StatusBadRequest)
		return
	}

	GlobalClient.mu.Lock()
	GlobalClient.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	GlobalClient.mu.Unlock()

	fmt.Println("Sent initiate call websocket message successfully")
	w.WriteHeader(http.StatusOK)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Working?")
	query := r.URL.Query().Get("query")
	fmt.Println("query:", query)

	if len(query) < 3 {
		http.Error(w, "query too short or doesn't exist", http.StatusBadRequest)
		return
	}

	endpoint := "http://localhost:8090/user-search?query=" + url.QueryEscape(query)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+UserJWT)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "backend error", resp.StatusCode)
		return
	}

	var results []SearchedUser
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)

}

func SSEHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	origin := r.Header.Get("Origin")

	allowed := map[string]bool{
		"http://localhost:5173": true,
		"http://localhost:5174": true,
	}

	if allowed[origin] {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ctx := r.Context()

	for {
		select {
		case msg := <-SSEChannel:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			flusher.Flush()
		case <-ctx.Done():
			return
		}
	}
}
