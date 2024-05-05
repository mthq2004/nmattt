package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
)

// GenerateRSAKeyPairString generates RSA public and private keys and returns pointers to their string representations
func GenerateRSAKeyPairString(bits int) (*string, *string, error) {
	// Generate primes p and q
	p, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		fmt.Println("Failed to generate prime p:", err)
		return nil, nil, err
	}

	q, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		fmt.Println("Failed to generate prime q:", err)
		return nil, nil, err
	}

	// Calculate modulus N = p * q
	N := new(big.Int).Mul(p, q)

	// Calculate Euler's totient function φ(N) = (p-1)(q-1)
	phi := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))

	// Choose public exponent e
	e := 65537

	// Find private exponent d
	d := new(big.Int).ModInverse(big.NewInt(int64(e)), phi)

	// Public key
	publicKey := &rsa.PublicKey{
		N: N,
		E: e,
	}

	// Marshal public key to PEM format
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, err
	}
	pubKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	}
	pubKeyString := string(pem.EncodeToMemory(pubKeyBlock))

	// Private key
	privateKey := &rsa.PrivateKey{
		PublicKey: *publicKey,
		D:         d,
	}

	// Marshal private key to PEM format
	privKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privKeyBytes,
	}
	privKeyString := string(pem.EncodeToMemory(privKeyBlock))

	return &pubKeyString, &privKeyString, nil
}
