package sshcryptdata

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

func EncodeSignature(sig *ssh.Signature) string {
	return fmt.Sprintf("%s %s", sig.Format, base64.StdEncoding.EncodeToString(sig.Blob))
}

func DecodeSignature(sig string) (*ssh.Signature, error) {
	parts := strings.SplitN(sig, " ", 3)
	if len(parts) < 2 {
		return nil, errors.New("Invalid signature format")
	}
	format := parts[0]
	blob, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	return &ssh.Signature{
		Format: format,
		Blob:   blob,
	}, nil
}
