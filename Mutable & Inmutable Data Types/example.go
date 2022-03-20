package main

import "fmt"

func changeFist(slice []int) {
	slice[0] = 100
}
func main() {
	var x []int = []int{3, 4, 5}
	fmt.Println(x)
	changeFist(x)
	fmt.Println(x)
}
