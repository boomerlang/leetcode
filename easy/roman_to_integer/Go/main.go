package main

import (
	"fmt"
)

func main () {
	inputs := [...]string{
			"XLV",
			"III",
			"LVIII",
			"MCMXCI",
			"DXCVIII",
			"XLIX",
			"DLV",
			"LXXIX",
			"IV",
			"XXI",
			"MMCCCXCIX",
			"MMCDXXV",
			"MMCCC",
			"MMCDXXV",
			"MDCCCLXXXIV",
			"MDCXCV",
			"MCMXCIV",
			"II",
			"D","L","V", "X",
	}
	
	for _, s := range inputs {
		rez := romanToInt(s)
		fmt.Println(s, ":", rez)
	}
	
}