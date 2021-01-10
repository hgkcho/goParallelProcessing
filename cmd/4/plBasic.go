package main

import "fmt"

func plBasic() {
	generator := func(done <-chan interface{}, d ...int) <-chan int {
		intStream := make(chan int, len(d))
		go func() {
			defer close(intStream)
			for _, v := range d {
				select {
				case <-done:
					return
				case intStream <- v:
				}
			}
		}()
		return intStream
	}
	multiply := func(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
		multiplyStream := make(chan int)
		go func() {
			defer close(multiplyStream)
			for i := range intStream {
				hage := i * multiplier
				select {
				case <-done:
					return
				case multiplyStream <- hage:
				}
			}
		}()
		return multiplyStream
	}

	add := func(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
		addStream := make(chan int)
		go func() {
			defer close(addStream)
			for i := range intStream {
				hage := i + additive
				select {
				case <-done:
					return
				case addStream <- hage:
				}
			}
		}()
		return addStream
	}

	done := make(chan interface{})
	defer close(done)

	data := []int{1, 2, 3, 4, 5, 7}
	intStream := generator(done, data...)
	pl := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	for v := range pl {
		fmt.Println(v)
	}

}
