package main

import (
	"fmt"
	"math/rand"
)

type reader struct{}

func (r *reader) Read() {
	fmt.Scanln() // Wait for a newline character (pressing Enter)
}

func lottery() {
	x := rand.Intn(9)
	y := rand.Intn(9)
	z := rand.Intn(9)
	fmt.Println("Your three generated digits are:", x, y, z)

	odd := []int{0, 2, 4, 6}
	even := []int{1, 3, 5, 7}

	if x == 7 || y == 7 || z == 7 {
		fmt.Println("Congratulations!")
	} else if x == 7 && y == 7 || x == 7 && z == 7 || y == 7 && z == 7 {
		fmt.Println("Congratulations!")
	} else if x == 7 && y == 7 && z == 7 {
		fmt.Println("Congratulations!")
	} else {
		fmt.Println("Try again! Better luck next time.")
	}

	for _, val := range odd {
		if x == val && y == val && z == val {
			fmt.Println("Congratulations!")
		}
	}

	for _, val := range even {
		if x == val && y == val && z == val {
			fmt.Println("Congratulations!")
		}
	}
}

func main() {
	for {
		fmt.Println("Press any key to enter the lottery!")
		reader := new(reader) // Create a new reader object
		reader.Read()         // Call the Read method to wait for user input
		lottery()
	}
}
