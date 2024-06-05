// Package jwt provides structures and functions for handling JSON Web Tokens (JWT).
package jwt

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

// JwtEncode takes multiple byte slices, encodes each to base64, and joins them with dots.
// This function is typically used to encode the header, payload, and signature parts of a JWT.
func JwtEncode(parts ...[]byte) string {
	var result []string
	// Iterate over each part, encode to base64, and append to result slice.
	for _, p := range parts {
		result = append(result, base64.StdEncoding.EncodeToString(p))
	}
	// Join the encoded parts with dots and return the result.
	return strings.Join(result, ".")
}

// JwtMarshal marshals a given data structure into a JSON byte slice.
// It panics if the marshaling process encounters an error.
func JwtMarshal(data interface{}) []byte {
	partJSON, err := json.Marshal(data)
	// Check for errors during marshaling and panic if any occur.
	if err != nil {
		panic(err)
	}
	// Return the marshaled JSON byte slice.
	return partJSON
}

