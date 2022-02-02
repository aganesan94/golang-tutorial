package main

import "fmt"

func main() {

	// Declaring a constant, cannot be changed
	// once assigned.
	const totalCourses int = 50

	// Variable, the value can be modfied
	// at a later point as well
	var coursesUnderReview uint = 20

	//assignments for variables can be done as follows
	myName := "Arun Ganesan"

	// More about the types can be reviewed here https://pkg.go.dev/fmt
	fmt.Printf("Welcome to %v's course. \n, There are a total of %v courses to pick and choose from out of which %v courses are still under review.", myName, totalCourses, coursesUnderReview)

}
