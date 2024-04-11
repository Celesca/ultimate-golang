package main

import "fmt"

type (
	Player struct {
		Name string
	}
	Display struct{}
)

func (p *Player) Move() {
	fmt.Println("Move")
}

func (p *Player) Jump() {
	fmt.Println("Jump")
}

func (d *Display) ShowScoreBoard() {
	fmt.Println("ShowScoreBoard")
}

func main() {
	player := &Player{}
	display := &Display{}

	player.Move()
	player.Jump()

	display.ShowScoreBoard()

}
