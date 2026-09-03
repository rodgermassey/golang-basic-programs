package main

import (
	"fmt"
)

// SWITCH STATEMENT IN GOLANG
// we can switch values of all built in types except slices, maps, channels, functions and structs which contain fields of these types
// we can switch all the values that can be compared using == operator
func main() {

	spellings := []string{"a", "cow", "elephant", "smile", "gopher", "octopus", "anthropologist"}
	//No break statements are required between cases
	//ANCHOR - BREAK statement can be used if we want to exit any case early and exit the entire switch block as well
	//Golang does include a keyword fallthrough which lets one case continue on to the next one but think twice before implementing it that way it might not be idiomatic

	//ANCHOR - EXPRESSION SWITCH STATEMENT
	for _, word := range spellings {
		switch size := len(word); size { // No paranthesis around the value being compared i.e. size
		case 1, 2, 3, 4: // we can group multiple cases together, by comma separated values
			fmt.Println(word, "is a short word")
		case 5: // we can see that no curly braces are required to group multiple statements in case block wordlen is also case block scoped
			wordlen := len(word)
			fmt.Println(word, " is exactly of length", wordlen)
		case 6, 7, 8, 9: // if there is an empty case we simply leave that case block empty
		default: // default case (optional)
			fmt.Println(word, "is a long word")
		}
	}

	//ANCHOR - BLANK SWITCH STATEMENT
	//when we don't specify the value against which we are switching, golang compares the cases on the basis boolean comparision for each case
	spellings = []string{"hi", "salutations", "hello"}
	for _, word := range spellings {
		switch wordlen := len(word); { // here we've skipped the value against which we are comparing, therefore it is called blank switch statement
		case wordlen < 5:
			fmt.Println(word, " is a short word")
		case wordlen > 10:
			fmt.Println(word, " is a long word")
		default:
			fmt.Println(word, " is exactly the right length")
		}
	}

	// If you find that you have written a blank switch where all of your cases are equality comparisons against the same variable
	// then it is better to replace it with the expression switch statement and use the values as the cases

	/*Favor blank switch statements over if/else chains when you have
	  multiple related cases. Using a switch makes the comparisons
	  more visible and reinforces that they are a related set of concerns.
	*/
}
