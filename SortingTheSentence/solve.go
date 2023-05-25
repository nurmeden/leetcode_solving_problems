package sortingthesentence

import (
	"strconv"
	"strings"
)

func sortSentence(s string) string {
	result := ""
	sliceS := strings.Split(s, " ")
	isDone := false
	for !isDone {
		isDone = true
		for i, _ := range sliceS {
			if i != 0 {
				intStrСurr, _ := strconv.Atoi(string(sliceS[i][len(sliceS[i])-1]))
				intStrPrev, _ := strconv.Atoi(string(sliceS[i-1][len(sliceS[i-1])-1]))
				if intStrСurr < intStrPrev {
					sliceS[i], sliceS[i-1] = sliceS[i-1], sliceS[i]
					isDone = false
				}
			}
		}
	}
	for i, value := range sliceS {
		if i != len(sliceS)-1 {
			result += value[:len(value)-1] + " "
		} else {
			result += value[:len(value)-1]
		}
	}
	return result
}
