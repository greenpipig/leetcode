package main

import (
	"fmt"
)

func powerf2(x int, n int) int {
	if n == 0 {
		return 1
	} else {
		return x * powerf2(x, n-1)
	}
}

func addBinary(a string, b string) string {
	if a == "0" && b == "0" {
		return "0"
	}
	if len(b) > len(a) {
		a, b = b, a
	}
	str := ""
	for i := 0; i < len(a); i++ {
		if i+1 > len(b) {
			str = string(a[len(a)-1-i]) + str
			continue
		}
		if string(a[len(a)-1-i]) == "1" && string(b[len(b)-1-i]) == "1" {
			str = "2" + str
		} else if string(a[len(a)-1-i]) == "0" && string(b[len(b)-1-i]) == "0" {
			str = "0" + str
		} else {
			str = "1" + str
		}
	}
	add := 0
	str1 := ""
	for i := len(str) - 1; i >= 0; i-- {
		if string(str[i]) == "2" && add == 1 {
			str1 = "1" + str1
			add = 1
		} else if string(str[i]) == "1" && add == 1 {
			str1 = "0" + str1
			add = 1
		} else if string(str[i]) == "1" && add != 1 {
			str1 = "1" + str1
			add = 0
		} else if string(str[i]) == "2" && add != 1 {
			str1 = "0" + str1
			add = 1
		} else if string(str[i]) == "0" && add == 1 {
			str1 = "1" + str1
			add = 0
		} else {
			str1 = "0" + str1
		}

	}
	if add == 1 {
		str1 = "1" + str1
	}
	return str1
}

func main() {
	fmt.Println(addBinary("100", "110010"))
}
