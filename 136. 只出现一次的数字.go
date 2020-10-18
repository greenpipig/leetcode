package main

/*主要学习位运算
任意一个数和0异或仍然为自己：
a⊕0 = a

任意一个数和自己异或是0：

a⊕a=0

异或操作满足交换律和结合律：

a⊕b⊕a=(a⊕a)⊕b=0⊕b=b  */


func singleNumber(nums []int) int {
	var res int
	for _, n := range nums{
		res = n ^ res
	}
	return res
}
