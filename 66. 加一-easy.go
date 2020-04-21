package main

import "fmt"

func plusOne(digits []int) []int {
	if digits[len(digits)-1]!=9{
		digits[len(digits)-1]+=1
		return digits
	}else {
		for i:=len(digits)-1;i>=0;i--{
			if digits[i]+1==10{
				digits[i]=0
				if digits[0]==0{
					return append([]int{1},digits...)
				}
			}else {
				digits[i]+=1
				break
			}
		}

	}
	return digits
}

func main()  {
	fmt.Println(plusOne([]int{9,9,9}))
}
