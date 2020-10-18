package main

func twoSum(numbers []int, target int) []int {
	for i,num:=range numbers{
		if num>=target{
			break
		}
		for j:=i+1;j<len(numbers);j++{
			if num+numbers[j]==target{
				return []int{i+1,i+2}
			}
		}
	}
	return []int{1}
}

//双指针法
func twoSum1(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right{
		if numbers[left] + numbers[right] ==  target{
			return []int{left+1, right+1}
		}else if numbers[left] + numbers[right] > target{
			right--
		}else{
			left++
		}
	}
	return []int{}
}
