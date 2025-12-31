package main

import (
	Api "example/backend/api"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("testing backend!")

	http.HandleFunc("/credentials", Api.CredentialsHandler)

	http.ListenAndServe(":5000", nil)

}
