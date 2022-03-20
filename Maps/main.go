package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"camilo": 1,
		"sergio": 2,
	}

	fmt.Println(mp["camilo"])

	mp["camilo"] = 10

	fmt.Println(mp["camilo"])

	delete(mp, "sergio")

	fmt.Println(mp)

	//check is key exist
	val, ok := mp["camilo"]

	if ok {
		fmt.Println(val)
	}
}
