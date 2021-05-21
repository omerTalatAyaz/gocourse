package for_range

import "fmt"

//1-10 arasındaki sayılardan tek olanları toplayan program.
func Demo2() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	toplam := 0
	for _, num := range nums {
		if num%2 != 0 {
			toplam = num + toplam
		}
	}
	fmt.Println("toplam", toplam)

}
