package main

func main() {
	linked_list := &LinkedList{}
	linked_list.add_node(1)
	linked_list.add_node(2)
	linked_list.add_node(3)
	linked_list.add_node(4)
	linked_list.print_list()
	linked_list.remove_node(40)
	linked_list.print_list()
}