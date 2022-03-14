package binary_search_tree

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func (node *Node) search(value int) *int {
	if node == nil {
		return nil
	}
	if node.value == value {
		return &node.value
	}
	if node.value < value {
		return node.left.search(value)
	} else {
		return node.right.search(value)
	}
}

func (node *Node) insert(value int) {
	if node.value <= value {
		if node.left == nil {
			node.left = &Node{
				value: value,
			}
			return
		}
		node.left.insert(value)
	} else {
		if node.right == nil {
			node.right = &Node{
				value: value,
			}
			return
		}
		node.right.insert(value)
	}
}

func (node *Node) print() {
	if node == nil {
		return
	}
	node.left.print()
	fmt.Println(node.value)
	node.right.print()
}

type Tree struct {
	root *Node
}

func New() *Tree {
	return &Tree{}
}

func (tree *Tree) Insert(value int) {
	if tree.root == nil {
		tree.root = &Node{
			value: value,
		}
		return
	}

	tree.root.insert(value)
}

func (tree *Tree) Search(value int) *int {
	return tree.root.search(value)
}

func (tree *Tree) Print() {
	if tree.root == nil {
		return
	}
	tree.root.print()
}
