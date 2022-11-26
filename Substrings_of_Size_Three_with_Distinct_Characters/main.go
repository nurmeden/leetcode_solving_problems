package main

import "fmt"

func main() {
	s := "xyzzaz"
	fmt.Println(countGoodSubstrings(s))
}

func countGoodSubstrings(s string) int {
	strs_arr := []string{}
	count := 0
	res := 0
	result := ""
	for i := 0; i < len(s); i++ {
		if i <= len(s)-3 {
			for j := i; j < i+3; j++ {
				if check(result, string(s[j])) {
					count++
				}
				result += string(s[j])
			}
			if count == len(result) {
				res++
			}
			strs_arr = append(strs_arr, result)
			result = ""
			count = 0
		} else {
			strs_arr = append(strs_arr, s[len(s):])
		}
	}
	return res
}

func check(s string, r string) bool {
	for _, value := range s {
		if string(value) == string(r) {
			return false
		}
	}
	return true
}

// func countGoodSubstrings(s string) int {
// 	count := 0
// 	for i := 0; i < len(s); i++ {
// 		str1 := ""
// 		mk := make(map[string]int)
// 		for j := i; j < len(s); j++ {
// 			str1 += string(s[j])
// 			mk[string(s[j])] = count
// 			if len(str1) == 3 && len(mk) == 3 {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }

// func countGoodSubstrings(s string) int {
// 	numGoodSubs := 0
// 	for i := 0; i < len(s)-2; i++ {
// 		if s[i] != s[i+1] && s[i] != s[i+2] && s[i+1] != s[i+2] {
// 			numGoodSubs++
// 		}
// 	}
// 	return numGoodSubs
// }
