package main

import (
	Auth "example/backend/api"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("testing backend!")

	http.HandleFunc("/credentials", Auth.CredentialsHandler)

	http.ListenAndServe(":5000", nil)
}
