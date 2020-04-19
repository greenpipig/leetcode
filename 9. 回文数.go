package main

import "strconv"

//回文数第九题
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x >= 0 && x < 10 {
		return true
	} else {
		stra := strconv.Itoa(x)
		for i := 0; i < len(stra)/2; i++ {
			if stra[i] != stra[len(stra)-1-i] {
				return false
			}
		}
		return true
	}
}
