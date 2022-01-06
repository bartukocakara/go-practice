package tests

func Topla(sayilar []int) int {
	var toplam int
	for i := range sayilar {
		toplam = toplam + sayilar[i]
	}
	return toplam
}