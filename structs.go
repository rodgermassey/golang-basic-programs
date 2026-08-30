package main

import (
	"fmt"
)

// STRUCTS IN GOLANG
// STRUCTS are defined using the keyword type then the name of the struct and then the keyword struct
// structs can be defined inside the function or at the package level
func main() {
	type person struct {
		name   string
		age    int
		height int
	}

	//UNLIKE MAPS, STRUCTS don't have any difference between assigning an empty struct literal and not assigning a value at all, both cases will have the 0 values for all the fields of the struct
	var peter person //this defines peter variable of person type's 0 value
	//which will be the struct's each field value to be set to its type's 0 value
	// name = "", age = 0 and height = 0

	bob := person{} //struct variable can be assigned to a variable as well

	//ASSIGNING VALUES TO STRUCTS (THE 2 STYLES specified below cannot be mixed)
	//1. we can assign values directly as comma separated values, values should follow the order of struct's keys and each key must be assigned a value
	//KEY ORDER MATTERS HERE
	julia := person{
		"Julia",
		20,
		150,
	}

	//2. we can assign values with the key names, looks similar to map
	//here we use the key names and we can skip the values of the keys
	//the keys which are not assigned will have the 0 values of their respective types
	//KEY ORDER DOESN'T MATTER HERE
	beth := person{
		name:   "Beth",
		height: 148, //it is important to add a comma at the last key value pair
	}

	//If we initialize the struct without using key names and then later on add more keys to the struct then our code will no longer compile and we will have to add the newly added keys to our old struct literals

	//ACCESSING VALUES IN STRUCT
	//just like we use [] for maps for reading and writing
	//we use (.) dot notation for both reading and writing to a struct
	//values is a struct are accessed using (.) dot notation
	bob.name = "Bobby"

	fmt.Println(peter, bob, julia.age, beth.height)

	//ANONYOUS STRUCT
	//We can skip the struct name and directly use the definition of the struct
	//mainly used for translating external data to and from into a struct (marshalling) & for writing tests
	//1. we declare the variable of a particular anonymous struct and then assign values to it
	var uranus struct {
		name   string
		salary int
	}

	uranus.name = "Uranus"
	uranus.salary = 1000

	//2. we declare the struct and assign the values to the variable altogether
	hilda := struct {
		name    string
		address string
	}{
		name:    "Hilda",
		address: "Jungle",
	}

	//type of uranus and hilda is anonymous structs
	fmt.Println(uranus, hilda)

	//COMPARING AND CONVERTING STRUCTS
	//In go 2 named structs can be compared if they have the same field names, in the same order and has the same number of fields and the type of both the variables are the same struct

	type person1 struct {
		name string
		age  int
	}

	type person2 struct {
		name string
		age  int
	}

	var p1 person1
	// var p2 person2

	// fmt.Println(p1 == p2) //can't be compared as both the structs have different names

	type person3 struct {
		age  int
		name string
	}

	// var p3 person3
	// fmt.Println(p1 == p3) //can't compare as the order of fields is different

	p4 := struct {
		name string
		age  int
	}{name: "P4",
		age: 12,
	}

	p1 = person1{
		name: "P1",
		age: 12,
	}

	p4 = p1 //assignment is also possible if at least one variable's type is anonymous struct and they have same field names, order and types

	fmt.Println(p1 == p4) //comparable and assignment is possible if using anonymous struct
}
