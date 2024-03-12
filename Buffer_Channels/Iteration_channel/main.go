package main

import "fmt"

func main() {
	val := make(chan string)
	go attack(val)
	for {
		msg, open := <-val
		if !open {
			fmt.Println("Channel is closed")
			break
		}
		fmt.Println(msg)
	}
}

func attack(val chan string) {

	for i := 1; i <= 3; i++ {
		val <- fmt.Sprint("Attack at ", i)
	}
	close(val) // this avoids dealdock
}
