# PasswordCheck
A Go Lang implementation of a CLI password checker in Windows

* *setPass.go* - allows you to set a password, which will be stored in a txt file
  * password is hashed by **SHA512** hash function and PBKDF2 algorithm

## Notes

* Needs a Go Lang Compiler (https://golang.org/dl/)
* Set *Go* to *PATH*

## Running

* Run setPass.go (by typing *go run setPass.go* in command line)
  * The salt of PBKDF2 can be chaged manually in code 
* If some libraries are not installed, type: *go get \[library name\]*
* It will create a file 'index.txt' with the salt and password 
* Then run checkPass.go (by typing *go run checkPass.go* in command line)

* You can create an executable by building it:
  (go build checkPass.go)
