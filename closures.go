package main

import (
	"fmt"
	"sort" //using sort standard package to sort the slices
)

type person struct {
	Firstname string
	Lastname  string
	Age       int
}

// ANCHOR - RETURNING FUNCTION FROM FUNCTION
// used in middleware when creating a web server
func makeMultiplier(base int) func(int) int {
	return func(value int) int {
		return base * value
	}
}

// ANCHOR - CLOSURES
// closures are basically anonymous functions which are declared inside another functions
// the declared function encloses all the data from the outer function and can be reused
func main() {
	persons := []person{
		{"Michael", "Jordan", 23},
		{"Ray", "Allen", 24},
		{"Kobe", "Bryant", 25},
	}

	//ANCHOR - FUNCTION AS PARAMETER VALUE
	// here we're passing a function as a parameter to another function these are called higher order functions
	//sort.Slice is taking a function as a parameter a and b are the indices of the slice that we pass as the first parameter
	/*Passing functions as parameters to other functions is often useful
	for performing different operations on the same kind of data.
	*/
	sort.Slice(persons, func(a, b int) bool {
		return persons[a].Age > persons[b].Age
	})

	fmt.Println(persons)

	//NOTE- here we have recieved closured function in base2 and base3 
	//the returned function will remember the base that has been passed 2 and 3 in our case
	base2 := makeMultiplier(2)
	base3 := makeMultiplier(3)

	fmt.Println(base2(5), base3(7))
}
