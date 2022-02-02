package main

import "fmt"

func main() {

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

	fmt.Printf("Let us send %v, %v an email at %v", firstName, lastName, email)
}
