package main

import "fmt"

func main() {
	// Mutable data
	var x []int = []int{3, 4, 5} //Slide
	y := x
	y[0] = 100
	fmt.Println(x, y)

	var x1 map[string]int = map[string]int{"hello": 5} //map
	y1 := x1
	y1["y"] = 100
	fmt.Println(x1, y1)

	//Inmutable

	var a [2]int = [2]int{3, 4}
	b := a
	b[0] = 100
	fmt.Println(a, b)

}
