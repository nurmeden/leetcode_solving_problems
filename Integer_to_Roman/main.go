package main

import (
	"fmt"
)

type Roman struct {
	Roman_symbols string
	num           int
}

func main() {
	num := 555
	intToRoman(num)
}

func intToRoman(num int) string {
	arr := []Roman{}
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	i := 0
	for keys, values := range m {
		arr = append(arr, Roman{keys, values})
		i++
	}

	fmt.Println(arr)
	return ""
}
