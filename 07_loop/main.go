package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	var sum int
	for i := 1; i <= 10; i++ {
		sum += i
	}

	// Go 1.12.2 +
	//for i := range 3 { // [0,2]
	//	fmt.Println(i)
	//}

}
