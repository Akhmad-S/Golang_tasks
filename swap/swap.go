// Swap elements of array
package main

import "fmt"

func main() {
	workArray := [10]uint8{}
	var a, b uint8

	for i := range workArray {
		fmt.Scan(&workArray[i])
	}
	for j := 0; j < 3; j++ {
		fmt.Scan(&a, &b)
		workArray[a], workArray[b] = workArray[b], workArray[a]
	}
	for _, elem := range workArray {
		fmt.Print(elem, " ")
	}
}
