package main

func searchInsert(nums []int, target int) int {
	for index, num := range nums {
		if num >= target {
			return index
		}
	}
	return len(nums)
}

func main() {
	searchInsert([]int{1, 2, 3, 4, 5}, 3)
}
