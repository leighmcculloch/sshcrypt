package sshcryptagent

import (
	"fmt"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func GetSigners() ([]ssh.Signer, error) {
	conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		return nil, fmt.Errorf("Could not connect to the SSH Agent socket. %s", err)
	}

	sshAgent := agent.NewClient(conn)
	signers, err := sshAgent.Signers()
	if err != nil {
		return nil, fmt.Errorf("Could not retrieve signers from the SSH Agent. %v", err)
	}

	if len(signers) == 0 {
		return nil, fmt.Errorf("There are no SSH keys added to the SSH Agent. Check that you have added keys to the SSH Agent and that SSH Agent Forwarding is enabled if you are using this remotely.")
	}

	return signers, nil
}
