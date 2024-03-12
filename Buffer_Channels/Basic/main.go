package main

import "fmt"

/*func main() {

	val := make(chan string)

	val <- "Shubham" // this will create deadlock cause the deafult capacity of channel is 0 and it cannot hold value
	fmt.Println(<-val)

}*/

func main() {
	// Buffer channel
	val := make(chan string, 1) // make(chan TYPE, LENGTH)

	val <- "Shubham" // it will hold the value
	//val <- "NMikhil" // this will again create deadlock
	fmt.Println(<-val)

	val <- "Nikhil"
	fmt.Println(<-val)
}
