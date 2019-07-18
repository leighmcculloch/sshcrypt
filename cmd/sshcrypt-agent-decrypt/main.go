package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"4d63.com/sshcrypt/lib/sshcryptactions"
	"4d63.com/sshcrypt/lib/sshcryptagent"
	"4d63.com/sshcrypt/lib/sshcryptdata"
)

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
	}

	data, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		fail(err)
	}

	cipherTextPackages := strings.Split(string(data), "\n")
	if len(cipherTextPackages) == 0 {
		fail(fmt.Errorf("Error: At least one piece of encrypted data must be provided."))
	}
	for _, cipherTextPackage := range cipherTextPackages {
		if len(cipherTextPackage) == 0 {
			continue
		}

		challenge, cipherText, err := sshcryptdata.DecodeChallengeCipherText(cipherTextPackage)
		if err != nil {
			fail(err)
		}

		for _, signer := range signers {
			sig, err := sshcryptactions.Sign(signer, challenge)
			if err != nil {
				fail(err)
			}

			clearText, ok, err := sshcryptactions.DecryptWithPassword(sig.Blob, cipherText)
			if err != nil {
				fail(err)
			}
			if ok {
				os.Stdout.Write(clearText)
				os.Exit(0)
			}
		}
	}

	fmt.Println("Decryption not possible")
	os.Exit(1)
}
