package main

import "fmt"

type car struct {
	id   string
	name string
	tyre Wheel
}
type Wheel struct {
	radius   float64
	material string
}

func main() {
	car1 := car{
		id:   "HR-27B-6839",
		name: "CRETA",
		tyre: Wheel{
			radius:   52.3,
			material: "Nylon",
		},
	}

	car2 := car{}

	car2.id = "MH-89S-4T93"
	car2.name = "VERNA"
	car2.tyre.material = "Polyster"
	car2.tyre.radius = 45.9

	fmt.Println(car1)
	fmt.Println(car2)
}
