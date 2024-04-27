package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// GenerateRSAKeyPair tạo một cặp khóa RSA với kích thước đã cho và trả về public key và private key dưới dạng string.
func GenerateRSAKeyPair(bits int) (*string, *string, error) {
	// Tạo cặp khóa RSA
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, fmt.Errorf("lỗi khi tạo private key: %v", err)
	}

	// Tạo public key từ private key
	publicKey := &privateKey.PublicKey

	// Convert private key thành PEM format
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	privateKeyStr := string(privateKeyPEM)

	// Convert public key thành DER format
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, fmt.Errorf("lỗi khi chuyển public key sang DER format: %v", err)
	}
	// Encode public key thành PEM format
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	publicKeyStr := string(publicKeyPEM)

	return &publicKeyStr, &privateKeyStr, nil
}
