package main

import (
	"fmt"
	"strconv"
)

//functions are values in golang
/*The type of a function is built out of the keyword func and the types of the parameters and return values
  2 functions who have the same number of parameters with same types
  and same number of return values with same type have the same signature
*/

type calcFuncType func(int, int) int

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

//ANCHOR - all the functions add, sub, multiply and divide are of type calcFuncType
//therefore we can assign them as values to the opMap

var opMap = map[string]calcFuncType{
	"+": add,
	"-": sub,
	"*": multiply,
	"/": divide,
}

func main() {
	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"2", "|", "3"},
		[]string{"a", "/", "3"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("Invalid expression min. length should be 3")
			continue
		}

		//NOTE - need to extract it separately, as if we use shorthand of if the operatorFunction and ok variables will only be if block scoped
		operatorFunction, ok := opMap[expression[1]]
		if !ok {
			fmt.Println("Invalid operator, not found in the list")
			continue
		}

		//NOTE - need to extract these separately as strconv.Atoi returns 2 values the integer and the error
		//if strconv.Atoi cannot convert it returns 0 value of int i.e. 0 and the error
		//therefore cannot directly pass into the operatorFunction
		op1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println("The first operand for the expression", expression, "is not valid")
			continue
		}

		op2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println("The second operand for the expression", expression, "is not valid")
			continue
		}

		fmt.Println(operatorFunction(op1, op2))
	}
}
