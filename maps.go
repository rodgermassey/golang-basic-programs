package main

import (
	"fmt"
)

// MAPS in go are used to store key value pairs like objects in js
// MAPS are not comparable using == or != they can only be compared against nil
// The key of the map has to be a comparable type therefore slice and map cannot be used as the keys of the map
// ORDER OF KEYS doesn't matter in maps
func main() {
	//ZERO VALUE IS NIL MAP of a MAP
	//NIL MAP LENGTH = 0
	//READING NIL return NIL VALUE OF MAP
	//WRITING TO NIL MAP PANICS
	var a map[string]int //this creates a nil map, the keys will be string values and the values will integer values

	//EMPTY MAP LITERAL
	//EMPTY MAP LENGTH = 0
	//READ AND WRITE IS POSSIBLE ON EMPTY MAP
	var b = map[string]int{}

	//NON EMPTY MAP
	var c = map[string][]string{
		"cars":  []string{"BMW", "NISSAN", "TOYOTA"},
		"bikes": []string{"HERO", "JAWA", "TVS"},
	}
	fmt.Println(a, b, c, len(c))

	//USING MAKE to create a map
	//USE it when we know only the number of key value pairs not the exact values
	//MAPS created with MAKE has initial length = 0, and they can also outgrow the specified size
	d := make(map[string]int, 10) // 10 is the size, length = 0, can outgrow the size i.e. 10
	fmt.Println(d, len(d))

	//READING AND WRITING TO A MAP
	e := map[string]int{}
	e["JAVA"] = 5
	e["JAVASCRIPT"] = 8
	fmt.Println(e["JAVA"])
	fmt.Println(e["GOLANG"]) //Returns the 0 value for the value type of the map i.e. int in our case
	fmt.Println(e)
	e["GOLANG"]++
	fmt.Println(e["JAVASCRIPT"])
	fmt.Println(e["GOLANG"])

	//COMMA-OK IDIOM
	//if ok is true then the key is present
	//Used to find whether a key is present in the map or not, rather than judging based on the 0 value of the value type\
	//v and ok variables are not block scoped as they are not enclosed in the if block as the variables value and ok on line number 72
	v, ok := e["JAVA"]
	fmt.Println(v, ok) // v prints the value, ok stores the boolean value based on whether the key is present or not

	v, ok = e["RUST"]
	fmt.Println(v, ok)

	//DELETING KEYS FROM MAP
	delete(e, "JAVA")
	fmt.Println(e)

	//USING MAPS AS SETS
	//Golang doesn't have inbuilt set implementation but maps can be used to convert slices into sets and then can be used for quick-lookup
	slice := []int{1, 2, 3, 4, 5, 6, 1, 2, 3}
	var m = map[int]struct{}{}
	for _, v := range slice {
		m[v] = struct{}{}
	}

	fmt.Println(len(slice), len(m))

	// this value and ok variables are block scoped therefore we can use them again as we're doing on line 75
	if value, ok := m[100]; !ok {
		fmt.Println("100 doesn't exist in set",value)
	}

	if value, ok := m[2]; ok {
		fmt.Println("2 exists in set", value)
	}

	fmt.Println("querying m[100]", m[100]) // returns {} because this is the 0 value for a struct
	fmt.Println("querying m[2]", m[2]) // also returns {} because we are storing the 0 value in the map for this key,
	//therefore we should always use comma,ok idiom while checking for keys in the map
