package main

import "fmt"

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	result := ""
	max_len := 1
	l := len(s)
	if len(s) >= 1 {
		for i := 0; i < l; i++ {
			for j := i; j < l; j++ {
				for _, value := range s[i : j+1] {
					if Checker(result, value) {
						result += string(value)
					} else {
						// if len(result) > max_len {
						//     max_len = len(result)
						// }
						result = ""
						result += string(value)
					}
				}
				if len(result) > max_len {
					max_len = len(result)
				}
				result = ""
			}
		}
		return max_len
	} else {
		return 0
	}
}

func Checker(s string, r rune) bool {
	for _, v := range s {
		if v == r {
			return false
		}
	}
	return true
}
