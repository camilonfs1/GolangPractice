package main

import "fmt"

func test(x int) {
	fmt.Printf("test", x)
}

func test2(x int) int {
	return x
}

func test3(x int) (int, int) {
	return x, x
}

func test4(x int) (string, int) {
	return "val", x
}

func tes5(x int) (a string, b int) {
	a = "val"
	b = x
	return
}

func main() {
	test(10)
	ans1, ans2 := test4(10)
	fmt.Println(ans1, ans2)
}
