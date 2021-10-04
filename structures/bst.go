package main

import "fmt"

type Leaf struct {
	data  int
	left  *Leaf
	right *Leaf
}

type Tree struct {
	root *Leaf
	size int
}

// Constructor for the tree
func createTree() *Tree {
	return &Tree{
		root: nil,
		size: 0,
	}
}

// Method to insert a new leaf
func (t *Tree) insert(data int) {
	t.size += 1
	newLeaf := &Leaf{
		data:  data,
		left:  nil,
		right: nil,
	}

	if t.root == nil {
		t.root = newLeaf
	} else {
		curr := t.root
		for {
			if data < curr.data {
				if curr.left == nil {
					curr.left = newLeaf
					return
				}
				curr = curr.left
			} else {
				if curr.right == nil {
					curr.right = newLeaf
					return
				}
				curr = curr.right
			}
		}
	}
}

// Find a specific leaf in the tree
func (t *Tree) search(data int) *Leaf {
	curr := t.root

	for curr.data != data {
		if curr != nil {
			if curr.data > data {
				curr = curr.left
			} else {
				curr = curr.right
			}
			if curr == nil {
				return nil
			}
		}
	}
	return curr
}

// In order traversal Left -> Root -> Right
func inorder(root *Leaf) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Println(root.data)
	inorder(root.right)
}

// Pre order traversal Root -> Left -> Right
func preorder(root *Leaf) {
	if root == nil {
		return
	}
	fmt.Println(root.data)
	preorder(root.left)
	preorder(root.right)
}

// Post order traversal Left -> Right -> Root
func postorder(root *Leaf) {
	if root == nil {
		return
	}
	postorder(root.left)
	postorder(root.right)
	fmt.Println(root.data)
}

// Main entry point
func main() {

	tree := createTree()
	values := []int{2, 6, 3, 18, 9, 5, 21}

	for _, val := range values {
		tree.insert(val)
	}

	fmt.Println(tree.search(18))
	fmt.Printf("\n=================================\n\n")
	inorder(tree.root)
	fmt.Printf("\n=================================\n\n")
	preorder(tree.root)
	fmt.Printf("\n=================================\n\n")
	postorder(tree.root)
}
