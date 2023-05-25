package main

import (
	"fmt"
)

func main() {
	command := "G()(al)"
	fmt.Println(Interpret(command))
}

func Interpret(command string) string {
	result := ""
	for i := 0; i < len(command); i++ {
		if command[i:i+2] == "()" {
			result += "o"
		} else if command[i:i+4] == "(al)" && i < len(command)-5 {
			result += "al"
		} else {
			result += "G"
		}
	}
	return result
}
