package conditionals

import "fmt"

func Workshop1(x, y, z int) {
	//3 adet değişken tanımla.
	// Ekrana en büyük olanı yazdır.

	/*if x > y {
		if x > z {
			fmt.Println(x, "en büyük sayıdır")
		} else {
			fmt.Println(z, " en büyük sayıdır")
		}

	} else if x < y {
		if y > z {
			fmt.Println(y, "en büyük sayıdır")

		} else {
			fmt.Println(z, "en büyük sayıdır")
		}
	}*/

	enBüyük := x
	if y > enBüyük {
		enBüyük = y
	}
	if z > enBüyük {
		enBüyük = z
	}
	fmt.Printf("en büyük sayı: %v", enBüyük)

}
