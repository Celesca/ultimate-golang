package main

import "fmt"

// Liskov - function ของลูกไม่ควรจะขัดกับตัวแม่
type (
	Aircraft interface {
		Fly()
	}

	Boeing747 struct {
		Aircraft
	}

	Spitfire struct {
		Aircraft // Embedded หรือ Inheritance ก็ได้
	}
)

func (b *Boeing747) Fly() {
	fmt.Println("Fly")
}

func (b *Spitfire) Fly() {
	fmt.Println("Fly")
}

func (b *Spitfire) Fire() {
	fmt.Println("Fire")
}

func main() {

	spitfire := &Spitfire{}
	boeing747 := &Boeing747{}

	spitfire.Fly()
	spitfire.Fire()
	boeing747.Fly()

}
