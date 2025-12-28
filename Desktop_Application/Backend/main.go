package main

import (
	Api "example/backend/api"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("testing backend!")
	go Api.WebsockClient()

	http.HandleFunc("/credentials", Api.CredentialsHandler)

	http.ListenAndServe(":5000", nil)

}
