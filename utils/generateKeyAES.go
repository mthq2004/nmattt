package utils

import (
    "crypto/rand"
    "encoding/hex"
    "fmt"
)

// GenerateAESKey128Bit generates a random 128-bit key for AES encryption.
func GenerateAESKey128Bit() (*string, error) {
    key := make([]byte, 16) // 128-bit key
    _, err := rand.Read(key)
    if err != nil {
        return nil, fmt.Errorf("error generating AES key: %v", err)
    }
    keyString := hex.EncodeToString(key)
    return &keyString, nil
}
