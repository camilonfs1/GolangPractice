package main

import "fmt"

/* (& *)
*	& = pinter
*	* = references
 */
func main() {
	x := 7
	fmt.Println(&x) //Get location for 7 store, memory location
	fmt.Println(x)  //Get value

	y := &x
	fmt.Println(y) //Get location for 7 store, memory location

	*y = 9 //Modify value on 7 memory location
	fmt.Println(x, y)

}
