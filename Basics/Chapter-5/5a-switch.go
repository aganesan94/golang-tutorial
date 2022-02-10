package main

import (
	"fmt"
)

func main() {

	// Switch statement example
	city := "London"

	switch city {
	case "New York":
		// booking for New York conference
	case "Singapore", "Hong Kong":
		// booking for Singapore & Hong Kong conference
	case "London", "Berlin":
		// booking for London & Berlin conference
	case "Mexico City":
		// booking for Mexico City conference
	default:
		fmt.Print("No valid city selected")
	}
}
