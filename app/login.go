package app

import (
	"./modules/users"
	"bufio"
	"fmt"
	"os"
	"strings"
	"utilities/converter"
	"utilities/crypto"
	"utilities/cmd"
)


// Alternativ løsning på skjult passord-input. https://play.golang.org/p/l-9IP1mrhA

func Login() (string, bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	brukernavn, _ := reader.ReadString('\n')
	brukernavn = converter.Remove_0a(strings.TrimSpace(brukernavn))

	passord := cmd.Invisible_in(STDIN_DESCRIPTION)

	return checkPassword(brukernavn, passord)
}

func checkPassword(brukernavn, passord string) (string, bool) {
	hash_db, _, online := users.FetchUser(brukernavn)   // Henter passord fra db
	match := crypto.CheckPasswordHash(passord, hash_db) // Sammenligner passord (plain text) fra stdin og hashet fra passord fra db

	if hash_db != "" {
		if match == true {
			Welcome(converter.StringToBool(online))
		} else {
			fmt.Println("\nPassword or username is wrong...")
			tryAgain()
		}
	} else {
		fmt.Println("\nThe user does not exist")
		tryAgain()
	}
	return brukernavn, match
}

func tryAgain() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to try again (y/n)? ")
	svar, _ := reader.ReadString('\n')
	svar = converter.Remove_0a(svar)
	switch svar {
	case "y":
		Login()
	case "n":
		os.Exit(0)
	default:
		fmt.Println("You have to choose between 'y' or 'n'.")
	}
}
