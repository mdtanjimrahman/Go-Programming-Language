package main

import "fmt"

var (
	day   string = "on Monday"
	date  int    = 15
	month string = "February"
	year  int    = 2025

	myName string = "Tanjim"
	myAge  int    = 24
)

func main() {
	name, language, lessons := "Tanjim ", "Learning GoLang ", 101
	fmt.Print(name, language, lessons, "\n \n") //have to mannually put spaces when declaring for readable output

	fmt.Println(name, language, day, date, month, year) //space in between variable comes auto for readable output

	title := "GoLang"
	title, myName, myAge := "Learning GoLang ", "with Tanjim ", 25 //reassigning values (have to be same data type)
	fmt.Print("\n", title, myName, myAge, "\n")

	one, two := "Python ", "GoLang "
	one, two = two, one //swapping values
	fmt.Print("\n", one, two, "\n")

}
