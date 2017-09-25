package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"4d63.com/sshcrypt/lib/sshcryptactions"
	"4d63.com/sshcrypt/lib/sshcryptdata"
)

var password string

func init() {
	var passwordPtr = flag.String("p", "", "Password to use for encryption")
	flag.Parse()
	password = *passwordPtr
}

func fail(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	data, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		fail(err)
	}

	cipherText, err := sshcryptactions.EncryptWithPassword([]byte(password), data)
	if err != nil {
		fail(err)
	}

	encodedCipherText := sshcryptdata.EncodeCipherText(cipherText)

	fmt.Println(encodedCipherText)
}
