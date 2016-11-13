package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/leighmcculloch/sshcrypt/lib/sshcryptactions"
	"github.com/leighmcculloch/sshcrypt/lib/sshcryptdata"
)

var password string

func init() {
	var passwordPtr = flag.String("p", "", "Password to use for decryption")
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

	cipherText, err := sshcryptdata.DecodeCipherText(string(data))
	if err != nil {
		fail(err)
	}

	clearText, ok, err := sshcryptactions.DecryptWithPassword([]byte(password), cipherText)
	if err != nil {
		fail(err)
	}
	if !ok {
		fail(fmt.Errorf("Decryption not possible"))
	}
	os.Stdout.Write(clearText)
}
