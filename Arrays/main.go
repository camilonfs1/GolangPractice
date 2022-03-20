package main

import "fmt"

func main() {
	arr := [5]int{1, 5, 3, 6, 5}

	fmt.Println("Hello", arr)

	arr2D := [5][5]int{
		{1, 5, 3, 6, 5},
		{1, 5, 3, 6, 5},
		{1, 5, 3, 6, 5},
		{1, 5, 3, 6, 5},
		{1, 5, 3, 6, 5},
	}

	fmt.Println("Hello", arr2D)

}
