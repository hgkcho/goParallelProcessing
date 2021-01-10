package main

import (
	"testing"
)

func BenchmarkGeneric(b *testing.B) {
	done := make(chan interface{})
	defer close(done)
	b.ResetTimer()
	for range toString(done, take(done, repeat(done, "a"), b.N)) {
	}
}

func BenchmarkTyped(b *testing.B) {
	repeat := func(done <-chan interface{}, values ...string) <-chan string {
		valueStream := make(chan string)
		go func() {
			defer close(valueStream)
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

	take := func(done <-chan interface{}, stringStream <-chan string, num int) <-chan string {
		takeStream := make(chan string)
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-stringStream:

				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done)
	b.ResetTimer()
	for range take(done, repeat(done, "a"), b.N) {
	}
}
