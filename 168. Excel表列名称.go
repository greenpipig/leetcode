package main

import "fmt"

//https://leetcode-cn.com/problems/excel-sheet-column-title/
// 本题考查了对与acsii编码的知识

func convertToTitle(n int) string {
	strNum := ""
	for n != 0 {
		strNum = string((n-1)%26+65) + strNum
		n = (n - 1) / 26
	}
	return strNum
}

func main() {
	fmt.Println(convertToTitle(701))
}
