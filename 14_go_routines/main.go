package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("Hello main thread 1")

	time.Sleep(1 * time.Second)

	// hello() -> Synchronous programming มันจะมาไม่ถึง จะรอให้ hello() ทำงานเสร็จก่อน
	go hello() // -> Asynchronous programming มันจะไม่รอ concurrency ตัวลูก

	// main จะทำงานปกติ มันจะคอลก็ต่อเมื่อสองอันบนทำงานเสร็จแล้ว

	fmt.Println("Hello main thread 2")

	// ข้อจำกัดของ goroutine คือ มันยังคงไล่จากบนลงล่างเสมอ

	var wg sync.WaitGroup
	wg.Add(3)

	for i := range []int{1, 2, 3} {
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println("Hello from another thread", i)
		}(i, &wg)
	}

	wg.Wait()

}

func hello() {
	// ทำงานไปเรื่อยๆ
	for {
		fmt.Println("Hello from another thread")
	}
}
