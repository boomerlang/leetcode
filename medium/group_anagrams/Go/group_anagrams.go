package main

// Cathegory: Medium

// Success
// Details
// Runtime: 38 ms, faster than 66.98% of Go online submissions (1,723,930) for Group Anagrams.
// Memory Usage: 9.6 MB, less than 8.02% of Go online submissions (1,723,930) for Group Anagrams.

// 118 / 118 test cases passed.
// 	Status: Accepted
// Runtime: 38 ms
// Memory Usage: 9.6 MB

// Copyright 2022 Bogdan Peta, All Rights Reserved

import (
	// "fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {

	acc := map[string][]string{}

	for _, w := range strs {
		str_w := []rune(w)
		
		sort.Slice(str_w, func(i int, j int) bool { return str_w[i] < str_w[j] })
		
		w1 := string(str_w)

		acc[w1] = append(acc[w1], w)
	}

	rez := [][]string{}
	for _, v := range acc {
		if v != nil {
			rez = append(rez, v)
		}
	}

	return rez
}
