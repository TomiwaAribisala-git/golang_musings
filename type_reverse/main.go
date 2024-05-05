package main

import (
	"fmt"

	"github.com/zhexuany/wordGenerator"
	"golang.org/x/example/hello/reverse"
)

func main() {
	for {
		fmt.Println("Welcome to the type reverse game!")
		randomWord := wordGenerator.GetWord(5)
		result := reverse.String(randomWord)
		fmt.Println(result)

		var input string
		fmt.Printf("Type the reverse of %s:", randomWord)
		fmt.Scan(&input)

		if input == result {
			fmt.Println("You are correct!")
		} else {
			fmt.Println("You are wrong, try again!")
		}
	}
}
