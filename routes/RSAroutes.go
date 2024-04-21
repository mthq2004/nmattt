package routes

import (
	"net/http"

	"github.com/congmanh18/NMATTT_AESRSA/handler"
)

// SetupAESRoutes sets up the AES encryption and decryption routes.
func RSARoutes() {
	http.HandleFunc("/RSA/encryption", handler.EncryptionRSAHandler())
	http.HandleFunc("/RSA/decryption", handler.DecryptionRSAHandler())
}
