package main

import (
	"fmt"
	"net/http"

	"github.com/omr50/P2P_Video_Calling_App/internal/api"
	"github.com/omr50/P2P_Video_Calling_App/internal/sock"

	"github.com/omr50/P2P_Video_Calling_App/internal/auth"
)

func main() {
	fmt.Println("Starting Server!")
	api.InitDB()
	api.InitRedisClient()

	http.HandleFunc("/login", auth.LoginHandler)
	http.HandleFunc("/signup", auth.SignupHandler)
	http.HandleFunc("/protected", auth.ProtectedHandler)
	http.HandleFunc("/user-search", api.EmailSearch)
	http.HandleFunc("/ws", sock.WebsockHandler)
	http.ListenAndServe(":8090", nil)
}
