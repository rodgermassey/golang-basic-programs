package main

import (
	"fmt"
)

// ANCHOR - INTERFACES IN GOLANG
// 1. Interfaces is the only abstract type in go, interface's names usually ends with 'er' (Stringer, io.Reader, io.Closer, etc.)
// 2. Methods defined by interface are called to be in the method set of the interface
// 3. In go interfaces are implicit we don't have to implement an interface if a type has the methods defined in the methodset of the interface then that
// type automatically satisifes that interface
// 4. Interfaces specify what callers need. The client code defines the interface to specify what functionality it requires.
// 5. Golang encourages decorator pattern, it promots factory functions to take in an instance of an interface and return another type that implements the same interface
// e.g. os.Open(fileName) returns an instance that satisfies the io.Reader interface
// 6. Interfaces can also be embedded in the structs.
type Logicprovider struct{}

// NOTE - no implementation relation between this method and interface type, it just simply creates a method without any knowledge of interface,
// unlike java where we have to tell what interface to implement
func (lp Logicprovider) Process(data string) string {
	return fmt.Sprintf("The input data is: %s", data)
}

// NOTE - Logic interface simply defines the method set
type Logic interface {
	Process(data string) string
}

// NOTE - Only the caller i.e. (Client)  knows the interface
// 1. there is nothing on LogicProvider to tell which interface it implements
// 2. It also ensures that in future a new logic provider can be passed in as long as it matches the client's needs
type Client struct {
	L Logic
}

func (c Client) Program() string {
	return c.L.Process("hello")
}

type embeddingStructAndInterfaces struct {
	Logic //NOTE - here the user will have to provide the struct value for logic such that the struct should satisfy this interface
}

func main() {
	c := Client{
		L: Logicprovider{},
	}

	fmt.Println(c.Program())
}
