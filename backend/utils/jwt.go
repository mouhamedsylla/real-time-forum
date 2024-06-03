package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub string `json:"sub"`
	Iat int64  `json:"iat"`
	Exp int64  `json:"exp"`
}

type JWT struct {
	Hd  Header
	Pld Payload
	Key   Key
}

type Key struct {
	Private *rsa.PrivateKey
	Public  *rsa.PublicKey
}

func (k *Key) GenerateKey() error {
	var err error
	k.Private, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	k.Public = &k.Private.PublicKey
	return nil
}

func (jwt *JWT) GenerateToken() string {
	jwt.Hd = Header{
		Alg: "RS256",
		Typ: "JWT",
	}

	jwt.Pld = Payload{
		Sub: "1234567890",
		Iat: time.Now().Unix(),
		Exp: time.Now().Add(time.Hour * 1).Unix(), // Expire dans une heure
	}

	headerJSON := MarshalPart(jwt.Hd)
	payloadJSON := MarshalPart(jwt.Pld)
	message := PartEncode(headerJSON, payloadJSON)

	msgHash := sha256.New()
	_, err := msgHash.Write([]byte(message))
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)

	err = jwt.Key.GenerateKey()
	if err != nil {
		panic(err)
	}

	signature, err := rsa.SignPSS(rand.Reader, jwt.Key.Private, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		panic(err)
	}

	Token := message + "." + PartEncode(signature)
	return Token
}

func VerifyToken(token string, pubKey *rsa.PublicKey) (*Payload, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token format")
	}

	message := parts[0] + "." + parts[1]
	signature, err := base64.StdEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, err
	}

	msgHash := sha256.New()
	_, err = msgHash.Write([]byte(message))
	if err != nil {
		return nil, err
	}
	msgHashSum := msgHash.Sum(nil)

	err = rsa.VerifyPSS(pubKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		return nil, errors.New("invalid token signature")
	}

	var payload Payload
	payloadJSON, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(payloadJSON, &payload)
	if err != nil {
		return nil, err
	}

	if payload.Exp < time.Now().Unix() {
		return nil, errors.New("token has expired")
	}

	return &payload, nil
}

func PartEncode(parts ...[]byte) string {
	var result []string
	for _, p := range parts {
		result = append(result, base64.StdEncoding.EncodeToString(p))
	}
	return strings.Join(result, ".")
}

func MarshalPart(data interface{}) []byte {
	partJSON, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return partJSON
}
