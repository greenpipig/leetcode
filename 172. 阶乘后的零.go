package main

import "fmt"

//https://leetcode-cn.com/problems/factorial-trailing-zeroes/

func trailingZeroes(n int) int {
	c:=0
	for n>0{
		c=c+n/5
		n=n/5
	}
	return c
}

func main(){
	fmt.Println(trailingZeroes(125))
}
