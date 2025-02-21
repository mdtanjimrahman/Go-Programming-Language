package main

import (
	"fmt"
)

var (
	numRow    int
	numColumn int
	iteration int
	height    int
	a         int
	b         int
	option    int
)

func multiplication() {
	fmt.Print("\nEnter number multiplication will start from: ")
	fmt.Scanln(&a)
	fmt.Print("Enter number multiplication will start to perform from: ")
	fmt.Scanln(&b)
	fmt.Print("Enter number of iteration per number to perform: ")
	fmt.Scanln(&iteration)

	for numRow := a; numRow <= iteration; numRow++ {
		for numColumn := b; numColumn <= iteration; numColumn++ {
			result := numRow * numColumn
			fmt.Printf("%d x %d = %d\t", numRow, numColumn, result)
		}
	}
}

func asterisks() {
	fmt.Print("\nEnter the row height of the trinangle: ")
	fmt.Scanln(&height)

	for a := 1; a <= height; a++ {
		for b := 1; b <= a; b++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}
}

func main() {
	fmt.Print("\nChoose a number (1 ~ Multiplications, 2 ~ Draw Asterisks Triangle): ")
	fmt.Scanln(&option)

	switch option {
	case 1:
		multiplication()
	case 2:
		asterisks()
	}
}
