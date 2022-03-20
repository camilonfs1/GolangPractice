package main

import "fmt"

func main() {
	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var s []int = x[1:3]

	fmt.Println("s", s)
	fmt.Println("length", len(s))
	fmt.Println("capacity", cap(s))

	s = x[:cap(s)]

	fmt.Println("s", s)
	fmt.Println("length", len(s))
	fmt.Println("capacity", cap(s))

}
