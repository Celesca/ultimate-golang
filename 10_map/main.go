package main

import "fmt"

var graph = map[int][]int{
	1: {2, 3},
	2: {4, 5},
	3: {6},
}

func main() {
	a := map[int]int{
		1: 1,
		2: 2,
		3: 10,
	}

	if _, ok := a[100]; ok {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}

	fmt.Println(a[100]) // Default value ของแต่ละ Type

	// Initialze empty map

	b := make(map[int]int) // สร้าง map ว่างๆ

	b[1] = 123

	// Implementing the graph using map
	dfs(graph, 1, make(map[int]bool))

}

func dfs(graph map[int][]int, vertex int, visited map[int]bool) {
	visited[vertex] = true
	for _, v := range graph[vertex] {
		fmt.Printf("->%d", v)
		dfs(graph, v, visited)
	}
}
