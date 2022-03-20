package main

import "fmt"

func changeValue(str *string) {
	*str = "change! 1"
}
func changeValue2(str string) {
	str = "change ! 2"
}
func main() {
	toChange := "hello"
	changeValue(&toChange)
	fmt.Println(toChange)

	changeValue2(toChange)
	fmt.Println(toChange)
}
