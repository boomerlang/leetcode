package main

// Cathegory: Hard
// Success
// Details
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Integer to English Words.
// Memory Usage: 2.3 MB, less than 46.67% of Go online submissions for Integer to English Words.

// 601 / 601 test cases passed.
// 	Status: Accepted
// Runtime: 0 ms
// Memory Usage: 2.3 MB
	
// copyright 2022 Bogdan Peta 

import (
	"fmt"
	"strconv"
	"strings"
)

var pairs0 map[string]string = map[string]string{
		//    0: "zero",
		   "1": "One ",
		   "2": "Two ",
		   "3": "Three ", 
		   "4": "Four ", 
		   "5": "Five ",
		   "6": "Six " ,
		   "7": "Seven ",
		   "8": "Eight ",
		   "9": "Nine ", 
		  "10": "Ten ",
		  "11": "Eleven ",
		  "12": "Twelve ",
		  "13": "Thirteen ",
		  "14": "Fourteen ",
		  "15": "Fifteen ",
	      "16": "Sixteen ",
		  "17": "Seventeen ",
		  "18": "Eighteen ",
		  "19": "Nineteen ",
		  "20": "Twenty ",
		  "30": "Thirty ",
		  "40": "Forty ",
		  "50": "Fifty ",
		  "60": "Sixty ",
		  "70": "Seventy ",
		  "80": "Eighty ",
		  "90": "Ninety ",
	}


func make_hundreds(rank string) string {

	if  val0, ok := pairs0[rank]; ok  {
		return val0 
	}
	rank_len := len(rank)
	if rank_len == 3 {
		if rank[0] == '0' {
			if  val0, ok := pairs0[string(rank[1])+string(rank[2])]; ok  {
				return val0
			} else {
				return pairs0[string(rank[1])+"0"] + pairs0[string(rank[2])] 
			}
		} else {
			if  val0, ok := pairs0[string(rank[1])+string(rank[2])]; ok  {
				return pairs0[string(rank[0])] + "Hundred " + val0
			} else {
				return pairs0[string(rank[0])] + "Hundred " + pairs0[string(rank[1])+"0"] + pairs0[string(rank[2])] 
			}
		}
	} else if rank_len == 2 {
		if rank[0] == '0' {
			return pairs0[string(rank[1])]
		} else {
			if  val0, ok := pairs0[string(rank[0])+string(rank[1])]; ok  {
				return val0
			} else {
				return pairs0[string(rank[0])+"0"] + pairs0[string(rank[1])]
			}
		}
	}
	return ""
}

func numberToWords(num int) string {

	if num == 0 {
		return "Zero"
	}
	
	num_en := ""
	num_str := strconv.Itoa(num)
	len_num_str := len(num_str)

	num_rank := ""
	
	j := 0
	rank := ""
	for i := len_num_str - 1; i >= 0 ; {
		
		if i - 2 <= 0 {
			rank = num_str[0:i+1]
		} else {
			rank = num_str[i-2:i+1]
		}
		
		num_rank = make_hundreds(rank)
		if j == 0 {
			num_en = num_rank + num_en
		} else if j == 1 && strings.TrimSpace(rank) != "000" {
			num_en = num_rank + "Thousand "+ num_en
		} else if j == 2 && strings.TrimSpace(rank) != "000" {
			num_en = num_rank + "Million "+ num_en
		} else if j == 3 {
			num_en = num_rank + "Billion "+ num_en
		}
		
		i -= 3
		j += 1
	}

    return strings.TrimSpace(num_en)
}
