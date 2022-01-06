package tests

import "testing"

func TestTopla(t *testing.T) {
	sonuc := Topla([]int{1, 3, 5})
	if sonuc != 9 {
		t.Error("Result", sonuc)
	}
}

func Topla(numbers []int) int {
	var total int
	for _, value  := range numbers {
		total += value
	}

	return total
}

// func TestTopla(t *testing.T) {
// 	sonuc := Topla([]int{2, 5})
// 	if sonuc != 7 {
// 		t.Error("Beklenen sonuc 7, elde edilen ", sonuc)
// 	}
// }

// func Topla(sayilar []int) int {
// 	var toplam int
// 	for i := range sayilar {
// 		toplam = toplam + sayilar[i]
// 	}
// 	return toplam
// }
