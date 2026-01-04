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
	http.HandleFunc("/serach", Api.SearchHandler)

	http.ListenAndServe(":5000", nil)

}
