package main

import (
	"fmt"
)

func main() {

	var firstName string
	var lastName string

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	if firstName == "Arun" {
		fmt.Print("First Names are equal")
	}

	if firstName == "Don" && lastName == "Quixote" {
		fmt.Print("firstName and lastName are equal")
	}

	if firstName == "Don" {
		fmt.Print("You got the first name right")
	} else {
		fmt.Print("You got the first name wrong!!!")
	}

}
