package main

import (
	"fmt"
)

func main() {
	ninja1, ninja2 := make(chan string), make(chan string)

	go attack(ninja1, "Ninja 1")
	go attack(ninja2, "Ninja 2")

	select {

	case msg := <-ninja1:
		fmt.Println(msg)

	case msg := <-ninja2:
		fmt.Println(msg)

	}

}

func attack(val chan string, msg string) {
	// for making the default cSE RUN  USE 	time.Sleep(time.Second * 3)
	val <- msg
}
