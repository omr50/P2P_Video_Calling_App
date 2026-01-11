package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/omr50/P2P_Video_Calling_App/internal/auth"
)

func FetchAndValidatePassword(email string, password string) bool {
	storedPaswordHashed, err := GetPassword(email)

	if err != nil {
		return false
	}

	return auth.ValidPassword(storedPaswordHashed, password)

}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup endpoint hit")
	w.Header().Set("Content-Type", "application/json")

	var user UserSignup
	json.NewDecoder(r.Body).Decode(&user)
	hashedPassword, err := auth.HashPassword(user.Password)
	user.Password = hashedPassword

	if err != nil {
		// 400 bad request
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	unusedEmail := UnusedEmail(user.Email)

	if !unusedEmail {
		// 409 conflict
		w.WriteHeader(http.StatusConflict)
		return
	}

	_, err = StoreUser(user)

	if err != nil {
		// 400 bad request
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 200 succeeded
	w.WriteHeader(http.StatusOK)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u User
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Println("jsondata: ", u)

	user, err := FetchUser(u.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "No user found!")
	}

	if auth.ValidPassword(user.Password, u.Password) {
		tokenString, err := auth.CreateToken(u.Email)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "No username found!")
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"token":    tokenString,
			"username": user.Username,
		})
		fmt.Println("successfully sending token")
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid Credentials")
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		isAuth := auth.IsAuthenticated(r)

		if !isAuth {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Blank or Invalid Token!")
			return
		}
		fmt.Println("Middleware works! Accessing protected Area")
		next.ServeHTTP(w, r)
	})

}
