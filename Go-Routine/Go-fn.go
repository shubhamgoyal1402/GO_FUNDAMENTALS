package main

import (
	"fmt"
	"time"
)

func main() {
	userId := 1
	go getUser(userId)
	time.Sleep(time.Millisecond * 100) // so that it gives time for go routine to complete
	fmt.Println("Go routine")
}

func getUser(id int) {
	time.Sleep(time.Millisecond * 50)
	fmt.Println(id)
}
