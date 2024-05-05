package main

import (
	"fmt"
)

func main() {
	fmt.Println("Name Concatenation!")
	concatenate()
}

func concatenate() {
	var FirstName string
	fmt.Println("Enter your first name: ")
	fmt.Scan(&FirstName)

	var LastName string
	fmt.Println("Enter your last name: ")
	fmt.Scan(&LastName)

	var Age int
	fmt.Println("Enter your age: ")
	fmt.Scan(&Age)

	CurrentYear := 2024
	YearOfBirth := CurrentYear - Age

	var Gender string
	fmt.Println("Enter your gender: ")
	fmt.Scan(&Gender)

	if Gender == "male" || Gender == "Male" || Gender == "MALE" {
		fmt.Printf("Welcome %s, son of %s, your year of birth is %d\n", FirstName, LastName, YearOfBirth)
	} else if Gender == "female" || Gender == "Female" || Gender == "FEMALE" {
		fmt.Printf("Welcome %s, daugther of %s, your year of birth is %d\n", FirstName, LastName, YearOfBirth)
	}

	fmt.Printf("Welcome, %s %s, you are %d years old!\n", FirstName, LastName, Age)
	fmt.Printf("Welcome, %s %s, your year of birth is %d", FirstName, LastName, YearOfBirth)
}
