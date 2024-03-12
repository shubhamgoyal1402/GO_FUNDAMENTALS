package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	name := []string{"Shubham", "manan", "raghav"} //, "srilekha", "dwij", "Ramesh", "aryan"}

	val := make(chan bool) // helps in not taking any more extra time that time.sleep takes for main fn not to exit
	for _, names := range name {
		go attack(names, val)
		//val <- false channel cannot contains two value for now as it holds only one at a time
		fmt.Println(<-val)
	}

	fmt.Println(time.Since(start))
}

func attack(name string, val chan bool) {

	val <- true
	fmt.Println("Attack at", name)

}
