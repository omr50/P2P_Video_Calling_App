package main

import (
	"fmt"
	"net/http"

	"github.com/omr50/P2P_Video_Calling_App/internal/auth"
)

func main() {
	fmt.Println("Starting Server!")
	http.HandleFunc("/login", auth.LoginHandler)
	http.HandleFunc("/protected", auth.ProtectedHandler)

	http.ListenAndServe(":8090", nil)

}
