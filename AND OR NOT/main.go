package main

import "fmt"

func main() {
	// AND &&
	// OR ||
	// NOT !

	x := 8
	val := (true || false) && !false || x > 7 || true

	fmt.Printf("%t", val)
}
