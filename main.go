package main

import (
	"fmt"
	"time"
)

func main() {
	//working("chase")

	//go working("dune")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	messages := make(chan string)
	go func() {
		messages <- "maker"
	}()
	msg := <-messages
	fmt.Println(msg)

	time.Sleep(time.Second)
	fmt.Println("DONE")
}
