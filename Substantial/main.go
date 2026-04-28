package main

import (
	"fmt"
)

//--------+--------//
//  AVL Tree Node  //
//--------+--------//

type Node struct {
	key         int
	height      int
	data        int
	left, right *Node
}

func newNode(newKey int, newData int) *Node {
	return &Node{
		key:   newKey,
		data:  newData,
		left:  nil,
		right: nil,
	}
}

func updateHeight(node *Node) int {
	var l int = 0
	var r int = 0
	var height int = 0
	if node.left == nil {
		l = 0
	} else {
		l = node.left.height
	}

	if node.right == nil {
		r = 0
	} else {
		r = node.right.height
	}

	if l > r {
		height = 1 + l
	} else {
		height = 1 + r
	}
	return height
}

//--------+--------//
//     AVL Tree    //
//--------+--------//

type AVLTree struct { // do we want to make it w generics?
	root *Node

	newTree func() *AVLTree
	insert  func(*AVLTree, int, int)
	remove  func(*AVLTree, int)
	get     func(*AVLTree, int) int
	hasKey  func(*AVLTree, int) bool
	size    func(*AVLTree) int
	height  func(*AVLTree) int
	maxKey  func(*AVLTree) int
	minKey  func(*AVLTree) int
	max     func(*AVLTree) *Node
	min     func(*AVLTree) *Node
}

func newTree() *AVLTree {
	return &AVLTree{
		root: nil,
	}
}

func insert(tree *AVLTree, key int, data int) {

}

func remove(tree *AVLTree, key int) {

}

func get(tree *AVLTree, key int) int {
	return 0
}

func hasKey(tree *AVLTree, key int) bool {
	return false
}

func size(tree *AVLTree) int {
	return 0
}

func height(tree *AVLTree) int {
	return 0
}

func maxKey(tree *AVLTree) int {
	return 0
}

func minKey(tree *AVLTree) int {
	return 0
}

func max(tree *AVLTree) *Node {
	return nil
}

func min(tree *AVLTree) *Node {
	return nil
}

func main() {
	fmt.Println("Hello World")
}
