package main

import (
	"fmt"
)

func main () {
	inputs1 := [...][]string{
		 []string{"5","2","C","D","+"},

		 []string{"5","-2","4","C","D","9","+","+"},

		 []string{"1", "C"},

		 []string{"1","D","D","D"},

	}
	
	for _, test := range inputs1 {
		sum := calcPoints(test)

		fmt.Println(sum)
	}
}