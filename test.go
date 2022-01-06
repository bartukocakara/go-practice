package main

import "fmt"

func main() {
	var width = 10
	var height = 5

	for i := 1; i <= height; i++ {
		for j := 1; j <= width; j++ {
			if i == 1 || i == height {
				fmt.Print("*")
			} else if j == 1 || j == 10 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}