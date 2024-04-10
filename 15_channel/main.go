package main

import "fmt"

// Channel เอาไว้ใช้กับ Goroutines และเอาไว้เก็บส่งข้อมูลได้
// ขั้นสูงกว่า wait โดยปกติมันไม่รู้ว่าจะต้องจบตอนไหน รับค่าตอนไหน
// channel buffered size เป็น 10 หมายถึงมันจะมี buffer ไว้ 10 ค่า

// close คือจะไม่สามารถรับ ส่งค่าได้อีก ป้องกันการ Deadlock ได้

// receive only channel รับได้อย่างเดียว ต้องใช้ chan <- แทน <- chan

// Go wait ไม่รู้จะจบตอนไหน แต่ถ้ามันครบรอบแล้ว มันก็จะหยุด
// function use case ที่มีการรับส่งข้อมูล ควรใช้ channel แทน wait

func main() {

	jobCh := make(chan int, 10)
	resultCh := make(chan int, 10)

	for i := range 10 {
		jobCh <- i + 1
	}

	close(jobCh) // กัน deadlock ทุกครั้งที่มี assign แล้วต้อง close ไม่รับค่าแล้ว ครบ buffered

	double(jobCh, resultCh)
	for range 10 {
		result := <-resultCh
		fmt.Println(result)
	}
}

//Channel เป็น pointer ที่ชี้ไปที่ address ของข้อมูล

func double(jobCh <-chan int, resultCh chan<- int) {
	for j := range jobCh {
		resultCh <- j * 2
	}
}
