package main

import (
	"fmt"
)

// ANCHOR - ANONYMOUS FUNCTIONS
// these are created by using the func keyword and then directly supplying the parameters and the return types
// In golang we can't create functions inside another function,
// but ANONYMOUS FUNCTIONS are an exception to this
func main() {
	for i := 0; i < 10; i++ {
		//As anonymous functions don't have names we need to invoke them as soon as we create them
		func(j int) {
			fmt.Println("current index is:", j)
		}(i)
	}
}
