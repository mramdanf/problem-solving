package linked_list

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Add(data int) {
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

func (ll *LinkedList) Print() {
	curr := ll.head
	for curr != nil {
		fmt.Print(curr.data, " ")
		curr = curr.next
	}
	fmt.Println()
}

func (ll *LinkedList) Remove(data int) {
	if ll.head == nil {
		return
	}

	if ll.head.data == data {
		ll.head = ll.head.next
		return
	}

	cur := ll.head
	for cur.next != nil && cur.next.data != data {
		cur = cur.next
	}

	if cur.next != nil {
		cur.next = cur.next.next
		return
	}
	
	if cur.next == nil {
		fmt.Println("Node not found.")
	}
}