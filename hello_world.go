// every go program must have a package main from where the compiler starts building and runtime starts executing
package main

import "fmt"

// main function is the beginning point of the go program
func main() {
	fmt.Println("Hello World!")
}

// go run file_name.go this command tells the compiler to compile create a temporary binary and then delete it once the work is done
// go build file_name.go this command tells the compiler to build a binary and save it.
// go build -o ./ch1/hello ./ch1/hello_world.go this way we can supply a custom output name for the binary.
// go fmt ./ch1/hello_world.go is used to format the go lang file
// goimports can also be downloaded and used for formatting the files and imports in the specified directory and all of it's subdirectories

// while we don't insert semicolons ourselves but the golang compiler inserts them for us so (go fmt) won't fix the format if we don't keep in mind the rules specified at https://go.dev/doc/effective_go#semicolons
// for idiomatic golang conventions and all refer to https://go.dev/wiki/CodeReviewComments