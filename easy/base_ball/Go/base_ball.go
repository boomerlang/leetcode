package main

// Cathegory: Easy
// Success
// Details
// Runtime: 6 ms, faster than 31.51% of Go online submissions for Baseball Game.
// Memory Usage: 2.8 MB, less than 15.07% of Go online submissions for Baseball Game.

// 39 / 39 test cases passed.
// 	Status: Accepted
// Runtime: 6 ms
// Memory Usage: 2.8 MB

// copyright 2022 Bogdan Peta 

import (
	// "fmt"
	"strconv"
)

func calcPoints(input []string) int {

	var out []int = []int{}

	for _, ch := range input {
		val, err := strconv.Atoi(string(ch))

		if err == nil {
			out = append(out, val)
		} else {
			if ch == "D" {
				crt := out[len(out) - 1]
				out = out[0:len(out) - 1]
				out = append(out, crt, crt * 2)
			}
			if ch == "C" {
				out = out[0:len(out) - 1]
			}
			if ch == "+" {
				x := out[len(out)-2:len(out)]
				out = out[0:len(out) - 2]
				out = append(out, x[0], x[1], x[0] + x[1])
			}
		}
	}
	
	sum := 0
	for _, elem  := range out {
		sum += elem
	}

	return sum
}
