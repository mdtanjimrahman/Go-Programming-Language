package main

import (
	"fmt"
)

var (
	monthlySalary int
	score         int
	option        int
)

func salary() {
	fmt.Print("\nEnter your monthly salary: ")
	fmt.Scanln(&monthlySalary)

	//Golang doesn't use break statements in switch blocks.
	//It automatically breaks after the matching case is executed.

	switch monthlySalary * 12 {
	case 36000:
		fmt.Print("\nYou are a junior developer\n")
	case 48000:
		fmt.Print("\nYou are on your way to becoming a senior developer\n")
	case 60000:
		fmt.Print("\nCongratulations! You are a senior developer\n")
	default:
		fmt.Print("\nYour annual salary: 35000\n")
	}
}

func scoring() {
	fmt.Print("\nEnter your exam score: ")
	fmt.Scanln(&score)

	switch {
	case score > 90:
		fmt.Print("\nCongratulations, you are a genius!\n")
	case score > 70 && score < 90:
		fmt.Print("\nGood Job\n")
	case score > 50 && score < 70:
		fmt.Print("\nDo better next time\n")
	case score < 50:
		fmt.Print("\nTry next time\n")
	}
}

func main() {
	fmt.Print("\nChoose a number (1 ~ Salary, 2 ~ Score): ")
	fmt.Scanln(&option)

	switch option {
	case 1:
		salary()
	case 2:
		scoring()
	}
}
