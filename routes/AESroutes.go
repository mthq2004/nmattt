package routes

import (
	"net/http"

	"github.com/congmanh18/NMATTT_AESRSA/handler"
	"gorm.io/gorm"
)

func AESRoutes(db *gorm.DB) {
	http.HandleFunc("/AES/encryption", handler.EncryptionAESHandler(db))
	http.HandleFunc("/AES/decryption", handler.DecryptionAESHandler())
}
