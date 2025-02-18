package main

import (
	"fmt"
	"math"
)

const ( //constant data type external declaration
	name  string = "Tanjim"
	title string = "Learning GoLang"
	age   int    = 24
)

var ( //float variable declaration
	p32 float32 = math.Pi //Uses 4 bytes of memory (prints less decimal points)
	p64 float64 = math.Pi //Uses 8 bytes if memory (prints furthur decimal points)
)

var ( //integer variable declaration
	credit1 int    = 30   //values can be positive or negative
	credit2 int8   = -30  //values can be in range of (-128 to 127)
	credit3 uint   = 33   //values can only be in positive range
	credit4 uint8  = 33   //values can be in range of (0 to 255)
	credit5 int16  = -365 //values can be in range of (-32,768 to 32,767)
	credit6 uint16 = 365  //values can be in range of (0 to 65,535)
)

var ( //variable declaration for conversion
	toHex     uint8 = 125  //value to be converted to Hexadecimal
	toDecimal uint8 = 0xc8 //value to be converted to Decimal
)

// multi-line string holding variable using backquotes (key with the symbol "~" on it)
var htmlstring = `<html>
    <head>
        <title></title>
    </head>
</html>`

var file = `home\project\go`

func main() {
	const daysPerWeek, daysPerMonth, daysPerYear = 7, 30, 365     //constant data type short declaration
	fmt.Print("\n", daysPerWeek, daysPerMonth, daysPerYear, "\n") //printing const value
	fmt.Print(name, title, age, "\n")                             //printing const value

	fmt.Print(p32, "\n", p64, "\n") //printing float values

	fmt.Print("\n", credit6, "\n\n") //printing integer value

	fmt.Printf("%s %s %f %d %s", name, title, p32, credit6, "\n\n") //printing with format specifier

	fmt.Printf("%x %d %s", toHex, toDecimal, "\n\n") //format specifier %x converts decimal to hexadecimal and %d does vice-versa

	floatToint := int(p32) //type conversion
	fmt.Print(floatToint, "\n\n")

	fmt.Print(htmlstring, "\n\n") //printing multi-line string

	fmt.Print(name+" "+title, "\n\n") //Concatenation

	fmt.Print("Enter your path: ")
	fmt.Scan(&file) //escape scequence (for this to work: run the program in terminal)
	fmt.Println("Your file is located")
}
