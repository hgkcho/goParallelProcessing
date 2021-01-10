package main

import "fmt"

func repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

func repeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	stream := make(chan interface{})
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()
	return stream
}

func toString(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case stringStream <- v.(string):
			}
		}
	}()
	return stringStream
}

func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream: //valueStreamはポインタなので <-valueStream で値を取り出す
			}
		}
	}()
	return takeStream
}

func plConvenient() {
	done := make(chan interface{})
	defer close(done)

	// stream := repeat(done, 1, 2, 3, 4)
	// for n := range take(done, stream, 10) {
	// 	fmt.Printf("%v ", n)
	// }

	// rd := func() interface{} { return rand.Int() }
	// for num := range take(done, repeatFn(done, rd), 10) {
	// 	fmt.Println(num)
	// }

	var msg string
	for token := range toString(done, take(done, repeat(done, "I", "am."), 5)) {
		msg += token
	}
	fmt.Printf("msg: %s\n", msg)

}
