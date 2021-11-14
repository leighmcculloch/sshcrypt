// +build !windows,!plan9

package sshcryptagent


import (
	"os/exec"
)


//

func MkCommand(cmdline string) *exec.Cmd {
	return exec.Command("sh", "-c", cmdline);
}