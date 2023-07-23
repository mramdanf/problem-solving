package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) add_node(data int) {
	node := &Node{data: data, next: nil}
	if ll.head == nil {
		ll.head = node
		return
	}

	cur := ll.head
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = node
}

func (ll *LinkedList) print_list() {
	curr := ll.head
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
}