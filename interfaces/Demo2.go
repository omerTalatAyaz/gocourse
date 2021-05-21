package interfaces

import "fmt"

type Mortgage struct {
	creditPaymentTotal float64
	adress             string
	rate               float64
}
type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}
type CreditCalculater interface {
	calculate() float64
}

func (m Mortgage) calculate() float64 {
	return m.creditPaymentTotal * m.rate / 100 / 12

}
func (c Car) calculate() float64 {
	return c.creditPaymentTotal * c.rate / 100 / 12
}

func CalculateMonthlyPayment(credits []CreditCalculater) float64 {
	paymentTotal := 0.0

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].calculate()
	}
	return paymentTotal
}
func Demo2() {
	credit1 := Mortgage{rate: 10, creditPaymentTotal: 100000, adress: "Ankara"}
	credit2 := Mortgage{rate: 12, creditPaymentTotal: 50000, adress: "İstanbul"}
	credit3 := Car{rate: 15, creditPaymentTotal: 60000, carInfo: "Mercedes"}
	credits := []CreditCalculater{credit1, credit2, credit3}
	total := CalculateMonthlyPayment(credits)
	fmt.Println("Toplam Ödeme:", total)

}
