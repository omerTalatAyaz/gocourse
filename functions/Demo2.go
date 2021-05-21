package functions

func Dörtİslem(sayi1, sayi2 int) (int, int, int, float64) {
	toplam := sayi1 + sayi2
	çarpım := sayi1 * sayi2
	fark := sayi1 - sayi2
	bölüm := (float64(sayi1) / float64(sayi2))
	return toplam, çarpım, fark, bölüm
}
