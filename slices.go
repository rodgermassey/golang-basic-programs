package main

import (
	"fmt"
)

// In place of arrays gophers tend to use slices as the length is not the part of the type slice
// so whenever we want the sequence of values to be stored we use the slices

//SLICES WORK WITH SLICES, ARRAYS AND STRINGS
// Slices cannot be compared using == or != the only thing the slices can be compared to is nil.
// The reflection package has a function DeepEqual that can be used to compare slices but this is primarily used for testing
func main() {
	var x []int                           // creates a nil slice
	fmt.Println(x == nil, len(x), cap(x)) // len can be used to get the length of the slice
	//cap is used to get the capacity of the slice
	// capacity is nothing but the memory allocated by the goruntime before which the go runtime will not reallocate the new underlying array for the slice


	//If values are known we can directly declare the slice with the known values
	z := []int{1,2,4,5,6}
	fmt.Println(z, len(y), cap(y))

	//One IMPORTANT thing to remember is to assign the value returned by the append function back to the slice variable that was passed in or else it wont update the previous underlying slice properties
	// GO is a primarily the pass by value language now append creates a copy of the slice that is passed in and appends the values to this copy and then returns the copy so we need to assign it back to the slice that was passed in
	// It is a compile time error if we don't assign the returned value from the append to a variable
	x = append(x, 4) // append is used to append the values at the end of the slices
	fmt.Println(x, len(x), cap(x))

	x = append(x, 5, 6, 7) // we can also append multiple values in a single go
	fmt.Println(x, len(x), cap(x))

	y := []int{89, 90, 91}
	var a = append(x, y...) // we use (slice...) operator in the append function to append a slice to another slice
	// like here we append slice y to slice x
	fmt.Println(a, len(a), cap(a))

	// CREATING SLICES FROM SLICES/ARRAYS
	var s = []int{10, 20, 30, 40, 50}
	s1 := s[:3]
	fmt.Println(s1) //prints 10,20,30

	s2 := s[3:]
	fmt.Println(s2) //prints 40, 50

	s3 := s[1:3]
	fmt.Println(s3) //prints 20,30

	s4 := s[:]
	fmt.Println(s4) //prints 10,20,30,40,50

	//IMPORTANT when creating slices from an already existing slice, the underlying array remains the same
	// therefore changing the value changes it in the original slice as well
	s4[2] = 99 // this will update the 2 index in the original (s) slice as well
	fmt.Println(s) // s will have 99 at the 2nd index OUTPUT => 10,20,99,40,50

	//SLICING FULL SYNTAX
	// s5 := [:2:2] //The last part tells to subtract the starting index from this max (low:high:max) new capacity is max - low
	// s6 := [2:4:4] //The new calculated capacity cannot be less than the length of the slice
	//IMPORTANT therefore it is idiomatic to set the max as the high index of the offset i.e. 4 in our case

	//COPYING A SLICE/ARRAY
	s7 := []int{1,2,3,4}
	s8 := make([]int,2)
	num := copy(s8,s7)
	fmt.Println(s8,num) //s8 is 1,2  and num is the number of elements copied from the source to the destination slice which is 2 as the length of the destination slice is 2
	// CAPACITY DOESN'T MATTER while COPYING, LENGTH MATTERS
	//whichever slice is smaller is smaller the source or the destination the elements are copied accordingly 
	//if the destination length is less only that much elements are copied
	//if the destination length is greater then all the remaining places are filled with the zero values
	
	//COPYING A PORTION OF SLICE/ARRAY TO ANOTHER SLICE
	s9 := make([]int, 2)
	copy(s9,s7[2:])//This will copy a portion of the slice/array to the destination
	fmt.Println(s9)

	//OVERWRITING SLICE with overlapping elements
	s10 := []int{1,2,3,4,5,6}
	copy(s10[:3], s10[2:]) //It will override the elements of offset 0 to 3 with the elements from the offset 2:5
	fmt.Println(s10) //3,4,5,4,5,6
}
