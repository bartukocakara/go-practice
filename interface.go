package main

import "fmt"

type FindVovel interface {
	VovelFinder() []rune
}

type MyString string

func main() {
	var name = MyString("Bartu Kocakara")
	var v FindVovel

	v = name
	fmt.Printf("chars are %c", v.VovelFinder())
}

func (ms MyString) VovelFinder() []rune {
	var chars []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'b' || rune == 'k' || rune == 't' {
			chars = append(chars, rune)
		}
	}

	return chars
}