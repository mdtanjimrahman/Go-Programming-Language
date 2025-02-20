package main

import (
	"fmt"
)

var num, numShift int

func main() {
	fmt.Print("********Enter Number and shifting value**********\n")
	fmt.Print("Enter you number: ")
	fmt.Scanln(&num)

	fmt.Print("Enter shifting value: ")
	fmt.Scanln(&numShift)

	fmt.Println("\n\nLeft Shift Result: ", (num << numShift))
	fmt.Println("Right Shift Result: ", (num >> numShift), "\n")
}
