package main

import "fmt"

func main() {
	// Generic คือสามารถรับได้ทั้ง int, float, string หรืออะไรก็ได้
	// เพื่อไม่ให้เกิด duplicate function

	numsInt := []int{1, 2, 3}
	numsFloat64 := []float64{1.1, 2.2, 3.3}

	sumInt := sum(numsInt)
	sumFloat64 := sum(numsFloat64)

	fmt.Println(sumInt)
	fmt.Println(sumFloat64)

}

type (
	Number interface {
		int | float64
	}

	// Struct with Generics Type

	Player[T Number] struct {
		Name   string
		Age    int
		Damage T
	}

	Database[T Number] interface {
	}
)

// Type T
func sum[T Number](nums []T) T {
	var sum T

	for _, n := range nums {
		sum += n
	}

	return sum
}
