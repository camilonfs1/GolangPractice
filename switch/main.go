package main

import "fmt"

func main() {
	ans := 2

	switch ans {
	case 1:
		fmt.Println("hola1")
	case 2, 3:
		fmt.Println("hola 2 or three")
	default:
		fmt.Println("hola default")
	}

	switch {
	case ans > 20:
		fmt.Println("hola 20 or more")
	default:
		fmt.Println("hola default")
	}
}
