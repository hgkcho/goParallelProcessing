package main

import (
	"fmt"
	"time"
)

func sleep(done, valueStream <-chan interface{} , d time.Duration) <-chan interface{} {
	sleepStream := make(chan interface{})
	go func() {
		defer close(sleepStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case sleepStream <- v:
				time.Sleep(d)
			}
		}
	}()
	return sleepStream
}

func queueTest() {
	done := make(chan interface{})
	defer close(done)
	start := time.Now()

	zeros := take(done, repeat(done,0),3)
	short := sleep(done, zeros, 1 * time.Second)
	long := sleep(done, short, 4 * time.Second)
	for v := range long {
		fmt.Println(v)
		fmt.Printf("Done: %v\n", time.Since(start))
	}
}