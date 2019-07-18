package sshcryptdata

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

func EncodeCipherText(cipherText []byte) string {
	return base64.StdEncoding.EncodeToString(cipherText)
}

func DecodeCipherText(encodedCipherText string) (cipherText []byte, err error) {
	return base64.StdEncoding.DecodeString(encodedCipherText)
}

func EncodeChallengeCipherText(challenge []byte, cipherText []byte) string {
	return fmt.Sprintf(
		"%s %s",
		base64.StdEncoding.EncodeToString(challenge),
		base64.StdEncoding.EncodeToString(cipherText))
}

func DecodeChallengeCipherText(encodedCipherText string) (challenge []byte, cipherText []byte, err error) {
	parts := strings.Split(encodedCipherText, " ")
	if len(parts) != 2 {
		return nil, nil, errors.New("Encrypted data is invalid format")
	}
	challenge, err = base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, nil, err
	}
	cipherText, err = base64.StdEncoding.DecodeString(parts[1])
	return challenge, cipherText, err
}
