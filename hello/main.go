package main

import "fmt"

var (
	day   string = "on Monday"
	date  int    = 15
	month string = "February"
	year  int    = 2025
)

func main() {
	name, language, lessons := "Tanjim", "Learning GoLang", 101
	fmt.Println(name, language, lessons)

	fmt.Println(name, language, day, date, month, year)
}
