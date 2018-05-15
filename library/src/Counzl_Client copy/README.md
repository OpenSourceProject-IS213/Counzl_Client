# Counzl_Client

## **The purpose of our Counzl_Client**
The purpose of 'Counzl' is to have a template for a program in go that should be easy to develop further (by adding more modules) for both beginners and experienced programmers. Currently there is just one modules now: authentication (login/sign up) and user management (except deleting users).   

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

## Adding Modules


