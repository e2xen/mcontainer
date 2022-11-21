package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"os/exec"
)

func main() {
	err := unix.Unshare(unix.CLONE_NEWPID |
		unix.CLONE_NEWUTS | unix.CLONE_NEWNET | unix.CLONE_NEWNS)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	err = unix.Sethostname([]byte("container"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	cmd := exec.Command("/bin/bash")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	cmd.Run()
}
