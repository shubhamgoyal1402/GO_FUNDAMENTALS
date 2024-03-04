package main

import "fmt"

func main() {

	names := []string{"ram", "sham", "raju", "pratham", "kiran"}
	ages := []int{40, 25, 18, 43, 26}
	// map structure : map_name := map[key_datatype]value_datatype{}
	details := map[string]int{} //declaration of a map
	for i := 0; i < len(names); i++ {
		details[names[i]] = ages[i] // initialzing a map
	}

	fmt.Println(details)
}
