package main

import (
	"fmt"
)

func main() {

	// Loop with a range
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
		fmt.Println(sum)
	}
	fmt.Println(sum)

	// loop lesser than range
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n)

	//infinite loop
	counter := 0
	for {
		counter++
		fmt.Println(counter)
	}
}
