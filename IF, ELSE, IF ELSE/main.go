package main

import "fmt"

func main() {
	age := 15

	if age >= 18 {
		fmt.Printf("hello")
	} else if age < 10 {
		fmt.Printf("hello2")
	} else {
		fmt.Printf("hello3")
	}

}
