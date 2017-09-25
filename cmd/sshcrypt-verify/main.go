package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"4d63.com/sshcrypt/lib/sshcryptdata"

	"golang.org/x/crypto/ssh"
)

var sig = flag.String("s", "", "signature to verify")
var pk = flag.String("k", "", "public-key")

func init() {
	flag.Parse()
}

func main() {
	var sshSignatures []*ssh.Signature
	for _, sig := range strings.Split(*sig, "\n") {
		sshSignature, err := sshcryptdata.DecodeSignature(sig)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sshSignatures = append(sshSignatures, sshSignature)
	}
	if len(sshSignatures) == 0 {
		fmt.Println("Error: At least one signature must be provided.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var sshPublicKeys []ssh.PublicKey
	for _, pk := range strings.Split(*pk, "\n") {
		sshPublicKey, err := sshcryptdata.DecodePublicKey(pk)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sshPublicKeys = append(sshPublicKeys, sshPublicKey)
	}
	if len(sshPublicKeys) == 0 {
		fmt.Println("Error: At least one public key must be provided..")
		flag.PrintDefaults()
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, sig := range sshSignatures {
		for _, pk := range sshPublicKeys {
			err = pk.Verify(data, sig)
			if err == nil {
				fmt.Println("Success")
				os.Exit(0)
			}
		}
	}

	fmt.Println("Failed")
	os.Exit(1)
}
