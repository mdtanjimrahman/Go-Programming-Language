package main

import (
	"fmt"
)

func main() {
	var first, second int
	fmt.Print("Enter 1st person's age: ")
	fmt.Scanln(&first)

	fmt.Print("Enter 2nd person's age: ")
	fmt.Scanln(&second)

	// Comparison operations and print the results
	equal := (first == second)
	greater := (first > second)
	lesser := (first < second)
	greaterEqual := (first >= second)
	lesserEqual := (first <= second)
	notEqual := (first != second)

	fmt.Println("== :", equal)
	fmt.Println("> :", greater)
	fmt.Println("< :", lesser)
	fmt.Println("<= :", lesserEqual)
	fmt.Println(">= :", greaterEqual)
	fmt.Println("!= :", notEqual)
}
