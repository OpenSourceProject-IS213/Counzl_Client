package cmd

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"strings"
	"utilities/converter"
)

func Invisible_in(description string) string {
	fmt.Print(description)

	byteInput, err := terminal.ReadPassword(0)
	if err != nil {
		fmt.Print("\n")
		log.Fatal(err)
	}
	passord := string(byteInput)
	return converter.Remove_0a(strings.TrimSpace(passord))
}
