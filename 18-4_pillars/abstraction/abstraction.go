package main

import "fmt"

type (
	Keyboard interface {
		Input()
	}

	MechanicalKeyboard struct {
		SwitchType string
		Size       string
		OS         string
	}

	NormalKeyboard struct {
		IsCanPress bool
	}
)

func (m MechanicalKeyboard) Input() {
	fmt.Println("Pressing the key on the mechanical keyboard: ", m.SwitchType, m.Size, m.OS)
}

func (m NormalKeyboard) Input() {
	fmt.Println("Pressing the key on the keyboard")
}

func main() {
	mechanicalKeyboard := MechanicalKeyboard{
		SwitchType: "Cherry MX Blue",
		Size:       "Full-size",
		OS:         "Windows",
	}

	normalKeyboard := NormalKeyboard{
		IsCanPress: true,
	}

	// Abstraction the name of function is the same but the implementation is different อยู่ที่มุมมอง Abstraction อยู่ในหัวที่เราคิดหลายแบบ
	// มองสิ่งเดียวกัน แต่หัวคิดต่างกัน
	mechanicalKeyboard.Input()
	normalKeyboard.Input()
}
