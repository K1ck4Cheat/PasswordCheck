package main

import (
	"fmt"
	"bytes"
	"syscall"
	"io/ioutil"
	"crypto/sha512"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/ssh/terminal"
)

func check(err error) {
	if err != nil {
		fmt.Printf("File Pass not found!")
		panic(err)
	}
}

func main() {

	fileContent, err := ioutil.ReadFile("input.txt")
	check(err)

	salt := append(fileContent[:32])
	realPass := append(fileContent[33:])
	// fmt.Printf("salt: %s\npass: %U\n", salt, realPass)

	fmt.Printf("Input the password: ")
	plainPass, err := terminal.ReadPassword(int(syscall.Stdin))
	check(err)

	inputPass := pbkdf2.Key(plainPass, salt, 100000, sha512.Size, sha512.New)
	// fmt.Printf("Wroten pass = %U\n", inputPass)

	ress := bytes.Compare(realPass, inputPass)
	if ress == 0 {
		fmt.Println("\n\nWelcome home, boy!")
	} else { 
		fmt.Println("\n\nOups, sometring went wrong!") 
	}
	
	fmt.Println("\n\nPress ENTER to exit")
	fmt.Scanln()
}