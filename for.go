package main

import (
	"fmt"
)

// FOR in golang
// FOR is the only loop available in golang
// The conditions are not enclosed in the paranthesis
func main() {
	//ANCHOR - 1. COMPLETE FOR STATEMENT
	//only short hand variable initialization will work, var will not work in for loop variable initialization
	//we can initialize 2 or more variables
	//shadowing will work as it works with if statement block, here in the for block as well.
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//ANCHOR - 2. CONDITION ONLY (Works as while loop in other languages)
	i := 1
	for i < 100 {
		fmt.Println(i)
		i *= 2
	}

	//ANCHOR - 3. INFINITE FOR STATEMENT (SKIPS INIT, CONDITION and CHANGE TO CONDITION VARIABLE)
	// for {
	// 	fmt.Println("HELLO")
	// }

	//ANCHOR - BREAK AND CONTINUE
	//BREAK is used to programatically exit the loops be it normal or infinite loop in golang
	//The example below is similar to do while loop as it starts from 0 and before breaking as well it executes what the loop was supposed to do
	i = 0
	for {
		fmt.Println("Hello", i)
		if i == 10 {
			break
		}
		i++
	}

	//CONTINUE is used to skip over rest of the body of the loop and proceed directly to next iteration
	//Technically we don't need a continue statement but it makes the code more readable and syntactically appealing
	//The commented code below can be written as following
	// for i := 1; i <= 100; i++ {
	// 	if i%3 == 0 {
	// 		if i%5 == 0 {
	// 			fmt.Println("FIZZBUZZ")
	// 		} else {
	// 			fmt.Println("FIZZ")
	// 		}
	// 	} else if i%5 == 0 {
	// 		fmt.Println("BUZZ")
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }

	//GO encourages short if bodies and as left aligned as possible, nested code is difficult to follow
	//using CONTINUE here will help make it easier
	//Replacing the chains of IF-ELSE with IF's and continue imporves the layout of conditions and lines up the conditions making it easier to read the code
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FIZZBUZZ")
			continue
		}
		if i%3 == 0 {
			fmt.Println("FIZZ")
			continue
		}
		if i%5 == 0 {
			fmt.Println("BUZZ")
			continue
		}
		fmt.Println(i)
	}

	//ANCHOR - 4. FOR-RANGE STATEMENT
	//for-range loop resembles iterators in different programming languages
	//it is used to iterate over some of GO's built-in types like arrays, strings, slices, maps and channels
	/*when iterating over maps with for range the order of the keys is not fixed, for 2 reasons :-
	1. People should not assume that the keys are ordered in the map while looping
	2. To prevent Hash Dos attack so that the attacker cannot send specially crafted data where all the keys would be stored in a single bucket thus making the server slow
	To prevent these issues GOlang's team made 2 changes they first changed the hash algorithm to include a random number that's generated every time a map variable is created
	and second they made a for-range iteration over a map vary a bit each time the map is looped over*/
	//There is one exception to this fmt.Println always prints the keys in ascending order so that debugging is easy

	/*when iterating over strings the for range loop iterates over runes,
	if there is a multi byte rune then it is converted into the utf-8 representation into a single 32-bit number
	and the offset is incremented by the number of bytes in the rune
	*/
	slice := []int{1, 2, 3, 4, 5, 6}

	for i, v := range slice { //The (v) is a copy value modifying this value will not modify the value in the original built-in type on which we are looping with for range loop
		fmt.Println(i, v)
	}

	//LABELLING FOR STATEMENT
	//Labelling is indented at the left most side of the enclosing function or the block
	//labelling is used with the for loops with the break and continue keywords
outer:
	for i := 0; i < 10; i++ {
		for j := 10; j > 0; j-- {
			fmt.Println("Printing for outer label", i, j)
			break outer //This line will simply take the control back to line 100 where outer label is placed and then the code that is there after the loop will start executing
		}
	}

}
