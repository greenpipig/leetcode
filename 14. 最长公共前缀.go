package main

//最长公共前缀14
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	same := ""
	samestr := ""
	min := strs[0]
	for _, str := range strs {
		if len(min) > len(str) {
			min = str
		}
	}
	for index, ch := range min {
		same += string(ch)
		for _, str := range strs {
			strjudge := ""
			for i := 0; i <= index; i++ {
				strjudge += string(str[i])
			}
			if strjudge != same {
				return samestr
			}
		}
		samestr += string(ch)
	}
	return samestr
}
