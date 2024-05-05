package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/congmanh18/NMATTT_AESRSA/model"
	"github.com/congmanh18/NMATTT_AESRSA/utils"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

func EncryptAES(plainText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(plainText))

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, plainText)

	// Prepend IV to ciphertext
	ciphertext = append(iv, ciphertext...)

	return ciphertext, nil
}

func DecryptAES(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext is too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	plaintext := make([]byte, len(ciphertext))

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}

// EncryptionAESHandler handles requests for AES encryption.
func EncryptionAESHandler(db *gorm.DB) http.HandlerFunc {
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
		data.Type = "AES"
		data.Key, _ = utils.GenerateAESKey128Bit()

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Encrypt the message
		encryptedMessage, err := EncryptAES([]byte(*data.Content), []byte(*data.Key))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Encode the encrypted message as base64
		encodedMessage := base64.StdEncoding.EncodeToString(encryptedMessage)
		data.Encrypted_content = encodedMessage
		data.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		db.Create(&data)

		response := struct {
			EncryptedMessage string `json:"encrypted_message"`
			Key              string `json:"key"`
		}{
			EncryptedMessage: encodedMessage,
			Key:              *data.Key,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// DecryptionAESHandler handles requests for AES decryption.
func DecryptionAESHandler() http.HandlerFunc {
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

		// Decode the encrypted message from base64
		encryptedMessage, err := base64.StdEncoding.DecodeString(*data.Content)
		if err != nil {
			http.Error(w, "Invalid base64 encoded message", http.StatusBadRequest)
			return
		}

		// Decrypt the message
		decryptedMessage, err := DecryptAES(encryptedMessage, []byte(*data.Key))
		data.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
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
