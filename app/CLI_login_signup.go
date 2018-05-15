package app

import (
	"gopkg.in/abiosoft/ishell.v2"
	"utilities/cmd"
)

func AccessCounzl() {
	shell := ishell.New()
	// display welcome info.
	shell.Print(cmd.ChangeColor("Sign in, sign up or sign up locally (sign up locally will not give you access to the server)", "blue"))
	shell.Print("\n" + cmd.ChangeColor("signin", "green") + " - Sign in; doesn't require a connection to the server")
	shell.Print("\n" + cmd.ChangeColor("signup", "green") + " - Online; You have to be connected to the server\n")
	shell.Println( cmd.ChangeColor("local_signup", "green") + " - Locally; doesn't require a connection to the server")

	shell.AddCmd(&ishell.Cmd{
		Name: "signup",
		Help: "You have to register online, we are just going to store the following info: " + cmd.ChangeColor("User ID", "red") + ", " + cmd.ChangeColor("Username", "red") + " & " + cmd.ChangeColor("A timestamp to indicate when you created the user", "red"),
		Func: func(c *ishell.Context) {
			SignUp(true)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "local_signup",
		Help: "Sign up locally, you will not have access to the online features",
		Func: func(c *ishell.Context) {
			SignUp(false)
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "signin",
		Help: "Sign in to use Counzl",
		Func: func(c *ishell.Context) {
			Login()
		},
	})
	shell.Run()
}
