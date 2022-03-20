package main

import "fmt"

func test5() {
	fmt.Println("hello")
}

func main() {
	x := test5
	x()

	test := func() {
		fmt.Println("hhello 2")
	}

	test()
}
