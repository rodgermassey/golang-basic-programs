package main

import (
	"fmt"
	"math/rand"
)

//IF in golang
//if is just like normal languages the only thing is we don't enclose the conditions in paranthesis in golang

func main() {
	n := rand.Intn(10)

	if n == 0 {
		fmt.Println("The number is too low")
	} else if n > 5 {
		fmt.Println("The number is too high", n)
	} else {
		fmt.Println("This is a good number", n)
	}

	//ANCHOR SOPING IF CONDITION VARIABLE
	// scoping variable to just condition and if else block
	//having this scope is very handy because the variables is only present when it's needed outside of that it is not accessible and ready for garbage collection
	// As with shadowing here n is shadowing the variable declared previously in the program
	//Technically we can use any simple statement after if but we'll not do that because it can make the program confusing so we'll use it to assign new variable values
	if n := rand.Intn(10); n == 0 {
		fmt.Println("n is too low")
	} else if n > 5 {
		fmt.Println("n is too high", n)
	} else {
		fmt.Println("n is a good number", n)
	}

	// fmt.Println(a) //This will throw error as a is not accessible outside the scope of the if else condition block
	fmt.Println(n) //this will run
}
