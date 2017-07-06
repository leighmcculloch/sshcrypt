package sshcryptactions

import "golang.org/x/crypto/scrypt"

func MakeKey(password, salt []byte) (key [keySize]byte, err error) {
	keyBytes, err := scrypt.Key(password, salt, 16384, 8, 1, 32)
	copy(key[:], keyBytes)
	return
}
