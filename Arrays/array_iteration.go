package main

import (
	"fmt"
)

func main() {
	numberList := [...]int{23, 34, 66, 78, 200, 97, 89}

	for i := 0; i < len(numberList); i++ {
		fmt.Print(numberList[i]*2, "\n")
	}

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
}
