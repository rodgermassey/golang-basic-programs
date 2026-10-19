package main

import (
	"fmt"
)

// ANCHOR - INTERFACES AND HOW THEY HANDLE NIL
// 1. interfaces have 2 pointer values under the hood one is pointer to the value type and one is the pointer to the value
// 2. an interface is a nil interface only when it's both value type and the value pointer are nil
// 3. only the value type cannot be nil as there cannot be a value in golang without a type
// 4. value pointer can be nil depending upon the value type if the 0 value of that type is nil then it can be nil
// 5. as we can invoke methods on nil concrete instances (nil structs) similarly we can invoke methods on interfaces which are assigned nil concrete instances
// 6. if an interface is nil invoking any method on that will cause a panic
// 7. if an interface is non-nil we can invoke methods on it, 
// But note that if the value is nil and the methods of
// the assigned type don’t properly handle nil, you could still trigger a panic.
func main() {

	var s *string
	var i interface{}

	fmt.Println(s == nil)
	fmt.Println(i == nil)
	i = s // NOTE - here when we assign the value s to the i, but s has the type *string and its value is nil,
	// so when we assign this to the empty interface i both the type and the value of s is copied thus the value type becomes *string and the value becomes nil
	// for an interface to be nil both should be nil
	fmt.Println(i == nil)
}
