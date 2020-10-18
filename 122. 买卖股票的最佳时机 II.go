package main

import "fmt"

func maxProfit(prices []int) int {
	maxPro:=0
	for i:=1;i<len(prices);i++{
		if prices[i-1]<prices[i]{
			maxPro+=prices[i]-prices[i-1]
		}
		if prices[i-1]>prices[i]{
			continue
		}
	}
	return maxPro
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5,6,2,1,7}))
}
