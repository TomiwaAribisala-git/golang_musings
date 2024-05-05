package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/example/hello/reverse"
)

func main() {
	for {
		word := bufio.NewReader(os.Stdin)
		fmt.Println("Please enter a word: ")

		input, _ := word.ReadString('\n')
		fmt.Println(reverse.String(input))

		revsentence := bufio.NewReader(os.Stdin)
		fmt.Println("Please enter a sentence: ")

		input1, _ := revsentence.ReadString('\n')
		fmt.Println(reverse.String(input1))
	}
}
