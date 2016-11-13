package sshcryptactions

import (
	"crypto/rand"

	"golang.org/x/crypto/ssh"
)

func Sign(signer ssh.Signer, data []byte) (*ssh.Signature, error) {
	sig, err := signer.Sign(rand.Reader, data)
	if err != nil {
		return nil, err
	}
	return sig, nil
}
