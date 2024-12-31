package data

type Node struct {
	Value any
	Left  *Node
	Right *Node
}

// This is a pretty "dumb" implementation of a binary tree. If I were going to really invest in this I would probably create pointer receiver
// functions to add a new value and it get inserted in the correct place as well as functions for removing data, resetting, etc.
type Btree struct {
	Root *Node
}
