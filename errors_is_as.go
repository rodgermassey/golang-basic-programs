package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

// ANCHOR - ERRORS.IS and ERRORS.AS
// 1. If sentinel errors are wrapped they cannot be matched exactly
// 2. If errors are wrapped then type assertion, type switches and unwrap cannot find the exact error type if we're using the custom error type
// 3. errors.Is is used to check for sentinel errors for the values or exact instances and errors.As is used to check for custom errors or type of the error

/* Use errors.Is when you are looking for a specific instance or specific values. Use errors.As when you are looking for a specific type.
*/

// ANCHOR - ERRORS.IS
// 1. To check if the returned error or any errors wrapped in this error was a sentinel error use errors.Is
// 2. errorsfunc (ce CustomErr) Is(sentinelErr error) bool {
// 	if se, ok := sentinelErr.(ce); ok {
// 		return reflect.DeepEqual(ce, se)
// 	}
// 	return false
// }.Is takes 2 parameters, first error that is being checked, second the instance against which we're comparing
// 3. errors.Is returns true if the error or any error in that error chain matches the sentinel error provided
// 4. implement Is method if using a custom error type

type CustomErr struct {
	Codes []int
}

func (ce CustomErr) Error() string {
	return fmt.Sprintf("codes: %v", ce.Codes)
}

func (ce CustomErr) Is(sentinelErr error) bool {
	if se, ok := sentinelErr.(CustomErr); ok {
		return reflect.DeepEqual(ce, se)
	}
	return false
}

func fileChecker(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("error occured while opening file: %w", err)
	}
	file.Close()
	return nil
}

func customErrorFileChecker(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return CustomErr{Codes: []int{1, 2, 3}}
	}
	file.Close()
	return nil
}

// ANCHOR - ERROR.AS
// 1. errors.As checks if a returned error or any error it wraps matches a particular type
// 2. errors.As takes 2 parameters, first error that is being checked, second a pointer to the variable of the type that we inted to check or a pointer to an interface that the error should meet
// 3. if function returns true, then it means the type matches some error in the chain, if false no match of that type of error was found
// 4. just errors.Is we can also override errors.As using reflect for our custom errors

func main() {
	err := fileChecker("non-existing-file.txt")
	if err != nil {
		if errorExists := errors.Is(err, os.ErrNotExist); errorExists {
			fmt.Println(err)
		}
	}

	err = customErrorFileChecker("non-existing-file.txt")
	if err != nil {
		var customErr CustomErr
		if errorTypeExists := errors.As(err, &customErr); errorTypeExists {
			fmt.Println(err, errorTypeExists)
		}
	}
}
