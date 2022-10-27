package main

// Cathegory: Medium
// Success
// Details
// Runtime: 12 ms, faster than 72.86% of Go online submissions for Integer to Roman.
// Memory Usage: 3.3 MB, less than 70.55% of Go online submissions for Integer to Roman.

// 3999 / 3999 test cases passed.
// 	Status: Accepted
// Runtime: 12 ms
// Memory Usage: 3.3 MB

// copyright 2022 Bogdan Peta 

import (
	// "fmt"
	"strconv"
	"strings"
)

var pairs map[int]byte = map[int]byte{
		   1: 'I', 
		   5: 'V',
		  10: 'X',
		  50: 'L',
		 100: 'C',
		 500: 'D',
		1000: 'M',
	}

var subtr_pairs map[int]string = map[int]string{
		  4: "IV",   9: "IX", 
		 40: "XL",  90: "XC",
		400: "CD", 900: "CM",
}

// I, II, III,  IV, V, VI, VII, VIII, IX, X
// 1 <= u <= 9
func unit_to_letter(u int) string {
	if u < 4 {
		return strings.Repeat("I", u)
	} else if ( u == 4 || u == 9 ) {
		return subtr_pairs[u]
	} else if ( u == 5 ) {
		return string(pairs[u])
	} else if ( u == 6 || u == 7  || u == 8) {
		return "V"+strings.Repeat("I", u - 5)
	}
	return ""
}

// X XX XXX XL L LX LXX LXXX XC
// 10 <= u <= 90
func decimals_to_letter(u int) string {
	if u < 40 {
		return strings.Repeat("X", u / 10)
	} else if ( u == 40 || u == 90 ) {
		return subtr_pairs[u]
	} else if ( u == 50 ) {
		return string(pairs[u])
	} else if ( u == 60 || u == 70  || u == 80) {
		return "L"+strings.Repeat("X", u / 10 - 5)
	}
	return ""
}

// C CC CCC D DX DXX DXXX CM
// 100 <= u <= 900
func hundreds_to_letter(u int) string {
	if u < 400 {
		return strings.Repeat("C", u / 100)
	} else if ( u == 400 || u == 900 ) {
		return subtr_pairs[u]
	} else if ( u == 500 ) {
		return string(pairs[u])
	} else if ( u == 600 || u == 700  || u == 800) {
		return "D"+strings.Repeat("C", u / 100 - 5)
	}
	return ""
}

// M MM MMM 
// 1000 <= u <= 3000
func thousands_to_letter(u int) string {
	if u <= 3000 {
		return strings.Repeat("M", u / 1000)
	} 
	return ""
}


func intToRoman(num int) string {

	units := map[int]int{
		4: 1000,
		3: 100,
		2: 10,
		1: 1,
	}
	
	num_roman := ""
	num_str := strconv.Itoa(num)
	len_num_str := len(num_str)

	for i := 0; i < len_num_str; {
		val, _ := strconv.Atoi(string(num_str[i]))
		
		unit := units[len_num_str - i]

		num0 := val * unit
		if unit == 1000 {
			num_roman += thousands_to_letter(num0)
		} else if unit == 100 {
			num_roman +=hundreds_to_letter(num0)
		} else if unit == 10 {
			num_roman +=decimals_to_letter(num0)
		} else if unit == 1 {
			num_roman +=unit_to_letter(num0)
		}

		i++
	}


    return num_roman
}

