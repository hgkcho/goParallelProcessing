package main

import "fmt"

// レキシカル拘束: コンパイルを駆使して拘束を強制する。レキシカルスコープを使って得知恵のスコープだけに影響範囲を留める
func rexical() {
	// ここでresultsチャネルへの書き込みを拘束している
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	// consumerはresultsを読み込みだけできる。
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("received: %d\n", result)
		}
		fmt.Println("Done")
	}

	results := chanOwner()
	consumer(results)
}
