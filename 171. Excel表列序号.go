package main

import (
	"fmt"
	"math"
)

//https://leetcode-cn.com/problems/excel-sheet-column-number/

func titleToNumber(s string) int {
	num:=0
	for index,s1:=range s{
		num+=int(math.Pow(26,float64(len(s)-index-1)))*(int(s1)-64)
		//ret = 26*ret + (int(c-'A')+1)
	}
	return num
}
func main()  {
	fmt.Println(titleToNumber("AB"))
}
