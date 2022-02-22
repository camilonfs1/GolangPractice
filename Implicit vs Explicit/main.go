package main

import "fmt"

func main() {
	var name string = "name example" // Explicit
	var number = 260                 // Implicit

	fmt.Println("------>", name, number)
	fmt.Printf("%T", number) //To cast var type

	var number2 = 260.8888 // Implicit
	fmt.Println("\n------")
	fmt.Printf("%T", number2) //To cast var type

	number3 := 6      //declaration + assignment
	var number4 = 260 //assignment

	fmt.Println("\n------", number4)
	fmt.Printf("%T", number3) //To cast var type
}
