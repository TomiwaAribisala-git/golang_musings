package main

import (
	"fmt"
)

func main() {
	table := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(table[0], table[1], table[2], table[4], table[5], table[6], table[7], table[8], table[9], table[10])
	for i := 2; i <= 12; i++ {
		fmt.Println(table[0]*i, table[1]*i, table[2]*i, table[4]*i, table[5]*i, table[6]*i, table[7]*i, table[8]*i, table[9]*i, table[10]*i)
	}
	fmt.Println()
}
