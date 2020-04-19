package main

func removeDuplicates(nums []int) int {
	if len(nums)<2{
		return len(nums)
	}
	slow:=0
	for {
		if len(nums)==slow+1{
			break
		}
		if nums[slow]==nums[slow+1]{
			nums=append(nums[:slow+1],nums[slow+2:]...)
		}else {
			slow++
		}
	}
	return len(nums)
}

//快慢指针法
func removeDuplicatesNIU(nums []int) int {
	i := 0
	for j := 1 ; j < len(nums) ; j++{
		if nums[j] != nums[i]{
			i += 1
			nums[i] = nums[j]
		}
	}
	return i+1
}


func main()  {
	removeDuplicates([]int{0,0,0,0,0,1,1,1,2})
}