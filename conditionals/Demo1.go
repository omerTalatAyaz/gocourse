package conditionals

import "fmt"

func Demo1() {

	var hesap float64 = 750

	var çekilmekistenen float64 = 1250

	if hesap > çekilmekistenen {
		fmt.Println("bakiye Yeterli")
	} else {
		fmt.Println("bakiye Yetersiz")
	}
}
func Demo2(x, y float64) {

	fuck := x + 5*y

	fmt.Println(fuck)

}
