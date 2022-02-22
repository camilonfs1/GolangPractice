package main

import "fmt"

func main() {
	fmt.Printf("%T \n", 10)

	fmt.Printf("Hello %T %v", 10, 10)

	fmt.Printf("\n %%", 0.10)

	fmt.Printf("\n Number : %9q", "Tim")

	var out string = fmt.Sprintf("Number: %07d is cool", 45)

	fmt.Println("\n", out)

}
