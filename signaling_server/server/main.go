package main

import (
	"fmt"
	"net/http"

	"github.com/omr50/P2P_Video_Calling_App/internal/api"
	"github.com/omr50/P2P_Video_Calling_App/internal/sock"
)

func main() {
	fmt.Println("Starting Server!")
	api.InitDB()
	api.InitRedisClient()

	http.HandleFunc("/login", api.LoginHandler)
	http.HandleFunc("/signup", api.SignupHandler)
	http.Handle("/user-search", api.AuthMiddleware(http.HandlerFunc(api.EmailSearch)))
	http.HandleFunc("/ws", sock.WebsockHandler)
	http.ListenAndServe("localhost:8090", nil)
}
