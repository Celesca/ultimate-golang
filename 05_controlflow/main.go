package main

import "fmt"

func main() {
	n := 3

	if n%2 == 0 {
		fmt.Println("is even")
	} else {
		fmt.Println("is not even")
	}

	score := 60
	if score >= 80 {
		fmt.Println("A")
	} else if score >= 70 {
		fmt.Println("B")
	} else if score >= 60 {
		fmt.Println("C")
	} else {
		fmt.Println("F")
	}

	// a is not global variables.
	if a := 10 > 20; a {
		fmt.Println("10 is greater than 20")
	} else {
		fmt.Println("10 is not greater than 20")
	}
}
