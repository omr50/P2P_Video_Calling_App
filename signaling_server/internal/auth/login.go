package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	Api "github.com/omr50/P2P_Video_Calling_App/internal/api"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("secret-key")

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func validPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true

}

func fetchAndValidatePassword(email string, password string) bool {
	storedPaswordHashed, err := Api.GetPassword(email)

	if err != nil {
		return false
	}

	return validPassword(storedPaswordHashed, password)

}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup endpoint hit")
	w.Header().Set("Content-Type", "application/json")

	var user Api.UserSignup
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, err := hashPassword(user.Password)
	user.Password = hashedPassword

	if err != nil {
		// 400 bad request
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	unusedEmail := Api.UnusedEmail(user.Email)

	if !unusedEmail {
		// 409 conflict
		w.WriteHeader(http.StatusConflict)
		return
	}

	_, err = Api.StoreUser(user)

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

	var u Api.User
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Println("jsondata: ", u)

	// user, err := Api.FetchUser(u.Email)

	if fetchAndValidatePassword(u.Email, u.Password) {
		tokenString, err := createToken(u.Email)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Errorf("No username found")
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"token": tokenString,
		})
		fmt.Println("successfully sending token")
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid Credentials")
	}
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}
	tokenString = tokenString[len("Bearer "):]

	fmt.Println("token:", tokenString)

	err := verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}

	fmt.Fprint(w, "Welcome to the protected aread!")
}
