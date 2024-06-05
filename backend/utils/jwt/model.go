// Package jwt provides structures and functions for handling JSON Web Tokens (JWT).
package jwt

import (
	"crypto/rsa"
	"encoding/pem"
)

// JWT represents a JSON Web Token with a header and payload.
type JWT struct {
	Header  Header // Header contains the algorithm and token type information.
	Payload Payload // Payload contains the claims of the token.
}

// Header represents the header part of a JWT, which typically includes
// the signing algorithm and the type of the token.
type Header struct {
	Alg string `json:"alg"` // Alg specifies the algorithm used to sign the token.
	Typ string `json:"typ"` // Typ specifies the type of the token, typically "JWT".
}

// Payload represents the payload part of a JWT, which typically contains
// the claims or assertions about the token's subject.
type Payload struct {
	Id  int    `json:"id"`  // Id is a unique identifier for the token's subject.
	Sub string `json:"sub"` // Sub is the subject of the token, typically a user ID or email.
}

// Key represents an RSA key pair used for signing and verifying JWTs.
type Key struct {
	Private *rsa.PrivateKey // Private is the RSA private key used for signing tokens.
	Public  *rsa.PublicKey  // Public is the RSA public key used for verifying tokens.
}

// PEMKey represents PEM encoded RSA keys.
// It contains PEM blocks for both the private and public keys.
type PEMKey struct {
	Private *pem.Block // PEM encoded private key block.
	Public  *pem.Block // PEM encoded public key block.
}

