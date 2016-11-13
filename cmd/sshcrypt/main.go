package main

import (
	"fmt"
	"os"
	"os/exec"
)

func printUsage() {
	fmt.Printf(
		`Usage: sshcrypt [command]
Commands:
	agent-encrypt - Encrypt data using the signature of a random challenge as the key
	agent-decrypt - Decrypt data using the signature of a random challenge as the key
	encrypt       - Encrypt data using a password
	decrypt       - Decrypt data using a password
	agent-sign    - Sign data with the SSH Agent
	verify        - Verify SSH signatures
`)
}

func fail(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	if len(os.Args) == 1 {
		printUsage()
		return
	}

	command := os.Args[1]
	switch command {
	case "agent-encrypt", "agent-decrypt", "encrypt", "decrypt", "agent-sign", "verify":
		cmd := exec.Command(fmt.Sprintf("sshcrypt-%s", command), os.Args[2:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fail(err)
		}
	default:
		printUsage()
	}
}
