package main

import (
	"fmt"
	"sort"
)

//https://leetcode-cn.com/problems/majority-element/

func majorityElement1(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

/*
算法描述
摩尔投票法（Boyer–Moore majority vote algorithm），也被称作「多数投票法」，算法解决的问题是：如何在任意多的候选人中（选票无序），选出获得票数最多的那个。

算法可以分为两个阶段：

对抗阶段：分属两个候选人的票数进行两两对抗抵消
计数阶段：计算对抗结果中最后留下的候选人票数是否有效

只适用于票数过半的情况！！！！！！！！！！！！！！！！
*/

func majorityElement(nums []int) int {
	major := 0
	count := 0
	for index, num := range nums {
		if index == 0 || count == 0 {
			major = num
			count = 1
			continue
		}
		if num == major {
			count++
		} else {
			count--
		}

	}
	return major
}

func main() {
	fmt.Println(majorityElement([]int{1, 54, 23, 62, 13, 1, 1, 1, 1, 5}))
}
