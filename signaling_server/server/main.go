package main

import (
	"fmt"
	"net/http"

	Api "github.com/omr50/P2P_Video_Calling_App/internal/api"
	"github.com/omr50/P2P_Video_Calling_App/internal/auth"
)

func main() {
	fmt.Println("Starting Server!")
	Api.InitDB()

	http.HandleFunc("/login", auth.LoginHandler)
	http.HandleFunc("/signup", auth.SignupHandler)
	http.HandleFunc("/protected", auth.ProtectedHandler)
	http.ListenAndServe(":8090", nil)
}
