package slices

import "fmt"

func Demo1() {

	isimler := make([]string, 3) // LİTERAL METHOD

	isimler[0] = "engin"
	isimler[1] = "derin"
	isimler[2] = "salih"
	isimler = append(isimler, "ali")
	fmt.Println(isimler)
	fmt.Println(len(isimler))
}
func Demo2() {
	sehirler := []string{"ankara", "istanbul", "izmir"} // MAKE METHOD
	sehirlerKopya := make([]string, len(sehirler))
	copy(sehirlerKopya, sehirler)
	sehirler = append(sehirler, "adana")
	fmt.Println(sehirler)
	fmt.Println(sehirlerKopya)
	fmt.Println(sehirler[1:3])
	fmt.Println(sehirler[1:])
	fmt.Println(sehirler[:2])
}
