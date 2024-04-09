package main

import "fmt"

func main() {
	name := hello("Folk")
	fmt.Println(name)

	newName := "Celesca"

	// Anonymous function can access global variable
	// Didn't copy in the parameters but access the global variable
	// Goroutines with anonymous function be carefully.

	func(name string) {
		fmt.Println("Hello ", name)
	}(newName)

	func() {
		newName = "Folk"
		fmt.Println(newName)
	}()

	result := func(a, b int) int {
		return a + b
	}(1, 2)

	fmt.Println(result)

	// Anonymous shouldn't be use in normally. It's hard with global variables.

}

func hello(name string) string {
	return "Hello " + name
}

// copy value into function without updating the original value

func add(a, b int) int {
	a += 1
	return a + b
}

func minus(a, b int) int {
	return a - b
}
