package main

import "fmt"

func main() {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	fmt.Println(restoreString(s, indices))
}

func restoreString(s string, indices []int) string {
	result := ""
	mapping := make(map[int]string)
	for i, value := range s {
		mapping[indices[i]] = string(value)
	}
	for i := 0; i < len(s); i++ {
		result += mapping[i]
	}
	return result
}
