// +build windows

package sshcryptagent


import (
	"os/exec"
)


//

func MkCommand(cmdline string) *exec.Cmd {
	return exec.Command("cmd.exe", "/c", cmdline);
}