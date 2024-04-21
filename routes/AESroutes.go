package routes

import (
	"net/http"

	"github.com/congmanh18/NMATTT_AESRSA/handler"
)

// SetupAESRoutes sets up the AES encryption and decryption routes.
func AESRoutes() {
	http.HandleFunc("/AES/encryption", handler.EncryptionAESHandler())
	http.HandleFunc("/AES/decryption", handler.DecryptionAESHandler())
}
