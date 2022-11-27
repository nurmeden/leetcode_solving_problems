package main

import "fmt"

func main() {
	x := 123
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	if x < 0 {
		flag := true
		if x < 0 {
			x = -x
			flag = false
		}
		sum := 0
		for x > 0 {
			a := x % 10
			sum = sum*10 + a
			x /= 10
		}
		if !flag {
			sum *= -1
		}
		return sum
	} else {
		return 0
	}
}
