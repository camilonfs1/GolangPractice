package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Type the year you were born: ")
	scanner.Scan()
	year, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("Type your name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Printf("Hello world %d %q", year, name)
}
