package main

import (
	m "using-go/mergesortedarr"
)

func main() {
	// binary_search_tree := &bst.BinarySearchTree{}
	// binary_search_tree.AddNode(binary_search_tree.Root, 10)
	// binary_search_tree.AddNode(binary_search_tree.Root, 5)
	// binary_search_tree.AddNode(binary_search_tree.Root, 15)
	// binary_search_tree.AddNode(binary_search_tree.Root, 20)
	// binary_search_tree.AddNode(binary_search_tree.Root, 17)
	// binary_search_tree.AddNode(binary_search_tree.Root, 4)
	// binary_search_tree.AddNode(binary_search_tree.Root, 6)
	// binary_search_tree.InOrderTraversal(binary_search_tree.Root)
	// fmt.Println()
	// binaryn(binary_search_tree.Search(binary_search_tree.Root, 110))
	// binary_search_tree.PrintAllLeaves(binary_search_tree.Root)
	// fmt.Println()
	// binary_search_tree.LevelOrderTraversal()
	// fmt.Println()

	// linkedList = &LinkedList
	// myLL := &ll.LinkedList{}
	// myLL.Add(1)
	// myLL.Add(2)
	// myLL.Add(3)
	// myLL.Print()
	// myLL.Remove(3)
	// myLL.Print()

	// m.merge([]int{1,2,3}, 3, []int{2}, 1)
	m.Merge([]int{1,2,3}, 3, []int{2}, 1)
	// 1,0,0,2,3,0,4,5,0
	// [1 0 2 -1]
	// []
	// []
}