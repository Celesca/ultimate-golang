package main

import "fmt"

type (
	Aircraft interface {
		Fly()
	}

	FighterAircraft interface {
		Aircraft
		Fire()
	}

	CommercialAircraft interface {
		Aircraft
		AirCondition()
	}

	Boeing747 struct {
		CommercialAircraft
	}
	Spitfire struct {
		FighterAircraft
	}
)

func (b *Boeing747) Fly() {
	fmt.Println("Fly")
}
func (b *Boeing747) AirCondition() {
	fmt.Println("AirCondition")
}
func (b *Spitfire) Fly() {
	fmt.Println("Fly")
}
func (b *Spitfire) Fire() {
	fmt.Println("Fire")
}

func main() {
	// การออกแบบ Interface ให้มีความเฉพาะเจาะจง ไม่ควรมี method ที่ไม่ได้ใช้
	spitfire := &Spitfire{}
	boeing747 := &Boeing747{}

	spitfire.Fly()
	spitfire.Fire()
	boeing747.Fly()
	boeing747.AirCondition()
}
