package main

import (
	"fmt"
	"time"
	//"rsc.io/quote"
)

func main() {
	//var name = "tomiwa"
	//const ticketPrice = 70
	//fmt.Println("My name is", name, "and my ticket price is", ticketPrice, "dollars")
	//fmt.Println(quote.Go())
	defer bran()
	sed()
	hon()
}

func sed() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second) // 1 second
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second) // 2 second
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func hon() {
	var matchPrice = 50
	fmt.Println(&matchPrice) // &matchPrice is the memory address of matchPrice
}
