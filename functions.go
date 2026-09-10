package main

import (
	"errors"
	"fmt"
)

// ANCHOR - OPTIONAL PARAMETERS FUNCTION
// there is no optional or named parameters in golang, so we can emulate it using a struct which has the fields that we'll use in the function
type OptionalStruct struct {
	Firstname string
	Lastname  string
	Age       int
}

func optionalParam(param OptionalStruct) {
	fmt.Println("running optional param function with parameter: ", param)
}

// ANCHOR - VARIADIC PARAMETERS AND SLICES
// here we specify (...int) type, by adding 3 dots before the type we specify that any number of parameters of this type are accepted in this function
// variadic parameter has to be the last parameter of the parameter list
// this variadic parameter is treated as a SLICE inside the function
// we can also supply a slice to this parameter by appending 3 dots after the slice that we're supplying as on LINE NUMBER 41
func variadicParam(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, val := range vals {
		out = append(out, base+val)
	}
	return out
}

// ANCHOR - MULTIPLE RETURN VALUES
//Enclose the return types in paranthesis when returning multiple values from a function
//error should always be the last returned value if we're returning an error
//No need to enclose return values in paranthesis supply comma separated values that are returned from the function
func mutipleReturn(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("Cannot divide by 0") //No need to enclose return values in paranthesis supply comma separated values that are returned from the function
	}
	return numerator / denominator, numerator % denominator, nil
}

//ANCHOR - NAMED RETURN VALUES
//enclose the return type parameters in paranthesis even if there is a single return type
//these variables i.e. dividend, remainder and err are declared and assigned 0 values
//these variables are available in the function block
//now an empty return statement would still return these variables, if any value is assigned to them that is returned or the 0 values are returned
/*If your function returns values, never use a blank return. It can
  make it very confusing to figure out what value is actually returned
*/
func namedReturnValues (numerator, denominator int) (dividend int, remainder int, err error) {
	if denominator == 0 {
		return 0, 0, errors.New("Cannot divide by 0") //No need to enclose return values in paranthesis supply comma separated values that are returned from the function
	}
	dividend, remainder, err = numerator / denominator, numerator % denominator, nil
	return
}

func main() {
	fmt.Println("Below example is of optional parameters")
	optionalParam(OptionalStruct{Firstname: "Test", Lastname: "Alpha"})

	fmt.Println("Below example is of variadic parameters")
	vals := variadicParam(2, 3, 4, 5)
	fmt.Println("with base 2", vals)
	vals = variadicParam(45)
	fmt.Println("with base 45", vals)
	vals = variadicParam(27, []int{4, 5, 6}...) //supplying a slice to a variadic function parameter
	fmt.Println("with bas 27", vals)

	fmt.Println("Below example is of MULTIPLE RETURN VALUES")
	dividend, remainder, err := mutipleReturn(5, 2) // assigning multiple returned values to variables (dividend, remainder and error)
	fmt.Println(dividend, remainder, err)
	dividend, remainder, err = mutipleReturn(5, 0)
	fmt.Println(dividend, remainder, err)

	fmt.Println("Below example is of NAMED RETURN VALUES")
	dividend, remainder, err = namedReturnValues(5,2)
	fmt.Println(dividend, remainder, err)
}
