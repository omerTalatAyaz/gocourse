package channels

import (
	"fmt"
	"time"
)

func CiftSayilar(CiftSayiCn chan int) {
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam = toplam + i
		fmt.Println("Çift Sayı Toplama Çalısıyor")
		time.Sleep(time.Second)

	}
	CiftSayiCn <- toplam
}
func TekSayilar(TekSayiCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam = toplam + i
		fmt.Println("Tek Sayı Toplama Çalısıyor")
		time.Sleep(time.Second)

	}
	TekSayiCn <- toplam
}
