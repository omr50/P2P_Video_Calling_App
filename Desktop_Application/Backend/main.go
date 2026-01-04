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

	http.ListenAndServe(":5000", nil)

}
