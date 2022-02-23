package main

import "fmt"

func main() {
	age := 15

	for age < 18 {
		age++
		fmt.Printf("hello")
	}

	for x := 0; x <= 5; x++ {
		fmt.Println(x)
	}

	fmt.Printf("---------------------\n")

	for x := 0; x <= 1000000; x += 2 {
		if x == 6 {
			break
		} else {
			fmt.Println(x)
			continue
		}
	}
}
