package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/leighmcculloch/sshcrypt/lib/sshcryptactions"
	"github.com/leighmcculloch/sshcrypt/lib/sshcryptagent"
	"github.com/leighmcculloch/sshcrypt/lib/sshcryptdata"
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

	var sigs []string
	for _, signer := range signers {
		sig, err := sshcryptactions.Sign(signer, data)
		if err != nil {
			fail(err)
		}
		sigs = append(sigs, sshcryptdata.EncodeSignature(sig))
	}

	fmt.Println(strings.Join(sigs, "\n"))
}
