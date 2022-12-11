package main

import "fmt"

func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
}

func uniqueMorseRepresentations(words []string) int {
	MorseArr := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	representations := make(map[rune]string)
	result_mapping := make(map[string]int)
	l := 0
	for i := 'a'; i <= 'z'; i++ {
		representations[i] = MorseArr[l]
		l++
	}
	for i := 0; i < len(words); i++ {
		result := ""
		for _, letter := range words[i] {
			result += representations[letter]
		}
		words[i] = result
	}
	for i := 0; i < len(words); i++ {
		for j := i; j < len(words); j++ {
			result_mapping[words[j]]++
		}
	}
	return len(result_mapping)
}
