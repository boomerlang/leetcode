package main

// Cathegory: Medium
// Success
// Details
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Permutations.
// Memory Usage: 2.6 MB, less than 70.74% of Go online submissions for Permutations.

// Accepted 1,404,570
// Submissions 1,879,152

// 26 / 26 test cases passed.
// 	Status: Accepted
// Runtime: 0 ms
// Memory Usage: 2.6 MB

// Date: 4.11.2022

// Copyright 2022 Bogdan Peta, All Rights Reserved

import (
	// "fmt"
)

// Generating permutation using Heap Algorithm
// Source:
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func generatePermutation(a *[]int, size int, n int, out *[][]int) {
    if (size == 1) {
			temp := make([]int, n)
			copy(temp, *a)
			*out = append(*out, temp)
    }
	
    for i := 0; i < size; i++ {
        generatePermutation(a, size - 1, n, out)
		
		if i < size - 1 {
			if size % 2 == 1 {
				temp := (*a)[0]
				(*a)[0] = (*a)[size - 1]
				(*a)[size - 1] = temp
				
			} else {
				temp := (*a)[i]
				(*a)[i] = (*a)[size - 1]
				(*a)[size - 1] = temp
				
			}	
		}
    }
}

func permute(nums []int) [][]int {
	out := [][]int{}
	n := len(nums)

	a := make([]int, n)
	copy(a, nums)

	generatePermutation(&a, n, n, &out)

	return out
}
