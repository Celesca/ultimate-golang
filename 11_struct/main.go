package main

import (
	"encoding/json"
	"fmt"
)

// embedded ร่วมกับ field ต่างๆ ให้ประกาศเป็น Public
type Player struct {
	Username string `json:"username"`
	Level    uint   `json:"level"`
	Status   string `json:"status"`
	Class    string `json:'class"`
}

func (p Player) Getusername() string {
	return p.Username
}

// use pointer to change value
// func (p Player) LevelUp() {} ถ้าแบบนี้มันจะ copy ค่าโดยปกติเข้าไป ฉะนั้นต้องใช้ pointer
func (p *Player) LevelUp() {
	p.Level++
}

func main() {

	//Create instant of struct
	player1 := Player{
		Username: "JohnDoe",
		Level:    1,
		Status:   "active",
		Class:    "warrior",
	}

	// Pass by reference (any คือ interface)
	jsonData, _ := json.MarshalIndent(&player1, "", "\t")
	// []Byte to string
	fmt.Println(string(jsonData))

	fmt.Println(player1.Getusername())

	player1.LevelUp()

	fmt.Println(player1.Level)

}
