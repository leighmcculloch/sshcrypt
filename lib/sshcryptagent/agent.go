package sshcryptagent

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type ReadWriter struct {
	io.Reader
	io.Writer
}

func NewReadWriter(r io.Reader, w io.Writer) io.ReadWriter {
	return &ReadWriter{r, w}
}

func GetSigners() ([]ssh.Signer, error) {
	sockAddr := os.Getenv("SSH_AUTH_SOCK")

	// ex: "socat.exe - UNIX-CONNECT:%a"
	socatFormat := os.Getenv("SSH_AUTH_SOCAT")
	var conn io.ReadWriter
	var err error
	if len(socatFormat) > 0 {
		socatCommand := fmtCommand(socatFormat, sockAddr)
		cmd := MkCommand(socatCommand)
		cmd.Stderr = os.Stderr
		pipeRd, err := cmd.StdoutPipe()
		pipeWr, err := cmd.StdinPipe()

		err = cmd.Start()
		if err != nil {
			return nil, fmt.Errorf("Could not start \"%s\": %s", socatCommand, err)
		}
		conn = NewReadWriter(pipeRd, pipeWr)
	} else {
		conn, err = net.Dial("unix", sockAddr)
		if err != nil {
			return nil, fmt.Errorf("Could not connect to the SSH Agent socket. %s", err)
		}
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

func fmtCommand(format string, address string) string {
	var out strings.Builder
	inPercent := false
	for _, c := range format {
		if inPercent {
			inPercent = false
			if (c == '%') {
				out.WriteRune('%');
			} else if (c == 'a') {
				out.WriteString(address);
			} else {
				out.WriteRune('%')
				out.WriteRune(c);
			}
		} else {
			if c == '%' {
				inPercent = true
			} else {
				out.WriteRune(c);
			}
		}
	}
	if inPercent {
		out.WriteRune('%')
	}
	return out.String()
}
