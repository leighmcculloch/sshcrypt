package sshcryptdata

import (
	"encoding/base64"
	"strings"

	"golang.org/x/crypto/ssh"
)

func DecodePublicKey(pk string) (ssh.PublicKey, error) {
	parts := strings.SplitN(pk, " ", 3)
	if len(parts) < 2 {
		return nil, nil
	}
	blob, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	return ssh.ParsePublicKey(blob)
}
