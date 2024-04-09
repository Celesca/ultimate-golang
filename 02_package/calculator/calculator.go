package calculator

import "fmt"

// Private because it starts with a lowercase letter
// func add(a, b int) int {
// 	return a + b
// }

// Public because it starts with a capital letter
func Add(a, b int) int {
	// ไม่ต้องประกาศชื่อ package ด้านหน้าเพราะว่าอยู่ใน package เดียวกัน
	fmt.Println("Multiply : ", multiply(a, b))
	return a + b
}
