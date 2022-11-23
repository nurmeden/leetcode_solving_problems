package main

import "fmt"

func main() {
	address := "1.1.1.1"
	fmt.Println(defangIPaddr(address))
}

func defangIPaddr(address string) string {
	var result string
	for i := range address {
		if address[i] == '.' {
			result += "[.]"
		} else {
			result += string(address[i])
		}
	}
	return result
}
