package main

import "fmt"

type vehicle struct {
	name string
	car
}
type car struct {
	id   string
	tank float64
}

func main() {
	obj := vehicle{}
	// can dirctly access without using car
	obj.name = "Vehicle 1"
	obj.id = "HR-34T-4322"
	obj.tank = 102.6

	fmt.Println(obj)
}
