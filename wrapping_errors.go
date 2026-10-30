package main

import (
	"errors"
	"fmt"
	"os"
)

// ANCHOR - WRAPPING ERRORS
// 1. while passing back errors through code we want some additional info (e.g. function name where error occured, operation trying to perform)
// such information can be added to the error while preserving the error it's called error wrapping
// 2. series of wrapped errors is called error chain.
// 3. fmt.Errorf uses %w verb to wrap the error, convention is to use (: %w) at the end of the error format string,
// and pass the error being wrapped as the 2nd parameter of the fmt.Errorf
// 4. to unwrap the errors, to get the error that was wrapped, std errors library gives us unwrap function,
// if there was no error was wrapped it returns nil
// 5. Don't usually call errors.Unwrap() we use erros.Is or errors.As to find a specific wrapped error
// 6. we can implement Unwrap method on our custom error type as well, it just returns the error, and now we can use this custom type to wrap errors

// NOTE - if we want to create a new error and add info from the old error rather than wrapping it 
// we use %v verb to append the error string to our error, rather than wrapping the error

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
	Err   error
}

func (se StatusErr) Error() string {
	return se.Message
}

func (se StatusErr) Unwrap () error {
	return se.Err
}

func fileChecker(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()
	return nil
}

func fileCheckerCustomErr(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return StatusErr{
			Status: NotFound,
			Message: "error opening file",
			Err: err,
		}
	}
	defer file.Close()
	return nil
}


func main() {
	err := fileChecker("./no_file.txt")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}

	customError := fileCheckerCustomErr("./no_file.txt")
	if customError != nil {
		fmt.Println(customError)
		if wrappedErr := errors.Unwrap(customError); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}
}
