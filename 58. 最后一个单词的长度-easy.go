package main

import "fmt"

func lengthOfLastWord1(s string) int {
	if len(s)==1&&string(s[0])==" "||len(s)==0{
		return 0
	}
	pin:=0
	judge:=false
	judge1:=false
	lenght:=len(s)-1
	if string(s[len(s)-1])==" "{
		judge1=true
		lenght=lenght-1
	}
	for i:=lenght;i>=0;i--{
		if string(s[i])==" "{
			pin=i
			judge=true
			break
		}
	}
	if judge1&&pin==0{
		return lenght-pin+1
	}
	if pin==0&&!judge&&!judge1{
		return len(s)
	}
	return lenght-pin
}

func lengthOfLastWord(s string) int {
	len1:=0
	for i:=len(s)-1;i>=0;i--{
			if string(s[i]) != " " {
				len1 += 1
			}else {
				if len1!=0{
					break
				}
				len1=0
			}
		}
	return len1
}


func main()  {
	fmt.Println(lengthOfLastWord(" 9 "))
}