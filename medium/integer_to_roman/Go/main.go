package main

import (
	"fmt"
)

func main () {
	inputs := [...]int{
		3,
		58,
		1994,
	}

	for _, s := range inputs {
		rez := intToRoman(s)
		fmt.Println(s, ":", rez)
	}
}