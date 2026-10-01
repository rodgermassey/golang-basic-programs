package main

import (
	"fmt"
)

//ANCHOR - CODING METHODS TO HANDLE NIL INSTANCES
// 1. The value reciever methods cannot handle nil instances
// and if called with a nil instance then they will panic
// 2. The pointer reciever methods don't panic and nil instances can invoke pointer reciever methods
// 3. But we have to handle pointer reciever methods exclusively for nil instances

// NOTE - implementing binary tree for integers
type Node struct {
	Value               int
	LeftNode, RightNode *Node
}

//NOTE - Both the insert and the contains method are pointer recievers 
// and handle the nil instances, 
// this is how we can utilize pointer reciever methods for nil instances
// if this was not there then the Insert would throw an error
func (n *Node) Insert(value int) *Node {
	// TBC from line 29
	// when the Insert is triggered from n.RightNode or n.LeftNode that doesn't exist
	// then n will be nil or for the first element n will be nil
	// then we insert create a node and return the node's address
	// (refer line 33 after this)
	if n == nil {
		return &Node{Value: value}
	}

	// if the value is bigger that the value at node then we move to the right subtree
	// and pass the value recursively until we find that the n.RightNode is nil
	// at that base case we trigger insert again (refer line 29 comment after this)
	// now the node is returned and is placed at n.RightNode which was previously nil
	// (refer line 41)
	if value > n.Value {
		n.RightNode = n.RightNode.Insert(value)
	} else if value < n.Value {
		n.LeftNode = n.LeftNode.Insert(value)
	}

	// now we return this node which has the updated right child and then this goes in the loop
	// code from here either exits or goes to either line 36 or line 38
	// and thus updates there left or right node with the updated left or right node respectively
	// which was returned from the neseted insert
	return n
}

func (n *Node) Contains(value int) bool {
	switch {
	case n == nil:
		return false
	case value > n.Value:
		return n.RightNode.Contains(value)
	case value < n.Value:
		return n.LeftNode.Contains(value)
	default:
		return true
	}
}

func main() {
	var binaryTree *Node
	//NOTE - we are able to call Insert on the binaryTree 
	// which is nil pointer for *Node type 
	// because pointer reciever methods can be called with nil instances
	// but value reciever methods cannot they would panic 
	binaryTree = binaryTree.Insert(5)
	binaryTree = binaryTree.Insert(3)
	binaryTree = binaryTree.Insert(6)
	binaryTree = binaryTree.Insert(8)

	fmt.Println("Checking for 12:", binaryTree.Contains(12))
	fmt.Println("Checking for 5:", binaryTree.Contains(5))
	fmt.Println("Checking for 3:", binaryTree.Contains(3))
	fmt.Println("Checking for 8:", binaryTree.Contains(8))
	fmt.Println("Checking for 6:", binaryTree.Contains(6))
}
