package main

import "fmt"

func main() {
	array := []string{"a", "b", "c", "d", "d"}

	fmt.Println("first", array)

	fmt.Println("second", unique(array))
}

func unique(array []string) []string {
	copy := make([]string, 0, len(array))
	keys := make(map[string]string)

	for _, v := range array {
		if _, ok := keys[v]; !ok {
			keys[v] = v
			copy = append(copy, v)
		}
	}

	return copy
}