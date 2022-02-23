package main

import "fmt"

func main() {

	x := 5
	y := 6
	val := x == y
	fmt.Printf("%t", val)

	val = x < y
	fmt.Printf("\n%t", val)

	val = x <= y
	fmt.Printf("\n%t", val)

	val = x >= y
	fmt.Printf("\n%t", val)

	a := "camilo"
	b := "camilo1"

	val = a == b
	fmt.Printf("\n%t", val)
	// (<, >, <=, >=, ==, !=)
}
