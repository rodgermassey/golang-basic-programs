package main

import (
	"fmt"
)

type MyInt int

// ANCHOR - TYPE ASSERTIONS ARE USED WITH INTERFACES 
// 1. type assertions are used to check the underlying concrete type of the interface
// 2. if our interface implements another interface then type assertion can point to that interface as well
// 3. Type assertions are different from type conversion, conversion is checked at compile time, assertion is checked at runtime, thus assertions cause panic
// conversion convert the values and assertions reveal the underlying types


func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
// NOTE - here we assert that i of type interface{} to be of type MyInt
// NOTE - if the type assertion fails our program panics for example i.(int) will cause a panic,
// this panics because go checks the compatibility with the underlying type, here i2 is of type MyInt and we are trying to add MyInt with int on line 22
	i2 := i.(MyInt) 
	fmt.Println(i2 + 1)


// ANCHOR - to avoid the panic we use the COMMA-OK idiom with the type assertion
	// if i2, ok := i.(int); ok {
	// 	fmt.Println(i2, ok)
	// }	

// COMMENT LINE 21 AND 22 WHILE RUNNING THE BELOW CODE AS WE'RE REUSING THE VARIABLE i2 OR JUST RENAME THE VARIABLE TO i3
// NOTE - error => cannot use i.(int) (comma, ok expression of type int) as MyInt value in assignment
// if i.(int) is used in place of MyInt in type assertion on line 30
	i3, ok := i.(int) // THIS DOESN'T PANIC BECAUSE OF COMMA OK IDIOM
	if ok {
		fmt.Println("i3 is compatible with the given type", i3)
	} else {
		fmt.Println("i3 is not compatible with the given type", i3)
	}

}
