package main

import (
	"fmt"
)

//ANCHOR - Pointers in golang holds the addresses of the memory locations
//1. the 0 value for a pointer is nil
//2. maps, slices, functions, interfaces and channels are of pointer types
//3. to create a pointer variable we use (*) indirection operator in front of the type e.g. var x *string, var y *int
//4. to access the address of a value we use (&) address of operator in front of value
//5. DEREFERENCING => to access the value stored at a particular address use indirection operator e.g. x = &y; to access the value of y we can say *x = y because x stores the address to the value stored at y
//6. while dereferncing keep in mind, nil pointers cannot be dereferenced, if the program tries to dereference the nil pointer it will cause a panic
//7. Pointers indicate that a parameter is mutable
//8. Pointers as parameters use dereferencing to update the value
//9. Pointers can be used to detect 0 values vs no values for the basic types, as empty pointers point to nil rather than 0 value
//	 create a function that return comma, ok idiom like values for checking this,
//	 also be careful as pointers make values mutable
//10.when working with JSON, use pointer values for the fields of struct that can be nullable

func main() {
	x := 10
	y := &x //here we store the address of the variable x into the variable y using &x address of operator
	fmt.Println(x, y, *y) //here *y is dereferencing to get the value stored at the address that is stored in y

	var z *string //here we create a nil pointer, *string shows a pointer type
	fmt.Println(z == nil, z)

	//ANCHOR - new function
	//the built-in new function creates a pointer variable, 
	//it creates a pointer to a zero value of the provided type

	a := new(int)//here new creates a pointer to a value of type int
	fmt.Println(a, *a)//here we can see that the value stored at that address is 0 value of type int i.e. 0

	//NOTE - to take the address or the pointer to the struct add & before the struct literal
	//this doesn't work with primitive literals like number, string and booleans
	b := &struct{a string}{a: "hello"}
	//c := &"hello" - this will not work as we cannot assign the address to the variable directly as we do with structs with primitive literals
	//NOTE - use fmt.Printf to print the key names as well for the struct
	fmt.Printf("%+v\n", b)//This prints the key names and the values
	fmt.Println(*b, b)//printing a struct doesn't print the key names to print keynames we have to use fmt.Printf("%v")

	//if there is a struct field that takes the pointer to a primitive value then we can create a function that returns the pointers to these values
	type Person struct {
		firstname string
		lastname *string
	}

	p := Person{
		firstname: "Mike",
		//lastname: "Tyson", this will not compile as lastname wants a pointer to a string
		lastname: stringP("Tyson"),//this is how we handle this situation
	}
	fmt.Printf("%+v\n", p)

	d := 10

	fmt.Println(d)
	failedUpdate(&d) //doesn't update as the address is updated not the value at that address
	fmt.Println("Failed update",d) //10

	update(&d)
	fmt.Println("Successful update", d) //20
}

//NOTE - this works because when a string or a constant literal is passed to this function,
//it is copied to the parameter variable which has a address
//therefore this address can be returned
func stringP(s string) (*string) {
	return &s
}


//ANCHOR - POINTERS are MUTABLE when passed as parameters
//below is the incorrect way to update the value of the pointer as we're updating px
//i.e. is the variable that holds the pointer, not the value that the pointer points to
func failedUpdate (px *int) {
	x := 20
	px = &x
}

//below is the correct way to update the value of the pointer as we're updating the value
//i.e. is stored at the address that is being pointed towards by this pointer
func update (px *int) {
	x := 20
	*px = x
}
