package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	if needle == "" || haystack == needle {
		return 0
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		str := ""
		for j := 0; j < len(needle); j++ {
			str = str + string(haystack[j+i])
		}
		if str == needle {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hell", "ll"))
}
