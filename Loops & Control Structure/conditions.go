package main

import (
	"fmt"
)

var (
	currYear       int = 2025
	year           int
	adName, adPass string = "admin", "aiubAdmin"
	uname, upass   string = "student", "aiub"
	name, pass     string
)

func main() {
	fmt.Println("\n*******Since when did you started programming?********")
	fmt.Print("Enter your starting year: ")
	fmt.Scanln(&year)

	exp := (currYear - year)

	if exp >= 4 {
		fmt.Print("\n\nYou are eligible to login as Admin\n\n")
		fmt.Print("Enter username: ")
		fmt.Scanln(&name)
		fmt.Print("Enter password: ")
		fmt.Scanln(&pass)

		if name == adName {
			if pass == adPass {
				fmt.Print("\nAuthetication Successful\n\n")
			} else {
				fmt.Print("\nWrong Password\n\n")
			}
		} else {
			fmt.Print("\nWrong Username\n\n")
		}
	} else if (exp < 4) || (exp >= 1) {
		fmt.Print("\n\nYou are eligible to login as Intern\n\n")
		fmt.Print("Enter username: ")
		fmt.Scanln(&name)
		fmt.Print("Enter password: ")
		fmt.Scanln(&pass)

		if name == uname {
			if pass == upass {
				fmt.Print("\nAuthetication Successful\n\n")
			} else {
				fmt.Print("\nWrong Password\n\n")
			}
		} else {
			fmt.Print("\nWrong Username\n\n")
		}
	} else {
		fmt.Print("\n\nYou are eligible to login. Better luck next year!\n\n")
	}
}
