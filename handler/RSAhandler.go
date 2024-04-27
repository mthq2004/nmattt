package handler

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"net/http"
	"time"

	"github.com/congmanh18/NMATTT_AESRSA/model"
	"github.com/congmanh18/NMATTT_AESRSA/utils"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

// EncryptRSA encrypts a message using RSA encryption.
func EncryptRSA(plainText []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		return nil, err
	}
	return encrypted, nil
}

// DecryptRSA decrypts a message using RSA decryption.
func DecryptRSA(encrypted []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encrypted)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}

// EncryptionRSAHandler handles requests for RSA encryption.
func EncryptionRSAHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
			// Allow all origins
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Allow only POST and OPTIONS
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		// Allow only Content-Type header
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// Parse request body
		var data model.Data
		data.ID = uuid.New().String()
		data.Type = "RSA"
		data.PublicKey, data.PrivateKey, _ = utils.GenerateRSAKeyPair(2048)

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Parse public key from PEM format
		block, _ := pem.Decode([]byte(*data.PublicKey))
		if block == nil || block.Type != "PUBLIC KEY" {
			http.Error(w, "Failed to decode PEM block containing public key", http.StatusBadRequest)
			return
		}
		pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		publicKey, ok := pubInterface.(*rsa.PublicKey)
		if !ok {
			http.Error(w, "Failed to parse public key", http.StatusBadRequest)
			return
		}

		// Encrypt the message
		encryptedMessage, err := EncryptRSA([]byte(*data.Content), publicKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Encode the encrypted message as base64
		encodedMessage := base64.StdEncoding.EncodeToString(encryptedMessage)
		data.Encrypted_content = encodedMessage
		data.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		db.Create(&data)
		// Send the encrypted message in the response
		response := struct {
			EncryptedMessage string `json:"encrypted_message"`
			PublicKey        string `json:"publicKey"`
			PrivateKey       string `json:"privateKey"`
		}{
			EncryptedMessage: encodedMessage,
			PublicKey:        *data.PublicKey,
			PrivateKey:       *data.PrivateKey,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// DecryptionRSAHandler handles requests for RSA decryption.
func DecryptionRSAHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Allow all origins
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Allow only POST and OPTIONS
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		// Allow only Content-Type header
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// Parse request body
		var data model.Data
		data.ID = uuid.New().String()
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Parse private key from PEM format
		block, _ := pem.Decode([]byte(*data.PrivateKey))
		if block == nil || block.Type != "RSA PRIVATE KEY" {
			http.Error(w, "Failed to decode PEM block containing private key", http.StatusBadRequest)
			return
		}
		privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Decode the encrypted message from base64
		encryptedMessage, err := base64.StdEncoding.DecodeString(*data.Content)
		data.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		if err != nil {
			http.Error(w, "Invalid base64 encoded message", http.StatusBadRequest)
			return
		}

		// Decrypt the message
		decryptedMessage, err := DecryptRSA(encryptedMessage, privateKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send the decrypted message in the response
		response := struct {
			DecryptedMessage string `json:"decrypted_message"`
		}{
			DecryptedMessage: string(decryptedMessage),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
