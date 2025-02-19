package main

import (
	"fmt"
)

var (
	a, b int     = 10, 21
	x, y float32 = 5, 21

	person1, person2, person3 int //can't be given any negative input to store
)

func basic_operation() {
	fmt.Print("Sum: ", a+b, "\n")
	fmt.Print("Subtract: ", b-a, "\n")
	fmt.Print("Multiply: ", a*b, "\n")
	fmt.Print("Division: ", y/x, "\n")
	fmt.Print("Remainder: ", b%a, "\n\n")
}

func avgWeight() {
	fmt.Print("\n**********Enter weights (kg)*********\n")
	fmt.Print("1st person's weight: ")
	fmt.Scanln(&person1)

	fmt.Print("2nd person's weight: ")
	fmt.Scanln(&person2)

	fmt.Print("3rd person's weight: ")
	fmt.Scanln(&person3)

	avg := (person1 + person2 + person3) / 3
	fmt.Print("The average weight is: ", avg, "\n\n")
}

func main() {
	basic_operation()
	avgWeight()
}
