package main

import "fmt"

func main() {

	// Declaring an array
	colors := []string{}
	colors = append(colors, "red", "green", "blue")
	fmt.Printf("Colors Array contains %v \n", colors)
	fmt.Printf("Total Length %v \n", len(colors))

	//Iterating over an array
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// Defining a predefined array
	integerArray := [5]int{1, 2, 3, 4, 5}

	//Iterating over an array
	for i := 0; i < len(integerArray); i++ {
		fmt.Println(integerArray[i])
	}

}
