package main

import (
	"fmt"
)

func main() {

	// Write a program that takes an array of elements and
	// prints the result of multiplying each element by 2.
	numberList := [...]int{23, 34, 66, 78, 200, 97, 89}

	fmt.Println()
	for i := 0; i < len(numberList); i++ {
		fmt.Print(numberList[i]*2, "\n")
	}
	fmt.Println()

	// Given two fixed arrays. The first one is an array of strings, and the second is a list of scores.
	// You need to use the following code to print this to the screen:

	// Your score is 90. Genius
	// Your score is 80. Good Job!
	// Your score is 50. Try Next time
	// (Can not change array order in declaration)
	str := [3]string{"Try Next time", "Good Job!", "Genius"}
	scores := [3]int{90, 80, 50}

	for i := 0; i < len(scores); i++ {
		var message string

		if scores[i] == 90 {
			message = str[2]
		} else if scores[i] == 80 {
			message = str[1]
		} else if scores[i] == 50 {
			message = str[0]
		}
		fmt.Printf("Your score is %d. %s\n", scores[i], message)
	}
	fmt.Println()

	// Write a program that takes exam scores as an array of elements and calculates the average
	// The array given in the code contains exam scores. Using a loop,
	// try to calculate the average and print it
	examscores := [5]int{24, 56, 78, 98, 67}
	sum := 0

	for i := 0; i < len(examscores); i++ {
		sum += examscores[i]
	}
	fmt.Print("\n", sum/len(examscores), "\n\n")
}
