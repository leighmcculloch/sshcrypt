package sshcryptactions

import (
	"golang.org/x/crypto/nacl/secretbox"
)

func DecryptWithPassword(password []byte, cipherText []byte) (clearText []byte, ok bool, err error) {
	var salt [salt_size]byte
	copy(salt[:], cipherText[:16])

	key, err := MakeKey(password, salt[:])
	if err != nil {
		return nil, false, err
	}

	clearText, ok = Decrypt(&key, cipherText[16:])
	return clearText, ok, nil
}

func Decrypt(key *[key_size]byte, cipherText []byte) (clearText []byte, ok bool) {
	var nonce [nonce_size]byte
	copy(nonce[:], cipherText[:nonce_size])
	clearText, ok = secretbox.Open(clearText[:0], cipherText[nonce_size:], &nonce, key)
	return clearText, ok
}
