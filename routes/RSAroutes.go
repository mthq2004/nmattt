package routes

import (
	"net/http"

	"github.com/congmanh18/NMATTT_AESRSA/handler"
	"gorm.io/gorm"
)

// SetupAESRoutes sets up the AES encryption and decryption routes.
func RSARoutes(db *gorm.DB) {
	http.HandleFunc("/RSA/encryption", handler.EncryptionRSAHandler(db))
	http.HandleFunc("/RSA/decryption", handler.DecryptionRSAHandler())
}
