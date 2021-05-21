package for_range

import "fmt"

func Demo1() {
	cities := []string{"ankara", "istanbul", "izmir"}
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
	for _, city := range cities {
		fmt.Println(city)
	}
}
