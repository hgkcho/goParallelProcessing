package main

import "sync"

func toInt(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case intStream <- v.(int):
			}
		}
	}()
	return intStream
}

// func badNaive() {
// 	rand := func() interface{} {return rand.Intn(500000000)}

// 	done := make(chan interface{})
// 	defer close(done)

// 	start := time.Now()

// 	randIntStream := toInt(done, repeatFn(done, rand))
// 	fmt.Println("Primes: ")
// 	for prime := range take(done, prinmeFinder(done,randIntStream), 10) {
// 		fmt.Printf("\t %d\n", prime)
// 	}
// 	fmt.Printf("Search took: %v", time.Since(start))
// }

func fanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})
	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	// TODO この処理をゴルーチンを起動してやる意味は？
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}
