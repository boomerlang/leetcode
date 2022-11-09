package main

import (
	"fmt"
)

func main () {
	inputs := [...][]string{
		[]string{"eat","tea","tan","ate","nat","bat"},
		[]string{""},
		[]string{"a"},
		[]string{"", ""},
		[]string{"c", "c"},
		[]string{"h", "h", "h"},
		[]string{"bdddddddddd","bbbbbbbbbbc"},
		[]string{"","b",""},
		[]string{"tea","","eat","","tea",""},
	}

	for _, s := range inputs {
		rez := groupAnagrams(s)
		fmt.Println(s, ":", rez)
	}
}
