package main

import "fmt"

func main() {
	depthFirstSearch(A, G)
}

const (
	A = iota
	B
	C
	D
	E
	F
	G
	Size
)

var adjacent= [][]int{
	{B, C},     // A
	{A, C, D},  // B
	{A, B, E},  // C
	{B, E, F},  // D
	{C, D, G},  // E
	{D},        // F
	{E},        // G
}


func member(n int, xs []int) bool {
	for _, x := range xs{
		if n ==x {return true}
	}
	return false
}

func dfs(goal int, path []int) {
	n := path[len(path) - 1]
	if n == goal {
			fmt.Println(path)
	} else {
			for _, x := range adjacent[n] {
					if !member(x, path) {
							dfs(goal, append(path, x))
					}
			}
	}
}

func depthFirstSearch(start, goal int) {
	path := make([]int, 0, Size )
	dfs(goal, append(path, start))
}