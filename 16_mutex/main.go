package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu      sync.Mutex
	balance int = 0
)

func main() {

	doneCh := make(chan bool, 3)

	go UpdateBalance(doneCh, 100)
	go UpdateBalance(doneCh, 50)
	go UpdateBalance(doneCh, 200)

	<-doneCh
	<-doneCh
	<-doneCh

	fmt.Println(balance)

}

// มันจะรัน Transaction ทีละอัน
func UpdateBalance(doneCh chan<- bool, amount int) {

	mu.Lock()
	fmt.Println("Updating balance")
	time.Sleep(1 * time.Second)
	balance += amount
	doneCh <- true // พร้อมที่จะ return

	mu.Unlock()
	fmt.Println("Balance updated")
}

// การป้องกัน Racing Condition ด้วย Mutex เช่น การล็อคกล่องใน Minecraft คนนึงกำลังเปิดอยู่
