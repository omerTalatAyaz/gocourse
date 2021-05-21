package arrays

import "fmt"

func Demo4() {
	var sayilar [3][2]int
	sayilar[0][0] = 1
	sayilar[0][1] = 3
	sayilar[1][0] = 5
	sayilar[1][1] = 2
	sayilar[2][0] = 4
	sayilar[2][1] = 6

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(sayilar[i][j])
			fmt.Print("")
		}
		fmt.Println()
	}
	fmt.Println(sayilar)
}
