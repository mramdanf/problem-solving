package bst

import "fmt"

type Node struct {
	Data int
	Left *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

func (b *BinarySearchTree) AddNode(node *Node, data int) *Node {
	if b.Root == nil {
		b.Root = &Node{Data: data}
		return b.Root
	}

	if node == nil {
		return &Node{Data: data}
	}

	if data <= node.Data {
		node.Left = b.AddNode(node.Left, data)
	} else if data > node.Data {
		node.Right = b.AddNode(node.Right, data)
	}


	return node
}

func (b *BinarySearchTree) InOrderTraversal(node *Node) {
	if node == nil {
		return
	}

	b.InOrderTraversal(node.Left)
	fmt.Print(node.Data, " ")
	b.InOrderTraversal(node.Right)
}

func (b *BinarySearchTree) PreOrderTraversal(node *Node) {
	if node == nil {
		return
	}

	fmt.Print(node.Data, " ")
	b.PreOrderTraversal(node.Left)
	b.PreOrderTraversal(node.Right)
}

func (b *BinarySearchTree) PostOrderTraversal(node *Node) {
	if node == nil {
		return
	}

	b.PostOrderTraversal(node.Left)
	b.PostOrderTraversal(node.Right)
	fmt.Print(node.Data, " ")
}

func (b *BinarySearchTree) Search(node *Node, data int) bool {

	if node == nil {
		return false
	}

	if node.Data == data {
		return true
	}

	if data != node.Data && data < node.Data {
		return b.Search(node.Left, data)
	}

	if data != node.Data && data > node.Data {
		return b.Search(node.Right, data)
	}

	return false
}

