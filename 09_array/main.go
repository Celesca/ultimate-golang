package main

import "fmt"

func main() {

	// array must declared fix size
	var a [3]int = [3]int{1, 2, 3}
	a[0] = 10
	fmt.Print("A : ")
	for i, v := range a {
		fmt.Print(i, v, " ")
	}

	// slice is dynamic array
	b := []int{1, 4, 5}
	fmt.Print(" B : ")
	for _, v := range b {
		fmt.Print(v, " ")
	}

	c := make([]int, 0) // Index at 0
	c = append(c, 1, 2, 3)
	fmt.Println(" C : ", c)

	// Array is the pointer, it is copied not pass by reference

	x := [3]int{1, 2, 3}
	double(x)

	fmt.Println("Inside main : ")
	for i := range x {
		fmt.Printf("%v\n", &x[i])
	}

	// However slice is pass by reference

	y := []int{1, 2, 3}
	z := double_slice(y)
	fmt.Println(y)
	fmt.Println(z)

}

func double(nums [3]int) {
	fmt.Println("Inside double : ")
	for i := range nums {
		fmt.Printf("%v\n", &nums[i])
		nums[i] *= 2
	}
}

// ถ้าไม่อยากให้มันยุ่งกับค่าด้านใน เราจะใช้ค่ามารับแทน
// ก็คือสร้างตัวอื่นมารับแทน
// In slice it will Pass by reference by default

func double_slice(nums []int) []int {
	newNums := make([]int, len(nums)) // new pointer reference will cause panick
	// ต้องจองให้มีขนาดเท่ากับตัวที่เราจะ copy มา
	for i := range nums {
		newNums[i] = nums[i] * 2
	}

	return newNums
}
