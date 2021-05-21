package arrays

import "fmt"

func Demo1() {

	sayilar := [5]int{12, 234, 53, 34, 34}

	fmt.Println(sayilar)

}
func Demo2() {
	var sehirler [5]string
	sehirler[0] = "ankara"
	sehirler[1] = "istanbul"
	sehirler[2] = "izmer"
	sehirler[3] = "adana"
	sehirler[4] = "diyarbakır"

	for i := 0; i <= 4; i++ {
		fmt.Println(sehirler[i])
	}
}
