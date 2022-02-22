package main

import "fmt"

func main() {
	var name string
	name = "name example"
	name = "Tim"

	var number uint16 = 260

	fmt.Println("Hello world!", name, number)

	number = number + 5
	fmt.Println("Hello world!", name, number)

}
