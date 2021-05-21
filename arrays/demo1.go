// dizideki sayıların en büyüğünü bulan algoritma

package arrays

import "fmt"

func Demo3() {
	sayilar := [...]int{20, 30, 58, 10, 8}
	enBuyuk := sayilar[0]

	for i := 0; i < len(sayilar); i++ {
		if enBuyuk < sayilar[i] {
			enBuyuk = sayilar[i]
		}

	}
	fmt.Println(enBuyuk)
}
