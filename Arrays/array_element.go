package main

import (
	"fmt"
)

var (
	option    int
	usertext  string
	userindex int
	userval   int
	lang      = [5]string{"Go", "Python", "Rust", "Java", "C++"}
	arr       = [...]string{"Go", "Python", "R", "C++", "Java", "C#"} //No fixed array size
)

func arrAccess() {
	//Access an array element
	fmt.Print("\nEnter array index to choose a programming language: ")
	fmt.Scanln(&userval)

	fmt.Print("\n", lang[userval], "\n\n") //without any input by default it access 0 or first index and return value
}

func arrModify() {
	//Modify an array
	fmt.Print("\nEnter a name of programming language: ")
	fmt.Scanln(&usertext)
	fmt.Print("Enter an index: ")
	fmt.Scanln(&userindex)

	if userindex > len(arr) {
		fmt.Print("\nArray has only ", len(arr), " indexes\n")
	} else {
		arr[userindex] = usertext
		fmt.Print(arr)
	}
}

func main() {
	fmt.Print("\nChoose a number (1 ~ Access arr, 2 ~ Modify arr): ")
	fmt.Scanln(&option)

	switch option {
	case 1:
		arrAccess()
	case 2:
		arrModify()
	}
}
