package main

import "fmt"

//https://leetcode-cn.com/problems/house-robber/

func get_max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return get_max(nums[0], nums[1])
	}
	nums[1]= get_max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		nums[i]=get_max(nums[i-1],nums[i]+nums[i-2])
		fmt.Println(nums[i])
	}
	return nums[len(nums)-1]
}

func main()  {
	fmt.Println(rob([]int{1,2,3,1}))
}