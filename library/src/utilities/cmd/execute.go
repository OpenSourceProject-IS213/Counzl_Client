package cmd

import (
	"log"
	"os/exec"
)

func Exe_cmd(cmd string) []byte {
	out, err := exec.Command(cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}
