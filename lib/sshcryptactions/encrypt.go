package sshcryptactions

import (
	"crypto/rand"

	"golang.org/x/crypto/nacl/secretbox"
)

func EncryptWithPassword(password []byte, clearText []byte) (cipherText []byte, err error) {
	var salt [saltSize]byte
	if _, err := rand.Read(salt[:]); err != nil {
		return nil, err
	}

	key, err := MakeKey(password, salt[:])
	if err != nil {
		return nil, err
	}

	cipherText, err = Encrypt(&key, clearText)
	if err != nil {
		return nil, err
	}

	return append(salt[:], cipherText...), nil
}

func Encrypt(key *[keySize]byte, clearText []byte) (cipherText []byte, err error) {
	var nonce [nonceSize]byte
	if _, err := rand.Read(nonce[:]); err != nil {
		return nil, err
	}
	return secretbox.Seal(nonce[:], clearText, &nonce, key), nil
}
