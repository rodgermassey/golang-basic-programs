package main

import (
	"fmt"
)

type AirthAdder interface {
	Adder(a, b int) int
}

type Airth func(a, b int) int

func (ai Airth) Adder(a, b int) int {
	return ai(a, b)
}

func realAdder(a, b int) int {
	return a + b
}

func main() {
	i := Airth(realAdder)

	fmt.Println(i.Adder(2, 4))
	realSum(i)
}

// NOTE - This is where we use interface and function types,
// 1. we create an interface that defines the function, which will be passed in the parameter list of other functions (AirthAdder)
// 2. then we create a function type of the same signature as that interface method (Airth)
// 3. (Adder on Airth satisfies AirthAdder)then we create a method on that function type so that it implements that interface created before, this method simply calls the function type on which it is being defined
// 4. (realAddder typecasted to Airth and passed to realSum which accepts the interface AirthAdder) we create a function that actually does satisfy the function type, then we type cast our function to that type so that the method is attached onto it, 
// and then we pass it into the 3rd party function that accepts the interface, now here the 3rd party function only knows that the incoming value satisfies the interface
// and therefore it knows that it will have the required method
func realSum (f AirthAdder) {
	fmt.Println(f.Adder(2, 3))
}