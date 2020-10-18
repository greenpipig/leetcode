package main

import "fmt"

func maxProfit_slow(prices []int) int {
	maxPrice := 0
	for i, price := range prices {
		for j := i + 1; j < len(prices); j++ {
			nowPrice := prices[j] - price
			if nowPrice > maxPrice {
				maxPrice = nowPrice
			}
		}
	}
	return maxPrice
}

func maxProfitI(prices []int) int {
	if len(prices)==1||len(prices)==0{
		return 0
	}

	smallestNum:=prices[0]
	profile:=0
	for _,price:=range prices{
		if price-smallestNum>profile{
			profile=price-smallestNum
		}
		if price<smallestNum{
			smallestNum=price
		}
	}
	return profile
}


