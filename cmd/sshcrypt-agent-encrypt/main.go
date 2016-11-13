package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/leighmcculloch/sshcrypt/lib/sshcryptactions"
	"github.com/leighmcculloch/sshcrypt/lib/sshcryptagent"
	"github.com/leighmcculloch/sshcrypt/lib/sshcryptdata"
)

const challengeSize = 64

func fail(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	signers, err := sshcryptagent.GetSigners()
	if err != nil {
		fail(err)
	}
	if len(signers) == 0 {
		fail(fmt.Errorf("Error: At least one signer must be provided. Check that your SSH Agent has at least one key added."))
		return
	}

	data, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		fail(err)
	}

	var cipherTexts []string
	for _, signer := range signers {
		var challenge [challengeSize]byte
		if _, err := rand.Read(challenge[:]); err != nil {
			fail(err)
		}

		sig, err := sshcryptactions.Sign(signer, challenge[:])
		if err != nil {
			fail(err)
		}

		cipherText, err := sshcryptactions.EncryptWithPassword(sig.Blob, data)
		if err != nil {
			fail(err)
		}

		encodedCipherText := sshcryptdata.EncodeChallengeCipherText(challenge[:], cipherText)
		cipherTexts = append(cipherTexts, encodedCipherText)
	}

	fmt.Println(strings.Join(cipherTexts, "\n"))
}
