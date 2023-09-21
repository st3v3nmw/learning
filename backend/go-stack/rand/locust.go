package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	statusCode int
	timeTaken  int64
}

const N_GOROUTINES = 500
const TEST_TIME = 2
const URL = "http://example.com/"

func main() {
	var wg sync.WaitGroup
	results := make(chan Result, 100)

	fmt.Printf(
		"Running %ds tests against %s using %d goroutine(s) ... ",
		TEST_TIME,
		URL,
		N_GOROUTINES,
	)

	statusCodes := make(map[int]int)
	nRequests := 0.0
	var totalTimeTaken int64

loop:
	for timeout := time.After(TEST_TIME * time.Second); ; {
		wg.Add(1)

		select {
		case <-timeout:
			fmt.Println("DONE")
			break loop
		case result := <-results:
			statusCodes[result.statusCode] += 1
			totalTimeTaken += result.timeTaken
			// nRequests += 1

			// if !more {
			// 	close(results)
			// } else {
			// 	fmt.Println("true")
			// }
		default:
		}

		go func() {
			defer wg.Done()

			started := time.Now().UnixMilli()
			resp, err := http.Get(URL)
			timeTaken := time.Now().UnixMilli() - started

			if err != nil {
				results <- Result{1000, timeTaken}
			} else {
				results <- Result{resp.StatusCode, timeTaken}
			}
		}()

		nRequests += 1
		println(nRequests)
	}

	wg.Wait()

	fmt.Println(statusCodes, float64(totalTimeTaken)/nRequests, nRequests)
}
