package main

import (
	Api "example/backend/api"
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("testing backend!")
	go Api.WebsockClient()

	go func() {
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			fmt.Println("Background message")
		}
	}()

	http.HandleFunc("/credentials", Api.CredentialsHandler)

	http.ListenAndServe(":5000", nil)

}
