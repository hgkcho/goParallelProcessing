package main

import (
	"fmt"
	"time"
)


func okok() {
	ok := false
	cnt := 0
	unko := make(chan bool)
	go func() {
		defer close(unko)
		for {
			time.Sleep(1 * time.Second)
			unko <- false
			time.Sleep(1 * time.Second)
			unko <- true
		}
	}()
	for {
		select {
		case ok = <- unko:
			if ok {
				cnt++
				fmt.Println("OKOK")
			}
		}
		if cnt >=3 {
			return
		}
	}
}