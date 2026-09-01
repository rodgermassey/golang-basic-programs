package main

import (
	"fmt" //fmt is file scoped this import is only availalbe in this file and not the entire package
)

var a = 10 //varaibles defined outside of all the functions are in the package block/ package scope

//This program is to discuss about the various scope types in the golang
func main (){
	x := 20 //x is function scoped or block scoped as it is only available inside the main function

	fmt.Println(a,x)

	//we can also define our own block with it's own scope
	{
		y := 30
		fmt.Println(y) // is only available in this block
	}

	// fmt.Println(y) //y won't be available here

	//ANCHOR - SHADOWING
	//SHADOWED VARIABLES CAN BE DETECTED USING SHADOW LINTER IF WE USE MAKE FILE
	//variables of importance should be avoided to be assigned using := notation as it is very easy to shadow them
	//and with := we can assign multiple variables if there is atleast one new variable on the left hand side
	//(:=) shorthand notation only reuses variables within the current block, when using this make sure we don't have any variables from outer scope on the left hand side unless we want to shadow them 
	b := 10 //b is shadowed in the if block

	if b > 5 {
		fmt.Println(b) //prints 10 before shadowing
		b := 5 //shadowing is done when we declare a variable with the same name inside the block like here are using b
		fmt.Println(b) // prints 5 after shadowing
		//once a variable is shadowed there is no way to access it until the shadowing variable scope ends 
		}

	fmt.Println(b)// prints 10 as the shadowing is not in this block and

	c := 20
	if c > 5 {
		c, d := 5,30 //c is still shadowed
		fmt.Println(c, d)
	}
	fmt.Println(c)

	//Also remember not to shadow an imported package variable
	// fmt := "oops"
	// fmt.Println(fmt)
	//will throw an error type string has no field or method Println


	//ANCHOR - UNIVERSE BLOCK
	/*GOLANG is a small language with only 25 keywords
	  built-in types like int, string, constants like true and false
	  and functions like make, close etc. are not included in this list, neither is nil
	  Go considers these as predeclared identifiers and defines them in universe block
	  universe block contains all the other blocks.
	  Because these are in the universe block these can be shadowed in the inner blocks
	*/

	var true = "100" //this shadows the constant true
	fmt.Println(true)
	false := "200" //this shadows the constant false
	fmt.Println(false)

}