# Counzl_Client

## **The purpose of our Counzl_Client**
The purpose of 'Counzl' is to have a template for a program in go that should be easy to develop further (by adding more modules) for both beginners and experienced programmers. Currently there is just two modules: authentication (login/sign up) and user management (except deleting users).   

## Local and Online user type 
Our program is intended for online use, but the user can also create a local user. When the user is local it is stored without an ID and do not have access to any features that require connection to the server. 

# Running the program

We have tried to build the project to a binary file, by running go install project_name, but got an error message with some permission problems. We tried to troubleshoot by running this command to build a 'go bootstrap tool': 
``` bash 
$ sudo ./all.bash 
```
We got this error message even though my working tree is version 1.7.4:
``` bash
Set $GOROOT_BOOTSTRAP to a working Go tree >= Go 1.4
```
Unfortunately we do not know how to solve this problem, but the program works as it should by running this command:
``` bash
$ go run main.go
```

If you get an error message you have to change your GOPATH by running this command:
``` bash
export GOPATH=absolute/path/to/Counzl_Client/library
```
Note: to change the gopath permanently, insert the command above into ".profile" (or .bash_profile) located in the $HOME-folder and restart the terminal or execute the command: 
```bash
source .profile (or .bash_profile)
```
<br>

# Implemented functions Counzl_Client: 
Note: If a function starts with a capital letter it means that it is accessible outside of its package. <br>
If a function does not start with a capital letter, the it is only available in its package.
If you have any experience with Java you can see that this is very similar to 'public' and 'package-private'.  
## Overview:
To make the links work I had to write the headings in lower case. The correct function names are listed in this overview. The numbers represents go-files, the bullet point are functions.
1. [login.go](#login) <br>
* [Login](#Login)
* [check password](#checkpassword)
* [try again](#tryagain)
2. [sign_up.go](#sign_up) <br>
* [SignUp](#signup)
3. [user_management.go](#user_management) <br>
* [CreateUser](#createuser)
* [FetchUser](#fetchuser)
* [PrintLocalUsers](#printlocalusers)
* [serialize](#serialize)
* [deserialize](#deserialize)
4. [user_dialer.go](#user_dialer)
* [CheckCerts_client](#checkcerts_client)
* [dialMUX](#dialmux)
5. [iShell](#ishell)

## login

### Login
* This is the main function of the login.go. It reads the username and password inputs. 
* The password input is handled by cmd.Invisible_in() to make the input is invisible in the CLI (you will just see a key icon).
* The line feed ('**\n**' or '**\x0a**' closes the argument) at the end of the input arguments is removed by calling converter.Remove_0a(id). <br>
This is due to the fact that we do not want to store the username (or the password for that matter) with a line feed at the end. 
* Returns the result from the [checkPassword](#checkpassword)-function.

### checkPassword
* This function calls [FetchUser](#fetchuser) in user_management.go (we will get to this file in a bit) to see if the username is in the database and to fetch the hashed password. 
* The hashed password is compared to the password argument from Login() which is hashed as well. This is done by calling crypto.CheckPasswordHash(password_arg, hashed_pw_from_db) which return true are false. 
1. If the password is correct, the user gets access to Welcome() in CLI_welcome (here you can access all the different modules).
2. If the password is incorrect or the username does not exist in the DB, the tryAgain()-function is called

### tryAgain
* This function consists of an input from the stdin and a switch.
* The text before the input field asks if you want to try to login again. Input can just be 'y'(yes) or 'n'(no), if you type something else the function will tell you to type either 'y' or 'n' - nothing else. 
1. If yes: the Login()-function will be called and you can try again.
2. If no: the program will be killed with exit code 0.

## sign_up

### SignUp
* Reads username, password and password_compare from stdin - username and password is treated as in [Login](#login) (removal of '\n' and invisible password, etc.). 
* The password_compare is just to make sure the user typed the password correct ("type password again"-sort-of-thing.
* The user ID is generated on the server by calling [Check_Certs_client](), the idea is that if a user is created locally it does not have an ID and cannot access the modules that include networking - for example a chat module. 
1. If the password and password_compare is identical the function [Welcome](#Welcome) is called.
2. If they are not identical, SignUp is called (will cause an infinite loop if this happens every time).

## user_management
This file includes functions to store a user, print all the users in the local database and fetch **one** user.

### CreateUser
* Hashes the password.
* Creates user variable that points to the user struct. A user variable consists of ID, hashed password and online (if user is registered on the server). 
* [serialize](#serialize)s the user variable (this will allow you to store multiple values in one 'value' in the DB). 
* Opens the DB-file. This is standard procedure when you want to write or read to/from the DB-file. <br>
**NB!** The database has to be closed after a transaction to allow other processes to access the file.  
<br> 
* The update function is called and inserts the key (username) and the value (ID, hashed password and online as a byte slice). After the insertion the function returns an error to signal that the transaksjon is done. 

### FetchUser
* Open the DB-file.
* Calls the View-function. This will allow for multiple entries at once (which is one of the reasons why we chose buntDB).
* Searches through the file by the key (in this case: username) and retrieves one user. 
* [deserialize](#deserialize)s the value (as in hashmap value) to get all the different values (ID, etc.) and declares them as seperate values. For example:
```go
var pword = ""
var id = ""
var online = ""

err = db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(username)
		if err != nil {
			return err
		}
		value := deserialize([]byte(val))
		pword = value.Password
		id = value.ID
		online = value.online
		return nil
	})
``` 
This is an example of how powerful and handy the golang struct can be.  
### PrintLocalUsers
* This function iterates over the db-file and prints out all the users. 
### serialize
Writes the the user variable to a byte slice which it then returns. 
### deserialize
* Decodes the byte slice to a user variable (which again consists of ID, hashed password and online). 
* Returns a pointer to a distinct user variable. 
* If you want to print the ID you can just type: fmt.Printf("ID: %s" , user.ID). This will print the ID in string format. 

## iShell
CLI_login_signup.go and CLI_welcome.go represent two shell sections in the program. It is unnecessary to describe them into detail, screenshots will give you a better idea how iShell works. <br> 
Unfortunately the language is in norwegian, but it should be quite simple to relate the code to the screenshots. <br>

Note: the ChangeColor-function is located in the utilities package and changes the color of the text to the one that you desire (in this case "login" becomes green).
```go
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
```

**Screenshot:** <br>
![screenshot](https://github.com/BadNameException/Counzl_Client/blob/justanotherbranch/Screenshot_Counzl_Client.png)

## user_dialer
This file is one of the more complicated ones; since it includes key and certicate authentication. 

### CheckCerts_client
* Loads keypair (client.crt and client.key), returns error if the files do not exist, are expired, etc. 
* Reads the CA-certification and creates a CA certification pool.
* Creates a pointer to a TLS configuration which consists of the key, crt and CA certification pool (CA certificate).
* Calls the dialMUX with the tls configuration pointer as input. This authenticate the connection between the client and the server. 
* If a connection refuses to be established; the program will exit with exit code 1
 
### dialMUX
* Dials the server with connection type ("tcp" in this case), the address (IP and port number) and the TLS configuration pointer. 
* Writes username to the connection.
* Creates a buffer (this will prevent the whole message from being cached in RAM). 
* Reads from the connection, writes it to the buffer and prints out the message. 
