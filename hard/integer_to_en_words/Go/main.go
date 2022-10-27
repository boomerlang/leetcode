package main

import (
	"fmt"
)

func main () {
	inputs := [...]int{
			2147483647,
			12345,
			1234567,
			8519348,
			85193106,
			222,
			1000000,
			10000000,
			10001001,
			1000000000,
			1000100001,
			1,
			0,
	}

	for _, num := range inputs {
		rez := numberToWords(num);
		fmt.Println(num, ":", rez)
	}
}