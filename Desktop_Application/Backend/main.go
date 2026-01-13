package main

import (
	Api "example/backend/api"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("testing backend!")

	http.HandleFunc("/credentials", Api.CredentialsHandler)
	http.HandleFunc("/logout", Api.LogoutHandler)
	http.HandleFunc("/search", Api.SearchHandler)
	http.HandleFunc("/sse", Api.SSEHandler)
	http.HandleFunc("/call", Api.CallHandler)

	http.ListenAndServe("localhost:5000", nil)

}
