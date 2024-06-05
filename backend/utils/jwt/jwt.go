// Package jwt provides structures and functions for handling JSON Web Tokens (JWT).
package jwt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"os"
	"strings"
)

// GenerateKey generates a new RSA key pair and assigns it to the Key struct.
func (k *Key) GenerateKey() error {
	var err error
	k.Private, err = rsa.GenerateKey(rand.Reader, 2048) // Generate a new 2048-bit RSA private key.
	if err != nil {
		return err
	}

	k.Public = &k.Private.PublicKey // Assign the corresponding public key.
	return nil
}

// PEMfromKey converts the RSA key pair in the Key struct to PEM encoded format and returns a PEMKey struct.
func (k *Key) PEMfromKey() *PEMKey {
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",                  // PEM block type for private key.
		Bytes: x509.MarshalPKCS1PrivateKey(k.Private), // Encode the private key to PKCS#1 format.
	}

	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",                  // PEM block type for public key.
		Bytes: x509.MarshalPKCS1PublicKey(k.Public), // Encode the public key to PKCS#1 format.
	}

	return &PEMKey{
		Private: privateKeyPEM,
		Public:  publicKeyPEM,
	}
}

// SetPEMToFile writes the PEM encoded private and public keys to files.
func (pm *PEMKey) SetPEMToFile(path string) error {
	// Create and write the private key to a file.
	privateKeyFile, err := os.Create(path + "/private_key.pem")
	if err != nil {
		return err
	}
	pem.Encode(privateKeyFile, pm.Private)

	// Create and write the public key to a file.
	publicKeyFile, err := os.Create(path + "/public_key.pem")
	if err != nil {
		return err
	}
	pem.Encode(publicKeyFile, pm.Public)
	return nil
}

// KeyfromPrivateFile reads a PEM encoded private key from a file and assigns it to the Key struct.
func (k *Key) KeyfromPrivateFile(filePath string) error {
	privateKeyFile, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	block, _ := pem.Decode(privateKeyFile)
	if block == nil {
		return errors.New("failed to decode PEM block containing private key")
	}
	k.Private, err = x509.ParsePKCS1PrivateKey(block.Bytes) // Parse the PEM encoded private key.
	if err != nil {
		return err
	}
	return nil
}

// KeyfromPublicFile reads a PEM encoded public key from a file and assigns it to the Key struct.
func (k *Key) KeyfromPublicFile(filePath string) (err error) {
	publicKeyFile, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	block, _ := pem.Decode(publicKeyFile)
	if block == nil {
		return errors.New("failed to decode PEM block containing public key")
	}
	k.Public, err = x509.ParsePKCS1PublicKey(block.Bytes) // Parse the PEM encoded public key.
	if err != nil {
		return err
	}
	return nil
}

// GenerateToken generates a JWT with the specified user ID, signs it with the provided private key, and returns the token as a string.
func (jwt *JWT) GenerateToken(userId int, privateKey *rsa.PrivateKey) string {
	jwt.Header = Header{
		Alg: "RS256", // Signing algorithm.
		Typ: "JWT",   // Token type.
	}

	jwt.Payload = Payload{
		Id:  userId,             // User ID.
		Sub: "authentification", // Subject of the token.
	}

	// Marshal the header and payload to JSON.
	header_to_json := JwtMarshal(jwt.Header)
	payload_to_json := JwtMarshal(jwt.Payload)
	// Encode the header and payload to base64 and concatenate them with a dot.
	msg := JwtEncode(header_to_json, payload_to_json)

	// Create a SHA-256 hash of the message.
	msgHash := sha256.New()
	_, err := msgHash.Write([]byte(msg))
	if err != nil {
		panic(err)
	}

	msgHashSum := msgHash.Sum(nil)

	// Sign the hash using RSA PSS and the provided private key.
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		panic(err)
	}

	// Encode the signature to base64 and concatenate it to the message.
	token := msg + "." + JwtEncode(signature)
	return token
}

// VerifyToken verifies the signature of a JWT using the provided public key and returns the payload if valid.
func (jwt *JWT) VerifyToken(token string, publicKey *rsa.PublicKey) (*Payload, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token format") // Token must have three parts.
	}

	// Concatenate the header and payload parts.
	message := parts[0] + "." + parts[1]
	// Decode the signature from base64.
	signature, err := base64.StdEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, err
	}

	// Create a SHA-256 hash of the message.
	msgHash := sha256.New()
	_, err = msgHash.Write([]byte(message))
	if err != nil {
		return nil, err
	}
	msgHashSum := msgHash.Sum(nil)

	// Verify the signature using RSA PSS and the provided public key.
	err = rsa.VerifyPSS(publicKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		return nil, errors.New("invalid token signature") // Signature verification failed.
	}

	var payload Payload
	// Decode the payload from base64.
	payloadJSON, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON payload.
	err = json.Unmarshal(payloadJSON, &payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
