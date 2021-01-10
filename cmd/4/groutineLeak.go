package main

import (
	"fmt"
	"math/rand"
	"time"
)

func goroutineLeakConsumer() {
	do := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("do exit!")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Printf("get from strings: %s\n", s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	hage := func() <-chan string {
		strings := make(chan string, 3)
		d := []string{"hello", "haha", "unko"}
		for _, v := range d {
			strings <- v
		}
		return strings
	}

	done := make(chan interface{})
	terminated := do(done, hage())

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("cancelling do goroutine ...")
		close(done)
	}()

	<-terminated // 受信街
	fmt.Println("All Done.")
}

func goroutineLeakProducer() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStrream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exit")
			defer close(randStrream)
			for {
				select {
				case randStrream <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randStrream
	}
	done := make(chan interface{})
	rs := newRandStream(done)
	for i := 0; i <= 6; i++ {
		fmt.Printf("%d: %d \n", i, <-rs)
	}
	close(done)
	time.Sleep(1 * time.Second)
}
