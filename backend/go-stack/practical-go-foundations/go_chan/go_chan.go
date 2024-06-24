package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

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
}
