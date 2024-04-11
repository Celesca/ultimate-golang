package main

import "fmt"

type (
	Human struct{}

	Devil struct {
		Human // Embedded field
	}

	Elf struct {
		Human
	}
)

func (h Human) Walk() {
	fmt.Println("Walking")
}

func (h Human) Eat() {
	fmt.Println("Eating")
}

func (h Human) Breathing() {
	fmt.Println("Breathing")
}

func (d Devil) Mutate() {
	fmt.Println("Mutating")
}

func (e Elf) MagicSpell() {
	fmt.Println("Casting a magic spell")
}

func main() {

	human := Human{}
	devil := Devil{}
	elf := Elf{}

	fmt.Println("Human: ")
	human.Walk()
	human.Eat()
	human.Breathing()
	fmt.Println("")

	fmt.Println("Devil: ")
	devil.Walk()
	devil.Eat()
	devil.Breathing()
	devil.Mutate()
	fmt.Println("")

	fmt.Println("Elf: ")
	elf.Walk()
	elf.Eat()
	elf.Breathing()
	elf.MagicSpell()

}
