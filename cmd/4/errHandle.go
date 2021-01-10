package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Err      error
	response *http.Response
}

func checkStatus(done <-chan interface{}, urls ...string) <-chan Result {
	resultStream := make(chan Result)
	go func() {
		defer close(resultStream)
		for _, url := range urls {
			var result Result
			resp, err := http.Get(url)
			if err != nil {
				result.Err = err
			}
			result.response = resp
			select {
			case <-done:
				return
			case resultStream <- result:
			}
		}
	}()
	return resultStream
}

func goodErrHandle() {
	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://a.com", "b", "c", "d"}
	errCount := 0
	for result := range checkStatus(done, urls...) {
		if result.Err != nil {
			fmt.Printf("err: %v\n", result.Err)
			errCount++
			fmt.Printf("errCount: %d\n", errCount)
			if errCount >= 3 {
				fmt.Println("fuck")
				break
			}
			continue
		}
		fmt.Printf("response status: %v\n", result.response.Status)
	}
}

func badErrHandle() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan *http.Response {
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					continue
				}
				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()
		return responses
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://a.com"}
	for resp := range checkStatus(done, urls...) {
		fmt.Printf("reponse: %v\n", resp.Status)
	}
}
