package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"bilgisayar", 500, "XYZ", 0})
	fmt.Println(product{name: "bilgisayar", unitPrice: 500})

}

type product struct {
	name         string
	unitPrice    float64
	brand        string
	discountRate int
}
