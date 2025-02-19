//Ensure the result is always true in all operations

package main

import (
	"fmt"
)

func main() {
	a := 100
	b := 200

	fmt.Println(b > a)
	fmt.Println(!(a > b || b == a))
	fmt.Println((a != b) && (b > a))
	fmt.Println((b == a) || (a < b))
	fmt.Println(!(a > b))
	fmt.Println((a > b) || (b > a))
	fmt.Println(b > a)
}
