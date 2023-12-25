package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"operations.go"
)

func main() {

	// log is used to print logs to the console
	log.SetPrefix("Operations: ")
	log.SetFlags(0)

	// Println helps us print to the console
	fmt.Println("This application gets lets you find age and gender by name!")
	fmt.Println("Please enter your name here:")

	// The following lines are used to read input from console
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("There was an issue reading the input")
		return
	}

	// Here we use a function from an external module
	ageString, err := operations.GetAgeUsingName(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(ageString)

	genderString, err := operations.GetGenderUsingName(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(genderString)
}
