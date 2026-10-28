package main

import (
	// "errors"
	"fmt"
)

// ANCHOR - Errors in golang
// 1. In golang there is not inbuilt mechanism to check if the error was returned, we have to add check using if conditions to know if error was returned or not
// 2. The error is an inbuilt interface in golang that defines Error method, we can create custom errors if we have this method in our error type
// e.g. type error interface {
// Error() string
// }
// 3. Use strings for simple errors, we can create error strings using errors.New or fmt.Errorf, also error is the last returned value from the function

// NOTE: Example for point 3
func divide(a, b int) (int, error) {
	if b == 0 {
		// return 0, fmt.Errorf("Cannot divide %v by 0", a) //Using this method we can format the error string as well
		return 0, errors.New("Cannot divide by 0") //This simply uses the errors.New method to create an error string
	}
	return a / b, nil
}

func main() {
	if _, err := divide(2, 0); err != nil {
		fmt.Println(err)
	}

	v, err := divide(5, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}
