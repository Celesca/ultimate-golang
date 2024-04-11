package main

import "fmt"

type Airplane struct {
	Name string // attribute or field
}

// method of airplane
func (a Airplane) Fly() {
	fmt.Printf("%s is flying\n", a.Name)
}

func main() {
	spitfire := Airplane{Name: "Spitfire"}
	spitfire.Fly()

	boeing := Airplane{Name: "Boeing 747"}
	boeing.Fly()

}
