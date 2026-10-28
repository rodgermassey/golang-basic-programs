package main

import (
	"fmt"
	"archive/zip"
	"bytes"
)

//ANCHOR - SENTINEL ERRORS
// 1. Sentinel errors are the variables declared at the package level, by convention their names start with Err,
// with exception of EOF
// 2. Sentinel errors is an idiomatic way to tell the program that either the processing cannot start at all or cannot continue beyond this point
// 3. Sentinel errors should be treated as read only, as compiler doesn't enforce that you shouldn't change the values of these
// 4. we compare the value of the error with the sentinel error like in our example we compare it with zip.ErrFormat

func main() {
	data := []byte("This is not a file")
	notAZipFile := bytes.NewReader(data)

	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err != nil {
		if err == zip.ErrFormat { // NOTE - zip.ErrFormat is a sentinel error
			fmt.Println("Not a zip file", err)
		}
	}
}
