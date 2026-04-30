package main

import (
	"errors"
	"fmt"
)

//--------+--------//
//  AVL Tree Node  //
//--------+--------//

type Node struct {
	data        int
	left, right *Node
}

func newNode(newData int) *Node {
	return &Node{
		data:  newData,
		left:  nil,
		right: nil,
	}
}

//-------------+-------------//
//     Binary Search Tree    //
//-------------+-------------//

type BST struct { // do we want to make it w generics?
	root *Node

	NewTree      func() *BST
	Insert       func(*BST, int, int)
	insertHelper func(*BST, *Node, *Node) (*Node, error)
	Remove       func(*BST, int)
	Get          func(*BST, int) int
	size         func(*BST) int
	height       func(*BST) int
	max          func(*BST) *Node
	min          func(*BST) *Node
}

func NewTree() *BST {
	return &BST{
		root: nil,
	}
}

func Insert(tree *BST, data int) {
	var node *Node = newNode(data)
	tree.root, _ = insertHelper(tree, nil, node)
}

func insertHelper(tree *BST, currNode *Node, newNode *Node) (*Node, error) {
	// if the root of the tree is nil, set the tree to the new node
	if tree.root == nil {
		tree.root = newNode
		return tree.root, nil
	}
	// if the current node is nil, then set it there
	if currNode == nil {
		currNode = newNode
		return currNode, nil
	} else {
		//tree's root is not null so continue down the tree
		if newNode.data >= currNode.data {
			// go to the right
			insertHelper(tree, currNode.right, newNode)
		} else if newNode.data < currNode.data {
			// go to the left
			insertHelper(tree, currNode.left, newNode)
		}
	}
	return newNode, errors.New("Insertion failed, could not find proper placement.")
}

func Remove(tree *BST, data int) (int, error) {
	// call remove helper
	newRoot, err := removeHelper(tree, tree.root, data)
	// if the removehelper errors out, fail
	if err != nil {
		return 0, err
	}
	// if it did not error out, return the removed data
	tree.root = newRoot
	return data, nil
}

func removeHelper(tree *BST, node *Node, dataToRemove int) (*Node, error) {
	// cannot find the value to be removed
	if node == nil {
		return nil, errors.New("Value not found in tree.")
	}

	// searching for the piece of data to be removed
	var err error
	if dataToRemove < node.data {
		node.left, err = removeHelper(tree, node.left, dataToRemove)
		if err != nil {
			return nil, err
		}
	} else if node.data < dataToRemove {
		node.right, err = removeHelper(tree, node.right, dataToRemove)
		if err != nil {
			return nil, err
		}
	} else {
		// node with 0 or 1 child
		if node.left == nil {
			return node.right, nil
		}
		if node.right == nil {
			return node.left, nil
		}

		// node with 2 children
		curr := node.right
		for curr.left != nil {
			curr = curr.left
		}
		node.data = curr.data
		node.right, err = removeHelper(tree, node.right, curr.data)
		if err != nil {
			return nil, err
		}
	}
	return node, nil
}

func Get(tree *BST, targetData int) (*Node, error) {
	if tree == nil {
		return nil, errors.New("Tree does not exist. ")
	}

	var currNode *Node = tree.root
	for currNode != nil {
		if currNode.data == targetData {
			return currNode, nil
		}
		if targetData < currNode.data {
			// go to the left
			currNode = currNode.left
		} else if targetData > currNode.data {
			// go to the right
			currNode = currNode.right
		}
	}
	return nil, errors.New("Could not find node with that data")
}

func hasKey(tree *BST, key int) bool {
	return false
}

func size(tree *BST) int {
	return 0
}

func height(tree *BST) int {
	return 0
}

func maxKey(tree *BST) int {
	return 0
}

func minKey(tree *BST) int {
	return 0
}

func max(tree *BST) *Node {
	return nil
}

func min(tree *BST) *Node {
	return nil
}

func main() {
	fmt.Println("Hello World")
}
