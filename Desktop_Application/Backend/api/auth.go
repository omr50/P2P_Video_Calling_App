package Api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var UserJWT string

type User struct {
	Email    string
	Password string
}

type TokenResponse struct {
	Token string `json:"token"`
}

func CredentialsHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(u)

	requestBody := bytes.NewBuffer(jsonData)

	// send this user to the signaling server to authenticate
	resp, err := http.Post("http://localhost:8090/login", "application/json", requestBody)

	if err != nil {
		http.Error(w, "signaling server unavailable", http.StatusBadGateway)
		return
	}

	var token TokenResponse

	err = json.NewDecoder(resp.Body).Decode(&token)

	json.NewEncoder(w).Encode(map[string]string{
		"token": token.Token,
	})

	fmt.Println("Token: ", token)
	UserJWT = token.Token
}
