package main

import "fmt"

type (
	PlayerClass interface {
		Attack()
	}

	Warrior struct {
		Weapon string
	}

	Mage struct {
		Spell string
	}
)

func (w Warrior) Attack() {
	fmt.Println("Warrior is attacking with", w.Weapon)
}

func (m Mage) Attack() {
	fmt.Println("Mage is attacking with", m.Spell)
}

func PlayerAttack(playerClass PlayerClass) {
	// ข้างในมันต้องทำงาน method attack ได้
	playerClass.Attack()
}

func main() {
	warrior := Warrior{
		Weapon: "Sword",
	}

	mage := Mage{
		Spell: "Fireball",
	}

	// warrior.Attack()
	// mage.Attack()

	PlayerAttack(warrior)
	PlayerAttack(mage)
}
