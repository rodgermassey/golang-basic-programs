package main

import (
	"fmt"
	"time"
)

//ANCHOR - Methods and types in golang
//1. types in golang can be package scoped or the block scoped
//2. methods can only be package scoped, so if a type wants to have methods they need to be package scoped
//3. types can be concrete or abstract
//4. types can be created using primitive or compound types
//5. methods cannot be overloaded in go, 2 different types can have same method name but 2 methods in the same type cannot have the same name
//6. reciever name in the method is usually the first letter of the type, idiomatic go
//7. methods can only be declared in the same package as the type, though it can be a different file in the same package, but that approach too is non idiomatic
//8. Don't write getter or setter methods unless we want the struct to meet an interface.
//9. Exception to point 8 are multiple field updates as single operation or when the update is not direct, 
//the increment method on counter struct demonstrates both these properties

/*ANCHOR - METHOD SETS IN GOLANG
// 1. method sets are all the methods that are available on any struct instance
// 2. if the struct value is of value type it will have only the value reciever methods in the method set
// 3. if the value is of pointer type then the method set will have both the pointer and the value reciever methods
// 4. in normal go we don't see any difference as golang doesn't throw an error
// as long as it can find the address of the value type, 
// but while returning a struct from a function and then calling a method on it might cause an error 
// e.g. getUser().setName() where getUser returns a value instance of the user struct
*/

/*NOTE - POINTS FOR DETERMINING WHEN TO USE POINTER AND WHEN TO USE VALUE RECEIVERS
// Pointer receiver:
// 1. Use a pointer receiver if the method modifies the receiver.
// 2. Use a pointer receiver if the method needs to handle nil instances.
// 3. If the method doesn't modify the receiver, a value receiver can be used.
// 4. If the type has any pointer receiver methods, prefer pointer receivers
//    for all methods for consistency.
*/

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// NOTE - this is called reciever the parameter like variable between func keyword and the method name
func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

type Counter struct {
	Counter     int
	LastUpdated time.Time
}

// NOTE - pointer reciever method
func (c *Counter) Increment() {
	c.Counter++
	c.LastUpdated = time.Now()
}

// NOTE - value reciever method
func (c Counter) String() string {
	return fmt.Sprintf("Counter value is %d \n Updated at: %v", c.Counter, c.LastUpdated)
}

func main() {
	matt := Person{
		FirstName: "Matt",
		LastName:  "Hardy",
		Age:       29,
	}

	fmt.Println(matt.String()) //invoking a method on the instance of Person type

	counter := Counter{
		Counter:     0,
		LastUpdated: time.Now(),
	}
	fmt.Println(counter.String())
	//NOTE -even though Increment is pointer receiver method but we can invoke it directly on the instance
	//NOTE -GOLANG interally does (&counter).Increment() this
	//but the rules for the pointer parameters remain the same
	counter.Increment()
	fmt.Println(counter.String())
}
