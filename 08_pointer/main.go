package main

import "fmt"

func main() {
	b := new(int)
	c := new(*int)

	a := 1
	b = &a
	c = &b

	fmt.Printf("address a %v\n", &a)
	fmt.Printf("address b %v\n", &b)
	fmt.Printf("address c %v\n", &c)

	fmt.Println()

	fmt.Printf("value a %v\n", a)
	fmt.Printf("value b %v\n", b)
	fmt.Printf("value c %v\n", c)

	fmt.Println()

	fmt.Printf("*b %v\n", *b)
	fmt.Printf("*c %v\n", *c)
	fmt.Printf("**c %v\n", **c)

	// Pass by pointer
	d := 2
	var n *int = &d

	double(n)

	fmt.Println(*n)
	fmt.Println(d)

	// Pass by reference
	m := 2

	double(&m)

	fmt.Println(m)

}

func double(n *int) {
	*n *= 2
}
