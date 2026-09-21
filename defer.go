package main

import (
	"io"
	"log"
	"os"
	"fmt"
)

//ANCHOR - DEFER in GOLANG
//defer is a keyword in golang that is used inside a function
//defer is followed by a function call
//this deferred function call is run at the end of the function exit,
//meaning after the return statement of the function assigns the values to the return variables it executes the deferred function call
//1. we can provide multiple defer statments in a function, they run in last in first out order
//2. if a value is returned from the function that is supplied to the defer there is no way to read that value
//3. defer can overwrite the values of the named return variables of the function, this pattern is used to add contextual info about the errors

// NOTE - defer is used to run the cleanup functions or any job that has to be closed before the function returns
func main() {
	if len(os.Args) < 2 {
		//No need to run return after log.Fatal, it automatically exits the program
		log.Fatal("No file name specified")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	//ANCHOR - in our clean up function we are calling f.close() because
	// we want any file that was opened by this function should be closed
	// before this function returns
	defer f.Close()

	data := make([]byte, 2048)

	for {
		//NOTE - f.Read will only read the data upto the length of the slice that we supply to it
		// count is the number of bytes that has been read by the f.Read
		count, err := f.Read(data)
		//NOTE - os.Stdout.Write will write the data from this slice which is starting from 0 till count-1
		os.Stdout.Write(data[:count])

		if err != nil {
			//NOTE - this condition checks whether the error is END OF FILE or not
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}

	file, closer, err := getFile(os.Args[1])
	if err != nil {
		//NOTE - here we call log.Fatal as this is the function that calls getFile
		log.Fatal(err)
	}

	fmt.Println("File Descriptor:", file.Fd())
	fmt.Println("File Name:", file.Name())
	defer closer()
}

// NOTE - In go it is idiomatic that if a function allocates some resources
// it should also clear those resources, by supplying a closure
func getFile(fileName string) (*os.File, func(), error) {
	file, err := os.Open(fileName)

	if err != nil {
		//NOTE - We'll not use log.Fatal in this function we'll pass the error to the parent function
		//and the topmost function will run print all the error logs for entire error stack
		//log.Fatal(err)
		return nil, nil, err
	}

	return file, func() {
		file.Close()
	}, nil
}
