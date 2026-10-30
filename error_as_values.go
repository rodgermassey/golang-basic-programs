package main

import (
	"fmt"
)

// ANCHOR - Errors as values
// 1. since error is an interface, we can have our custom type for error that implements the error interface
// 2. even with custom error type always return error from your functions 
// so that it stays generic rather than adhereing to only a single type of error
// 3. don't return the uninitialized instance of the error type from the function 
// as the error is an interface and unititialized instance of our custom type will have the type field set
// 4. always return nil if the error is not to be returned, from the function rather than uninitialized instance of the error type
// 5. either return nil from the function or the instance of error interface type

/* When using custom errors, never define a variable to be of the type
   of your custom error. Either explicitly return nil when no error
   occurs or define the variable to be of type error.
*/

// NOTE - example
// 1. Here we declare a custom type Status based on int
// 2. we use it as an enum for different status values in the const block
// 3. we define our custom error struct in this case StatusErr
// 4. we define a Error method on our StatusErr struct to satisfy the error interface

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	status  Status
	message string
}

func (se StatusErr) Error() string {
	return se.message
}

// using this function to throw an error
func throwInvalidLoginError() (int, error) {
	return 0, StatusErr{
		status:  InvalidLogin,
		message: "The user is invalid",
	}
}

func main() {
	if _, err := throwInvalidLoginError(); err != nil {
		fmt.Println(err)
	}
}
