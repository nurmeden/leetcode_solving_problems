package main

import "fmt"

func main() {
	paths := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	fmt.Println(destCity(paths))
}

func destCity(paths [][]string) string {
	result := ""
	mapping := make(map[string]string)
	for i := 0; i < len(paths); i++ {
		mapping[paths[i][0]] = paths[i][1]
	}
	for _, values := range mapping {
		if len(mapping[values]) == 0 {
			result = values
		}
	}
	return result
}
