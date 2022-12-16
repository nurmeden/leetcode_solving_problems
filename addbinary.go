package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
	b := "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
	fmt.Println(addBinary(a, b))
	fmt.Println(binary_to_decimal2(a) + binary_to_decimal2(b))
}

func addBinary(a string, b string) string {
	a_decimal, _ := strconv.Atoi(a)
	b_decimal, _ := strconv.Atoi(b)
	dec := binary_to_decimal(a_decimal, len(a)) + binary_to_decimal(b_decimal, len(b))
	return decimal_to_binary(dec)
}

func binary_to_decimal2(a string) int {
	ss := len(a)
	r := 0
	san := 0
	for _, value := range a {
		s, _ := strconv.Atoi(string(value))
		pow_num := 1
		if ss == 0 {
			pow_num = 1
		} else {
			for i := 0; i < r; i++ {
				pow_num *= 2
			}
			r++
		}
		ss--
		san += (s * pow_num)
	}
	return san
}

func binary_to_decimal(a int, l int) int {
	ss := l
	if a == 0 {
		return 0
	} else if a == 1 {
		return 1
	} else {
		r := 0
		san := 0
		for i := 0; i < l; i++ {
			k := a % 10
			pow_num := 1
			if ss == 0 {
				pow_num = 1
			} else {
				for i := 0; i < r; i++ {
					pow_num *= 2
				}
				r++
			}
			ss--
			san += (k * pow_num)
			a /= 10
		}
		return san
	}
}

func decimal_to_binary(a int) string {
	ch := ""
	for a != 0 {
		san := a % 2
		s := strconv.Itoa(san)
		ch = s + ch
		a /= 2
	}
	return ch
}
