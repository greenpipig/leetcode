package main

import "fmt"

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	a, b := 0, x //设置两个值，表示最大值和最小值
	for b-a > 1 {
		c := (a + b) / 2
		if x > c*c {
			a = c
		} else if x == c*c {
			return c
		} else {
			b = c
		}
	}
	return a
}
func main() {
	fmt.Println(mySqrt(10))
}
