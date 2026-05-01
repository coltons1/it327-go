package main

import (
	"Substantial/Tree"
	"fmt"
)

func main() {
	fmt.Print("Hello!")
	var tree *Tree.BST = Tree.NewTree()
	tree.Insert(10)
	tree.Insert(928)
	tree.Insert(-1)
	tree.Insert(74)
	tree.Insert(8)
	tree.Insert(210)
	tree.PrintTree()
}
