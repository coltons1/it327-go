package Tree

import (
	"errors"
	"fmt"
)

//--------+--------//
//  BST Tree Node  //
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

func (node *Node) GetData() (int, error) {
	if node == nil {
		return 0, errors.New("Node is nil.")
	}
	retData := node.data
	return retData, nil
}

//-------------+-------------//
//     Binary Search Tree    //
//-------------+-------------//

type BST struct { // do we want to make it w generics?
	root     *Node
	treeSize int
}

func NewTree() *BST {
	return &BST{
		root:     nil,
		treeSize: 0,
	}
}

func (tree *BST) Insert(data int) {
	var node *Node = newNode(data)
	tree.root, _ = tree.insertHelper(tree.root, node)
}

func (tree *BST) insertHelper(currNode *Node, newNode *Node) (*Node, error) {
	// if the current node is nil, then set it there
	if currNode == nil {
		tree.treeSize++
		return newNode, nil
	} else {
		//tree's root is not null so continue down the tree
		if newNode.data >= currNode.data {
			// go to the right
			currNode.right, _ = tree.insertHelper(currNode.right, newNode)
		} else if newNode.data < currNode.data {
			// go to the left
			currNode.left, _ = tree.insertHelper(currNode.left, newNode)
		}
	}
	return currNode, nil
}

func (tree *BST) Remove(data int) (int, error) {
	// call remove helper
	newRoot, err := tree.removeHelper(tree.root, data)
	// if the removehelper errors out, fail
	if err != nil {
		return 0, err
	}
	// if it did not error out, return the removed data
	tree.root = newRoot
	tree.treeSize--
	return data, nil
}

func (tree *BST) removeHelper(node *Node, dataToRemove int) (*Node, error) {
	// cannot find the value to be removed
	if node == nil {
		return nil, errors.New("Value not found in tree.")
	}

	// searching for the piece of data to be removed
	var err error
	if dataToRemove < node.data {
		node.left, err = tree.removeHelper(node.left, dataToRemove)
		if err != nil {
			return nil, err
		}
	} else if node.data < dataToRemove {
		node.right, err = tree.removeHelper(node.right, dataToRemove)
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
		node.right, err = tree.removeHelper(node.right, curr.data)
		if err != nil {
			return nil, err
		}
	}
	return node, nil
}

func (tree *BST) Get(targetData int) (*Node, error) {
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

func (tree *BST) HasKey(targetData int) bool {
	// traverse through three until key is found, if key is not found then return false.
	var currNode *Node = tree.root
	for currNode != nil {
		if currNode.data == targetData {
			return true
		}
		if targetData < currNode.data {
			// go to the left
			currNode = currNode.left
		} else if targetData > currNode.data {
			currNode = currNode.right
		}
	}
	return false
}

func (tree *BST) Size() int {
	return tree.treeSize
}

func (tree *BST) Height() int {
	return tree.height(tree.root)
}

func (tree *BST) height(node *Node) int {
	if node == nil {
		return -1
	}

	var l int = tree.height(node.left)
	var r int = tree.height(node.right)

	if l > r {
		return 1 + l
	} else {
		return 1 + r
	}
}

func (tree *BST) MaxKey() (int, error) {
	node := tree.Max()
	if node == nil {
		return 0, errors.New("Tree is empty.")
	}
	return node.data, nil

}

func (tree *BST) MinKey() (int, error) {
	node := tree.Min()
	if node == nil {
		return 0, errors.New("Tree is empty.")
	}
	return node.data, nil
}

func (tree *BST) Max() *Node {
	if tree.root == nil {
		return nil
	}

	var currNode *Node = tree.root
	for currNode.right != nil {
		currNode = currNode.right
	}

	return currNode
}

func (tree *BST) Min() *Node {
	if tree.root == nil {
		return nil
	}

	var currNode *Node = tree.root
	for currNode.left != nil {
		currNode = currNode.left
	}

	return currNode
}

func (tree *BST) inorderPrint(node *Node) {
	if node == nil {
		return
	}
	tree.inorderPrint(node.left)
	fmt.Printf("%v\t", node.data)
	tree.inorderPrint(node.right)
}

func (tree *BST) PrintTree() {
	tree.inorderPrint(tree.root)
}

func (tree *BST) InorderArr() (arr []*Node) {
	return tree.inorder(tree.root)
}

func (tree *BST) inorder(node *Node) (arr []*Node) {
	if node == nil {
		return
	}
	tree.inorder(node.left)
	arr = append(arr, node)
	tree.inorder(node.right)
	return arr
}
