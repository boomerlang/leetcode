package main

// Cathegory: Easy
// Success
// Runtime: 8 ms, faster than 83.54% of Go online submissions for Roman to Integer.
// Memory Usage: 2.9 MB, less than 30.47% of Go online submissions for Roman to Integer.
// Date; 21.10.2022

// 3999 / 3999 test cases passed.
// 	Status: Accepted
// Runtime: 8 ms
// Memory Usage: 2.9 MB

// copyright 2022 Bogdan Peta 

import (
	// "fmt"
)

func romanToInt(s string) int {

	pairs := map[byte]int{
		'I': 1, 
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	subtr_pairs := map[string]int{
		"IV": 4,   "IX": 9, 
		"XL": 40,  "XC": 90,
		"CD": 400, "CM": 900,
	}

	sum := 0
	
	j := 0
	if len(s) == 1 {
		sum = pairs[s[0]]
	} else {
		for i := 1; i < len(s); {
			var next byte = 'Z'

			if i < len(s) - 1  {
				next = s[i + 1]
			}
			
			crt := s[i]
			prev := s[i - 1]
			
			group1 := string(prev) + string(crt)

			group2 := string(crt) + string(next)
			
			if val2, ok2 := subtr_pairs[group2]; ok2 {
				sum += pairs[prev] + val2
				i = i + 3
				j = i
				
				continue
			} else {
				if val, ok := subtr_pairs[group1]; ok {
					sum += val
				} else {
					sum += pairs[prev] + pairs[crt] 
				}
			}
			i = i + 2
			j = i
		}
		
		if j < len(s) + 1 {
			sum += pairs[s[j-1]]
		}
	}
    return sum
}

