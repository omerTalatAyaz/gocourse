package loops

import "fmt"

func Workshop() {

	aklımdakiSayi := 80
	tahminEdilenSayi := 0
	tahminSayısı := 0

	fmt.Println("Bir Sayı Tahmin Ediniz")
	tahminSayısı++
	fmt.Scanln(&tahminEdilenSayi)

	for tahminEdilenSayi != aklımdakiSayi {

		if tahminEdilenSayi < 80 {
			fmt.Println("Daha Büyük Bir Sayı Giriniz")
			tahminSayısı++
			fmt.Scanln(&tahminEdilenSayi)
		}
		if tahminEdilenSayi > 80 {
			fmt.Println("Daha Küçük Bir Sayı Giriniz")
			tahminSayısı++
			fmt.Scanln(&tahminEdilenSayi)
		}
	}

	fmt.Println(tahminSayısı, "seferde buldunuz.")
	if tahminSayısı <= 3 {
		fmt.Println("Helal")
	} else {
		fmt.Println("İdare Eder")
	}

}

//Kullanıcıdan bir sayı girmesini isteyiniz.
//kullanıcının girdiği sayının asal olup olmadığını bul.
func Workshop2() {

	sayi := 0

	fmt.Println("Bir Sayı Tahmin Ediniz")
	fmt.Scanln(&sayi)

	asalMi := true

	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			asalMi = false

		}

	}
	fmt.Println(asalMi)
}

//arkadaş sayılardan 220 ile 284 ü kanıtla...
func Workshop3() {
	sayı1 := 220
	sayı2 := 284
	sayı1Bölen := 0
	sayı2Bölen := 0
	for i := 1; i < sayı1; i++ {
		if sayı1%i == 0 {
			sayı1Bölen = sayı1Bölen + i

		}
	}
	for i := 1; i < sayı2; i++ {
		if sayı2%i == 0 {
			sayı2Bölen = sayı2Bölen + i

		}

	}
	if sayı1 == sayı2Bölen && sayı2 == sayı1Bölen {
		fmt.Printf("%v ve %v arkadaş sayılardır", sayı1, sayı2)
	} else {
		fmt.Println(sayı2, "ve", sayı1, "arkadaş sayılar değildir.")
	}
}
