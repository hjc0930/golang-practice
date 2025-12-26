package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "hello world"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	}
}
