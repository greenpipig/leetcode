package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	str := "1"
	for i := 1; i < n; i++ {
		strnew := ""
		counter := 1
		for j := 0; j < len(str)-1; j++ {
			//检索相同的字符长度
			if str[j] == str[j+1] {
				counter++
				continue
			}
			strnew = strnew + strconv.Itoa(counter) + string(str[j])
			counter = 1
		}
		//加上末尾的字符
		str = strnew + strconv.Itoa(counter) + string(str[len(str)-1])
	}
	return str
}

func main() {
	fmt.Println(countAndSay(6))
}
