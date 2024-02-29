package helper

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudfront/sign"
)

// URLSigner is a struct that holds the key ID, private key, raw URL, and hour.
type URLSigner struct {
	KeyID         string        // Credential Key Pair key ID
	PrivKeyBase64 string        // private key in base64
	Hour          time.Duration // hour duration
}

// Signer signs a URL to be valid, using the private key and credential pair key ID.
func (us *URLSigner) Signer(url string) (string, error) {
	// Load the private key, and return an error if it fails
	privKeyBytes, err := loadPrivateKey(us.PrivKeyBase64)
	if err != nil {
		return "", err
	}

	// Create a new URLSigner with the key ID and private key
	signer := sign.NewURLSigner(us.KeyID, privKeyBytes)
	signedURL, err := signer.Sign(url, time.Now().Add(us.Hour))
	return signedURL, nil
}

// loadPrivateKey loads a RSA private key from the file path.
func loadPrivateKey(privKeyBase64 string) (*rsa.PrivateKey, error) {
	// Decode the base64 string to bytes
	privKeyBytes, err := base64.StdEncoding.DecodeString(privKeyBase64)
	if err != nil {
		return nil, err
	}

	// Decode the PEM block to get the private key
	block, _ := pem.Decode(privKeyBytes)

	// Parse the private key
	// #TODO: handle fallback to PKCS1 if PKCS8 fails to parse
	privKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privKey.(*rsa.PrivateKey), nil
}
