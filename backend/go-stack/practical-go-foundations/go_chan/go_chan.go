package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	fmt.Println(sleepSort([]int{3, 8, 2, 4, 1, 9, 0, 1}))

	for i := 0; i < 3; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(time.Millisecond)

	messages := make(chan string, 2)
	go func() {
		messages <- "ping"
		messages <- "pong"
	}()
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	close(messages)
	fmt.Println(<-messages)
	messages <- "late?"
}

func sleepSort(vs []int) (sorted []int) {
	comms := make(chan int, len(vs))
	for _, v := range vs {
		v := v
		go func() {
			time.Sleep(time.Duration(v) * time.Millisecond)
			comms <- v
		}()
	}

	for range vs {
		sorted = append(sorted, <-comms)
	}
	return sorted
}
