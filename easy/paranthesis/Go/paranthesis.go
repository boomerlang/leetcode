package main

// Cathegory: Easy
// Success
// Details
// Runtime: 3 ms, faster than 35.99% of Go online submissions for Valid Parentheses.
// Memory Usage: 2 MB, less than 43.19% of Go online submissions for Valid Parentheses.

// 92 / 92 test cases passed.
// 	Status: Accepted
// Runtime: 3 ms
// Memory Usage: 2 MB

// copyright 2022 Bogdan Peta 

import (
	// "fmt"
)

func validParant(s string) bool {

	temp := []rune{}
	pairs := map[rune]rune{'(': ')', '[': ']', '{': '}'}

	for _, ch := range s {
		if len(temp) == 0 {
			temp = append(temp, pairs[ch])
		} else {
			last := temp[len(temp) - 1]
			if ch == last {
				temp = temp[0:len(temp) - 1]
			} else {
				temp = append(temp, pairs[ch])
			}
		}
	}
	sz := len(temp);

	return sz == 0
}
