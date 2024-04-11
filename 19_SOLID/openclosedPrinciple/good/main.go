package main

import "fmt"

type (
	Player struct{}
	Item   struct {
		Name      string
		Abilities []Ability
	}
	Heal        struct{}
	DamageBluff struct{}

	Ability interface {
		Execute()
	}
)

// struct player เวลาผ่าน Item เข้ามา ดู ability ที่มี แล้วให้ทำการ execute ตาม ability นั้นๆ
func (p *Player) Use(item Item) {
	fmt.Println("Use : ", item.Name)

	for _, item := range item.Abilities {
		item.Execute()
	}
}

func (h Heal) Execute() {
	fmt.Println("Heal")
}

func (d DamageBluff) Execute() {
	fmt.Println("Increase Damage 100%")
}

func main() {
	p := &Player{}
	steak := Item{
		Name: "Steak",
		Abilities: []Ability{
			Heal{},
			DamageBluff{},
		},
	}

	p.Use(steak)
}

// เราจะไม่ต้องมาแก้ตัว struct มันคือ composite Pattern เพิ่มใหม่ลงใน Client ได้เลย
// และเราสามารถเพิ่ม Ability ใหม่ๆ ได้โดยไม่ต้องแก้ struct ที่มีอยู่แล้ว
