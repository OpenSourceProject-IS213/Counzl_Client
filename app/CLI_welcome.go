package app

import (
	"./modules/users"
	"gopkg.in/abiosoft/ishell.v2"
	"utilities/cmd"
)

// Denne filen skal ha muligheten til å launche enhver modul.
func Welcome(online bool) {
	shell := ishell.New()
	shell.Print("\n" + cmd.ChangeColor("Welcome to our Open Source project, Counzl!", "brightBlue") + "                                                  ")
	shell.Print("\n" + "If this is your first time; type '" + cmd.ChangeColor("hjelp", "green") + "'!\nThis will give you a list of commands :)")
	shell.Println()
	// register a function for "greet" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "loading",
		Help: "Test graphics for 'loading'",
		Func: func(c *ishell.Context) {
			cmd.Loading()
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "printusers",
		Help: "Print all the users registered in this program (all the users on your own PC)",
		Func: func(c *ishell.Context) {
			users.PrintLocalUsers()
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "logout",
		Help: "Log out",
		Func: func(c *ishell.Context) {
			AccessCounzl()
		},
	})

	shell.Run()
}
