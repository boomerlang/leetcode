package main

import (
	"fmt"
)

func main () {
	inputs := [...][]int{
		[]int{1,2,3},
		[]int{0,1},
		[]int{1},

		[]int{1,2,3,4},

		[]int{1,2,3,4,5},

		[]int{1,2,3,4,5,6},

		[]int{1,2,3,4,5,6,7},

		// []int{1,2,3,4,5,6,7,8},
	}

	for _, a := range inputs {
		out := permute(a)
		fmt.Println("\n-----------------------------\n", a, ":", out, len(out))
	}

	}
