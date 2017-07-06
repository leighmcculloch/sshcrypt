package sshcryptactions

import (
	"golang.org/x/crypto/nacl/secretbox"
)

func DecryptWithPassword(password []byte, cipherText []byte) (clearText []byte, ok bool, err error) {
	var salt [saltSize]byte
	copy(salt[:], cipherText[:16])

	key, err := MakeKey(password, salt[:])
	if err != nil {
		return nil, false, err
	}

	clearText, ok = Decrypt(&key, cipherText[16:])
	return clearText, ok, nil
}

func Decrypt(key *[keySize]byte, cipherText []byte) (clearText []byte, ok bool) {
	var nonce [nonceSize]byte
	copy(nonce[:], cipherText[:nonceSize])
	clearText, ok = secretbox.Open(clearText[:0], cipherText[nonceSize:], &nonce, key)
	return clearText, ok
}
