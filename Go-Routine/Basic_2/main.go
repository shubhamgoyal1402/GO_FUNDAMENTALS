package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	evilninja := []string{"Shubham", "Ayush", "Shreyansh", "Rudranil", "Navya"}

	for _, name := range evilninja {

		go attack(name) // By go Routine all the ninjas are attacked within a second otherwise it would take  5sec
	}
	time.Sleep(time.Second) // this will increase the time of the mission in next we will use channels for minimal time
	fmt.Println(time.Since(start))

}
func attack(name string) {
	fmt.Println("attack ", name)
	time.Sleep(time.Second)

}
