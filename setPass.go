package main

import (
	"fmt"
	"io/ioutil"
	"crypto/sha512"
	"golang.org/x/crypto/pbkdf2"
)
	
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	var plainPass []byte 
	salt := []byte("U4lT3uwDiEcpNrFH0EYyP2sOYd8VEOy0")

	fmt.Printf("Input the real password: ")
	fmt.Scanf("%s", &plainPass)

	realPassKey := pbkdf2.Key(plainPass, salt, 100000, sha512.Size, sha512.New)

	password := []byte(string(salt) + "\n" + string(realPassKey))

	// fmt.Printf("Password: %U", password) 

	err := ioutil.WriteFile("input.txt", password, 0644)
	check(err)

}
