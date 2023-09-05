package main

import (
	rd "using-go/removeduplicate"
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
	// m.Merge([]int{1,2,3}, 3, []int{2}, 1)
	// 1,0,0,2,3,0,4,5,0
	// [1 0 2 -1]
	// []
	// []
	// r.RemoveElement([]int{0,1,2}, 2) // 2 -> 0 1
	// r.RemoveElement([]int{2,2,2,2,1}, 2) // 1 -> 1
	// r.RemoveElement([]int{0,2,2,2,1}, 2) // 2 -> 0 1
	// r.RemoveElement([]int{0,1,2,2,1}, 2) // 3 -> 0 1 1
	// r.RemoveElement([]int{0,1,4,2,1}, 2) // 4 -> 0 1 4 1
	// r.RemoveElement([]int{3,2,2,3}, 3) // 2 -> 2 2
	// duplicateZeros([]int{1,0,2,3,0,4,5,0}) // 1,0,0,2,3,0,0,4,5,0,0
	// duplicateZeros([]int{4,5,0}) // 4,5,0,0
	rd.RemoveDuplicate([]int{1,2,2,3}) // 1,2,3
	rd.RemoveDuplicate([]int{1,1,2,3}) // 1,2,3
	rd.RemoveDuplicate([]int{1,1,2,3,3}) // 1,2,3
	// fmt.Println(de.CheckDoubleExist([]int{10,2,5,3}))
	// fmt.Println(de.CheckDoubleExist([]int{3,1,7,11}))
	// fmt.Println(de.CheckDoubleExist([]int{-2,0,10,-19,4,6,-8}))
	// fmt.Println(vma.ValidMountainArray([]int{2,1})) // false
	// fmt.Println(vma.ValidMountainArray([]int{3,5,5})) // false
	// fmt.Println(vma.ValidMountainArray([]int{3,5,6})) // false
	// fmt.Println(vma.ValidMountainArray([]int{0,3,2,1})) // true
	// fmt.Println(vma.ValidMountainArray([]int{2,0,2})) // false
	// fmt.Println(vma.ValidMountainArray([]int{0,1,2,1,2})) // false
}

