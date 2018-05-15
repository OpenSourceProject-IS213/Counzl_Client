package app

import (
	"./modules/users"
	"bufio"
	"fmt"
	"os"
	"strings"
	"utilities/converter"
	"utilities/cmd"
)

const STDIN_DESCRIPTION, STDIN_DESCRIPTION2 = "Password: ", "Confirm password: "

func SignUp(online bool) {
	var tryAgain = online
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Choose username: ")
	u, _ := input.ReadString('\n')
	brukernavn := converter.Remove_0a(strings.TrimSpace(u))
	fmt.Println()

	passord := cmd.Invisible_in(STDIN_DESCRIPTION)
	fmt.Println()
	p_compare := cmd.Invisible_in(STDIN_DESCRIPTION2)

	fmt.Println()
	if passord == p_compare {
		if online == true {
			id := users.CheckCerts_client(brukernavn)
			users.CreateUser(converter.Remove_0a(id), brukernavn, passord, "true")
			Welcome(tryAgain)
		} else {
			users.CreateUser("", brukernavn, passord, "false")
			Welcome(tryAgain)
		}

	} else {
		fmt.Println("The passwords do not correlate!")
		SignUp(tryAgain)
	}
}
