package main

import (
	"flag"
	"fmt"
)

func main() {

	// Get the email flag and parse it
	emailFlag := flag.String("email", "email@example.com", "an email string")
	flag.Parse()

	// Put the email flag in email variable
	email := *emailFlag
	// Add first two characters to stars variable
	encripedMail := email[:2]

	// Loop through email address
	for i := 2; i < len(email); i++ {
		// While character is not "@" put "*" instead of character
		if string(email[i]) != "@" {
			encripedMail += "*"
		} else {
			// Add the rest of email address and break the loop
			encripedMail += email[i:]
			break
		}
	}
	fmt.Println(encripedMail)
}
