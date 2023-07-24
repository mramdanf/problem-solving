package main

import (
	"fmt"
	bst "using-go/bst"
)

func main() {
	binary_search_tree := &bst.BinarySearchTree{}
	binary_search_tree.AddNode(binary_search_tree.Root, 10)
	binary_search_tree.AddNode(binary_search_tree.Root, 5)
	binary_search_tree.AddNode(binary_search_tree.Root, 15)
	binary_search_tree.AddNode(binary_search_tree.Root, 20)
	binary_search_tree.AddNode(binary_search_tree.Root, 17)
	binary_search_tree.AddNode(binary_search_tree.Root, 4)
	binary_search_tree.AddNode(binary_search_tree.Root, 6)
	binary_search_tree.InOrderTraversal(binary_search_tree.Root)
	fmt.Println()
	binary_search_tree.PreOrderTraversal(binary_search_tree.Root)
	fmt.Println()
	binary_search_tree.PostOrderTraversal(binary_search_tree.Root)
	fmt.Println()
	fmt.Println(binary_search_tree.Search(binary_search_tree.Root, 5))
	fmt.Println(binary_search_tree.Search(binary_search_tree.Root, 17))
	fmt.Println(binary_search_tree.Search(binary_search_tree.Root, 100))
	fmt.Println(binary_search_tree.Search(binary_search_tree.Root, 110))
}