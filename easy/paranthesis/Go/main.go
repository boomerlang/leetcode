package main

import (
	"fmt"
)

func main () {
	inputs := [...]string{
		"([)]",
		"{[]}",
		"()[]{}",
		"()[({})]{[({})]}",
	}

	for _, test := range inputs {
		rez := validParant(test)
		fmt.Println(rez)
	}
}