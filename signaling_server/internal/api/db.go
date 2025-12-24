package Api

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

var Db *sql.DB

type User struct {
	Email    string
	Password string
}

type UserSignup struct {
	Email    string
	Username string
	Password string
}

func InitDB() {

	fmt.Println("Test 1")
	db, err := sql.Open("sqlite", "./auth.db")
	if err != nil {
		fmt.Println("DB Initialization failed!")
		os.Exit(1)
	}

	fmt.Println("Test 1")
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            email TEXT UNIQUE NOT NULL,
            username TEXT NOT NULL,
            password_hash TEXT NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );
    `)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("DB initialized")
	Db = db
}

func GetPassword(email string) (string, error) {
	query := "SELECT password_hash FROM users WHERE email = ?"

	var hashedPassword string

	err := Db.QueryRow(query, email).Scan(&hashedPassword)

	if err != nil {
		return "", fmt.Errorf("No user found with email %s", email)
	}

	return hashedPassword, nil
}

func UnusedEmail(email string) bool {
	query := "SELECT email FROM users WHERE email = ?"
	var returnedEmail string
	err := Db.QueryRow(query, email).Scan(&returnedEmail)

	// no user found
	if err != nil {
		return true
	}
	return false
}

func StoreUser(user UserSignup) (int64, error) {
	query := "INSERT into users (email, username, password_hash) VALUES ($1, $2, $3)"
	result, err := Db.Exec(query, user.Email, user.Username, user.Password)

	if err != nil {
		return 0, fmt.Errorf("Add User: %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("Add User: %v", err)
	}

	return id, nil
}

func FetchUser(email string) (UserSignup, error) {
	query := "SELECT (email, username, password_hash) FROM users WHERE email = ?"
	var user UserSignup

	err := Db.QueryRow(query, email).Scan(&user.Email, &user.Username, &user.Password)

	if err != nil {
		return UserSignup{}, err
	}

	return user, nil

}
