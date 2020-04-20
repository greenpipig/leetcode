package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		tmp := 0
		for j := i; j < len(nums); j++ {
			tmp = tmp + nums[j]
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}

//最优解，时间复杂度O(n),空间复杂度O(1)
func maxSubArrayNew(nums []int) int {
	l := len(nums)
	max := nums[0]

	for i := 1; i < l; i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{1, 2, 3, 4, 5}))
}
