package main

import "fmt"

func main() {

	vehicle := map[int]map[string]float64{
		1: {
			"Creta": 52.3,
		},
		2: {
			"Audi": 52.3,
		},
		3: {
			"BMW": 52.3,
		},
	}

	value, ok := vehicle[2] // getting the map with id 2

	fmt.Println(value, ok)

	value2, ok2 := value["Audi"] // to get the 2nd value of map
	fmt.Println(value2, ok2)

}
