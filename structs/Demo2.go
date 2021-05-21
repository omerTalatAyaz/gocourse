package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() {
	fmt.Println("Eklendi:", c.firstName)
}
func (c customer) update() {
	fmt.Println("Güncellendi:", c.lastName)
}
func Demo2() {
	c := customer{firstName: "omer", lastName: "ayaz", age: 29}
	c.save()
	c.update()
}
