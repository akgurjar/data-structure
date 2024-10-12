// package main

// import "fmt"

// type Node struct {
// 	data int
// 	next *Node
// }

// type List struct {
// 	head *Node
// }

// func (list *List) add(data int) {
// 	node := &Node{data: data}
// 	if list.head == nil {
// 		list.head = node
// 		return
// 	}
// 	last := list.head
// 	for last.next != nil {
// 		last = last.next
// 	}
// 	last.next = node
// }

// func (list *List) show() {
// 	current := list.head
// 	if current == nil {
// 		fmt.Println("List is empty")
// 		return
// 	}
// 	fmt.Println("List is:")
// 	for current != nil {
// 		fmt.Printf("%v \n", current.data)
// 		current = current.next
// 	}
// }

// func (list *List) reverse() {
// 	if list.head == nil {
// 		return
// 	}
// 	var prev *Node
// 	for list.head.next != nil {
// 		next := list.head.next
// 		list.head.next = prev
// 		prev = list.head
// 		list.head = next
// 	}
// 	list.head.next = prev
// }

// func (list *List) middle() {
// 	current := list.head
// 	if current == nil {
// 		fmt.Println("List is empty")
// 		return
// 	}
// 	counter := 0
// 	middle := list.head
// 	for current.next != nil {
// 		if counter == 1 {
// 			counter = 0
// 			middle = middle.next
// 		} else {
// 			counter += 1
// 		}
// 		current = current.next
// 	}
// 	if counter == 1 {
// 		fmt.Printf("Middle Node: %v & %v \n", middle.data, middle.next.data)
// 	} else {
// 		fmt.Printf("Middle Node: %v \n", middle.data)
// 	}
// }

// func main() {
// 	list := &List{}

// 	list.add(12)
// 	list.add(8)
// 	list.add(16)
// 	list.add(11)
// 	list.show()
// 	list.reverse()
// 	list.show()

// }
