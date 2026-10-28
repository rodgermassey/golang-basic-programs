package main

import (
	"fmt"
)

// ANCHOR - TYPE SWITCH is used when an interface can be one of multiple possible types
func main() {
	var i interface{}

	i = 12

	checkType(i)
}

// ANCHOR - OPTIONAL INTERFACES REFER https://chatgpt.com/share/6ae1a13d-c5e4-83ee-ae66-fc8eda5adaf3

func checkType(i interface{}) {
	switch j := i.(type) { // NOTE - can use shadowing while assigning the value to the j
	case nil:
		fmt.Println("the type is nil", j)
	case int, string: //NOTE - to handle multiple types, the static type of j in this block at compile time is interface{} but at runtime it will have a dynamic type
	// if we try to do (i + 1) it will cause compilation error as the static type is interface{}
		fmt.Println("type is int or string", j)
	default: // NOTE - to handle a type that is not handled by our function
		fmt.Println("cannot define the type", j)
	}
}
