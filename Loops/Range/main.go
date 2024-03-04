package main

import "fmt"

func main() {

	cars := []string{"HONDA CITY", "VERNA", "CRETA", "BMW", "AUDI"}

	for car := range cars {
		fmt.Println(cars[car]) // car represnt index
	}

	for i, car := range cars { // i represnt index and car represent element on that index
		fmt.Println(i, car)

	}
}
