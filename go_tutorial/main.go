package go_tutorial

import "fmt"

func main() {
	var name = "tomiwa"
	const ticketPrice = 70
	fmt.Println("My name is", name, "and my ticket price is", ticketPrice, "dollars")

	for {
		var surName string
		var entryPrice uint
		fmt.Println("Please enter your surName: ")
		fmt.Scan(&surName)
		fmt.Println("Please enter your entryPrice: ")
		fmt.Scan(&entryPrice)
		fmt.Println("Thank you for coming to the cinema", surName, "your entry price is", entryPrice, "dollars")
	}
}
