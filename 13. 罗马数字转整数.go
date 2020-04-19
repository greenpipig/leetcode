package main

//罗马数字转整数
func romanToInt(s string) int {
	name := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000}
	counter := 1
	all := 0
	roman := s
	for i := 0; i < len(roman); i++ {
		if i+1 == len(roman) {
			all += counter * name[string(roman[i])]
			break
		}
		if roman[i] == roman[i+1] {
			counter++
		} else if name[string(roman[i])] < name[string(roman[i+1])] {
			all += name[string(roman[i+1])] - counter*name[string(roman[i])]
			i += 1
			counter = 1
		} else {
			all += counter * name[string(roman[i])]

			counter = 1
		}
	}
	return all
}
