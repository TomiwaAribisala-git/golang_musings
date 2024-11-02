package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	readFile("./line.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	contents := string(databyte)
	fmt.Println(contents)

	newfile, err := os.Create("./log.txt")
	if err != nil {
		fmt.Printf("error creating new file: %v", err)
	}

	defer newfile.Close()
	io.WriteString(newfile, contents)
}
