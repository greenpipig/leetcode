package main

import (
	"math"
	"strconv"
)

//整数反转第7题,值得学习的是平方运算
func reverse(x int) int {
	judge:=false
	if x<0{
		x=-x
		judge=true
	}
	stra:=strconv.Itoa(x)
	revse:= func(a string)int {
		strrevse:=""
		for i:=0;i<len(a);i++{
			strrevse+=string(a[len(a)-1-i])
		}
		inta,_:=strconv.Atoi(strrevse)
		return inta
	}(stra)
	if judge{
		revse=-revse
	}
	if revse < math.MinInt32 || revse > math.MaxInt32 {//todo 学习平方运算
		return 0
	}
	return revse
}
func reverse_NB(x int) int {
	y := 0
	for x!=0 {
		y = y*10 + x%10//精华，爆赞
		if !( -(1<<31) <= y && y <= (1<<31)-1) {
			return 0
		}
		x /= 10
	}
	return y
}

