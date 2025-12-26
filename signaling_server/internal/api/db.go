package Api

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
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

	connStr := "postgres://appuser:apppassword@localhost:5432/appdb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("DB Initialization failed!")
		os.Exit(1)
	}

	fmt.Println("Connected to PostgreSQL")

	fmt.Println("Test 2")
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            email VARCHAR(255) UNIQUE NOT NULL,
            username VARCHAR(255) NOT NULL,
            password_hash TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `)

	fmt.Println("Test 3")
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("DB initialized")
	Db = db
}

func GetPassword(email string) (string, error) {
	query := "SELECT password_hash FROM users WHERE email = $1"

	var hashedPassword string

	err := Db.QueryRow(query, email).Scan(&hashedPassword)

	if err != nil {
		return "", fmt.Errorf("No user found with email %s", email)
	}

	return hashedPassword, nil
}

func UnusedEmail(email string) bool {
	query := "SELECT email FROM users WHERE email = $1"
	var returnedEmail string
	err := Db.QueryRow(query, email).Scan(&returnedEmail)

	// no user found
	if err != nil {
		return true
	}
	return false
}

func StoreUser(user UserSignup) (int64, error) {
	query := "INSERT INTO users (email, username, password_hash) VALUES ($1, $2, $3) RETURNING id"

	var id int64
	err := Db.QueryRow(query, user.Email, user.Username, user.Password).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("Add User: %v", err)
	}

	return id, nil
}

func FetchUser(email string) (UserSignup, error) {
	query := "SELECT email, username, password_hash FROM users WHERE email = $1"
	var user UserSignup

	err := Db.QueryRow(query, email).Scan(&user.Email, &user.Username, &user.Password)

	if err != nil {
		return UserSignup{}, err
	}

	return user, nil

}
