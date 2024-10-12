// Base Linked List Implementation
package LinkedList

import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func (list *List) show() {
	current := list.head
	if current == nil {
		fmt.Println("List is empty")
		return
	}
	fmt.Println("List is:")
	for current != nil {
		fmt.Printf("%v \n", current.data)
		current = current.next
	}
}
