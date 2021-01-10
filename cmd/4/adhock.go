package main

import "fmt"

// アドホック拘束: 規約によって拘束。チーム規模が大きくなるにつれて難しくなる。静的解析ツールとか必要になる。
func adhock() {
	data := make([]int, 4)
	for i := 0; i < 4; i++ {
		data[i] = i + 3
	}

	ld := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go ld(handleData)
	for num := range handleData {
		fmt.Println(num)
	}
}
