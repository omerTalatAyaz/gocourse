package maps

import "fmt"

func Demo1() {
	sozluk := make(map[string]string)
	sozluk["table"] = "masa"
	sozluk["book"] = "kitap"
	sozluk["pencil"] = "kalem"
	fmt.Println(sozluk["book"])
	fmt.Println("Eleman Sayısı:", len(sozluk))
	delete(sozluk, "book")
	fmt.Println(sozluk["book"])
	fmt.Println("Eleman Sayısı:", len(sozluk))

	deger, varMi := sozluk["book"]
	fmt.Println(deger)
	fmt.Println("listede olma durumu:", varMi)

	sozluk2 := map[string]string{
		"glass":     "bardak",
		"microphon": "mikrofon",
	}
	fmt.Println(sozluk2)
}
