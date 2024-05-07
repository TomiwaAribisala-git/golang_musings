package main

import (
	"fmt"

	"github.com/martinlindhe/unit"
)

func main() {
	var option int
	option1 := "Celsius to Kelvin"
	option2 := "Fahrenheit to Kelvin"

	fmt.Printf("Welcome to temperature unit conversion!\nEnter 1 for %s or 2 for %s: ", option1, option2)
	fmt.Scan(&option)
	if option == 1 {
		celsiusTokelvin()
	} else if option == 2 {
		fahrenheitTokelvin()
	} else {
		fmt.Println("You inputted an invalid number, enter 1 or 2")
	}
}

func celsiusTokelvin() {
	var input float64
	fmt.Println("Enter a celsius value: ")
	fmt.Scan(&input)
	result := unit.FromCelsius(input)
	fmt.Printf("%v°C is %vK", input, result)
}

func fahrenheitTokelvin() {
	var input float64
	fmt.Println("Enter a farenheit value: ")
	fmt.Scan(&input)
	result := unit.FromFahrenheit(input)
	fmt.Printf("%v°C is %vK", input, result)
}
