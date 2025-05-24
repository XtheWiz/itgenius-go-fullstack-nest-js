package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello from goroutine!"
	}()

	fmt.Println("Waiting for message from goroutine...")
	msg := <-ch
	fmt.Println(msg)
	fmt.Println("Main function finished.")
}
