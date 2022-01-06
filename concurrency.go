package main

import (
	"fmt"
	"time"
)

func main() {
	var heros = []string{"spiderman", "batman", "hulk", "thor", "flash"}
	go bulucuA(heros)
	go bulucuB(heros)
	fmt.Scanln()
}

func bulucuA(heros []string) {
	for i := 0; i < len(heros); i++ {
		if heros[i] == "flash" {
			fmt.Println("Bulucu A: " + heros[i]  + "buldu")
		}
		time.Sleep(time.Second)
	}
}

func bulucuB(heros []string) {
	for i := 0; i < len(heros); i++ {
		if heros[i] == "hulk" {
			fmt.Println("Bulucu B: " + heros[i] + "bulundu")
		}
		time.Sleep(time.Second)
	}
}