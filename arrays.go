package main

import (
	"fmt"
)

// In go length of the array is also a part of the type of the array
// therefor we cannot compare an array of length 3 to an array of length 4 and they cannot be type converted as well

func main() {
	// var x = [3]int{10, 20, 30}   // creates an array with 10, 20 and 30 elements at 0,1 and 2
	// var y [3]int                 // creates an array and initializes all the elements as 0
	// var z = [...]int{11, 21, 30} // in this syntax we don't have to specify the length of the array

	// we can also create sparse arrays in golang
	var a = [...]int{1, 4: 10, 5, 7: 32} // this syntax tells to put 1 at index 0, then 10 at index 4, and 32 at index 7 rest all places are filled with 0
	fmt.Println(a, len(a), cap(a)) // len is used to get the length of an array & cap is used to get the capacity of the array
}
