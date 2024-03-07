package main

import (
	"fmt"
)

func main() {

	ch := make(chan int) // channel

	go func() {
		for i := 0; i < 1; i++ {
			ch <- i // passing value in channel
			ch <- i + 1

		}
	}()
	j := <-ch

	fmt.Println(j)
	j = <-ch //taking the value out of channel
	fmt.Println(j)
	//j = <-ch  pulling more from channel than it have creates deadlock
	//	fmt.Println(j)
}
