package main

import "fmt"

func main() {
	var dict = map[string]int{
		"x":   0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}
	if val, ok := dict["foo"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("olmadı")
	}
}
