package main

import (
	"fmt"
)

// This program is about the make function
// 1. If the capacity of the slice is known it is better to create the slice initially with that capacity only using the make function
func main() {
	var x = make([]int, 5) // This creates a slice with length and capacity 5 and initializes all the elements with the 0 values
	fmt.Println(x, len(x), cap(x))

	x = append(x, 10) // The capacity of x was 5 and appending a 6th element doubled the capacity of the underlying slice and we reassigned it to x
	fmt.Println(x, len(x), cap(x))

	// capacity is an optional parameter to the make function while creating a slice
	y := make([]int, 5, 10) //This creates a slice with length 5 and capacity 10
	fmt.Println(y, len(y), cap(y))

	y = append(y, 20) // The capacity of y was already 10 therefore no new memory allocation required for the underlying array of the slice
	fmt.Println(y, len(y), cap(y))

	//If values are known we can directly declare the slice with the known values
	z := []int{1,2,4,5,6}
	fmt.Println(z, len(y), cap(y))


	//USING MAKE to create a MAP
	//USE it when we know only the number of key value pairs not the exact values
	//MAPS created with MAKE has initial length = 0, and they can also outgrow the specified size
	d := make(map[string]int, 10) // 10 is the size, length = 0, can outgrow the size i.e. 10
	fmt.Println(d, len(d))
}
