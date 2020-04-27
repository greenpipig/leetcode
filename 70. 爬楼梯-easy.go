package main

import "fmt"

//暴力递归，想象每次可以走一步或者走两步
func climb1(i int, n int) int {
	if i == n {
		return 1
	}
	if i > n {
		return 0
	}
	return climb1(i+1, n) + climb1(i+2, n)
}

//通过写一个map读取内存减少递归次数，记忆化递归
func climb2(i int, n int, mem map[int]int) int {
	if i == n {
		return 1
	}
	if i > n {
		return 0
	}
	if mem[i] > 0 {
		return mem[i]
	}
	mem[i] = climb2(i+1, n, mem) + climb2(i+2, n, mem)
	return mem[i]
}

//第三种通过斐波那契数列，寻找规律
func climbStairs(n int) int {
	//return climb1(0, n)

	//mem := make(map[int]int, n)
	//return climb2(0, n, mem)

	memList := make([]int, n)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	memList[0] = 1
	memList[1] = 2
	for i := 3; i <= n; i++ {
		memList[i-1] = memList[i-2] + memList[i-3]
	}
	return memList[n-1]
}

func main() {
	fmt.Println(climbStairs(4))
}
