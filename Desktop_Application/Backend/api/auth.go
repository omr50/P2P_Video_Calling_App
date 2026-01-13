package Api

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdh"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/websocket"
)

var UserJWT string

type Client struct {
	email    string
	username string
	conn     *websocket.Conn
	crypto   *CryptoSession
	isInCall bool
}

type CryptoSession struct {
	privateKey *ecdh.PrivateKey
	publicKey  *ecdh.PublicKey
	sharedKey  []byte
	aesGCM     cipher.AEAD
}

type User struct {
	Email    string
	Password string
}

type TokenResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func generateKeyPair() (*ecdh.PrivateKey, *ecdh.PublicKey, error) {
	curve := ecdh.P256()
	privateKey, err := curve.GenerateKey(rand.Reader)

	if err != nil {
		return nil, nil, err
	}

	return privateKey, privateKey.PublicKey(), nil
}

func computeSharedKey(privateKey *ecdh.PrivateKey, peerPublicKeyBytes []byte) ([]byte, error) {
	curve := ecdh.P256()
	peerPublicKey, err := curve.NewPublicKey(peerPublicKeyBytes)

	if err != nil {
		return nil, err
	}

	sharedSecret, err := privateKey.ECDH(peerPublicKey)

	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(sharedSecret)
	return hash[:], nil
}

func initAESGCM(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)

	return aesGCM, nil
}

func (cs *CryptoSession) Encrypt(plaintext []byte) (ciphertext, nonce []byte, err error) {
	if cs.aesGCM == nil {
		return nil, nil, fmt.Errorf("encryption not initialized")
	}

	nonce = make([]byte, cs.aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

	ciphertext = cs.aesGCM.Seal(nil, nonce, plaintext, nil)

	return ciphertext, nonce, nil
}

func (cs *CryptoSession) Decrypt(ciphertext, nonce []byte) ([]byte, error) {
	if cs.aesGCM == nil {
		return nil, fmt.Errorf("decryption not initialized")
	}

	plaintext, err := cs.aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func CredentialsHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(u)

	requestBody := bytes.NewBuffer(jsonData)

	// send this user to the signaling server to authenticate
	resp, err := http.Post("http://localhost:8090/login", "application/json", requestBody)

	if err != nil {
		http.Error(w, "signaling server unavailable", http.StatusBadGateway)
		return
	}

	var token TokenResponse

	err = json.NewDecoder(resp.Body).Decode(&token)

	json.NewEncoder(w).Encode(map[string]string{
		"token":    token.Token,
		"username": token.Username,
		"email":    token.Email,
	})

	fmt.Println("Token: ", token)
	UserJWT = token.Token
	GlobalUserEmail = token.Email

	go WebsockClient()
}
