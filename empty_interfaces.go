package main

import (
	"fmt"
	"io/ioutil" //These packages are imported from io and encoding respectively no direct path
	"encoding/json"
)

//ANCHOR - EMPTY INTERFACE IN GOLANG

// NOTE - 1. empty interface is used at the places where the type of data is not known or is upto the user that they can input anything
// like we're reading a sample json file below
// 2. the second use can be if we want to store data in a user created datastructure
// like if we implement a linked list and we want it to store any type of data in it's value field then we can use interface{}
// 3. if the value is stored in the empty interface and we want to read it back then we use type assertion and type switches

func main() {
	// the second {} is for the empty map
	data := map[string]interface{}{}
	file, err := ioutil.ReadFile("./sample.json")
	if err != nil {
		return
	}

	//defer file.Close() // In later go releases this is not used at the time of book it was needed

	json.Unmarshal(file, &data)

	fmt.Println(data)
}
