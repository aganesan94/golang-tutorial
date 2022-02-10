package main

import (
	"fmt"
	"strings"
)

func printHelloWorld() {
	fmt.Print("Hello World\n")
}

func main() {

	// Demonstrates function can be declared before main
	printHelloWorld()

	// Demonstrates function can be declared after main
	printHiThere()

	// Declaration of type of variable
	var firstName string
	var lastName string
	var email string

	// Message to display before asking for input.
	fmt.Println("First Name:")
	fmt.Scanln(&firstName)

	fmt.Println("Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Email: ")
	fmt.Scanln(&email)

	// validate user input
	isNameValid, isEmailValid := validateUserInput(firstName, lastName, email)

	if isNameValid && isEmailValid {
		fmt.Printf("Let us send %v, %v an email at %v", lastName, firstName, email)
	} else {
		fmt.Printf("Something is wrong with your first name, last name and email. Try again")
	}

}

func printHiThere() {
	fmt.Print("Hi There")
}

// Multipe return types can be specified.
func validateUserInput(firstName string, lastName string, email string) (bool, bool) {
	isNameValid := len(firstName) >= 2 && len(lastName) >= 2
	isEmailValid := strings.Contains(email, "@")
	return isNameValid, isEmailValid
}
