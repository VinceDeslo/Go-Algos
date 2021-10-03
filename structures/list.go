package main

import (
	"fmt"
)

/*
Insertion at end: O(n)
Insertion at index: O(i)
Removal at index: O(i)
Access at index: O(i)
List contents: O(n)
Concatenation: O(n1)
*/

type Node struct {
	value int
	next  *Node
}

type List struct {
	name string
	head *Node
	size int
}

// Constructor for a new list
func createList(name string) *List {
	return &List{name: name}
}

// Insertion method at end of list
func (l *List) insertAtEnd(value int) {
	new := &Node{value: value}
	last := l.head
	l.size += 1

	if last == nil {
		l.head = new
		return
	}
	for last.next != nil {
		last = last.next
	}
	last.next = new
}

// Method for insertion with node indexes
func (l *List) insertAtIndex(value int, index int) {
	curr, prev := l.head, l.head
	new := &Node{value: value}

	if index >= l.size {
		fmt.Println("Index out of bounds")
		return
	}
	for i := 0; i < index; i++ {
		prev = curr
		curr = curr.next
	}
	prev.next = new
	new.next = curr
}

// Accessor utilizing node indexes
func (l *List) getByIndex(index int) *Node {
	curr := l.head
	if index >= l.size {
		fmt.Println("Index out of bounds")
		return nil
	}
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	return curr
}

// Destructor utilizing node indexes
func (l *List) removeAtIndex(index int) {
	curr, prev := l.head, l.head

	if index >= l.size {
		fmt.Println("Index out of bounds")
		return
	}
	if index == 0 {
		l.head = l.head.next
		return
	}
	for i := 0; i < index; i++ {
		prev = curr
		curr = curr.next
	}
	if curr == nil {
		return
	}
	prev.next = curr.next
}

// Utility to concat two lists
func (l1 *List) concatList(l2 *List) {
	last := l1.head
	for last.next != nil {
		last = last.next
	}
	last.next = l2.head
}

// Recursive utility to view list contents
func printList(curr *Node) {
	if curr != nil {
		fmt.Println(*curr)
		printList(curr.next)
	}
}

// Main entry point
func main() {
	list := createList("Value List")
	fmt.Println("Initial list")
	list.insertAtEnd(7)
	list.insertAtEnd(12)
	list.insertAtEnd(3)
	list.insertAtEnd(14)
	printList(list.head)

	list.insertAtIndex(15, 2)
	fmt.Println("List post insertion")
	printList(list.head)

	list.removeAtIndex(2)
	fmt.Println("List post removal")
	printList(list.head)

	fmt.Println("Accessor calls")
	fmt.Println(list.getByIndex(0))
	fmt.Println(list.getByIndex(1))
	fmt.Println(list.getByIndex(2))
	fmt.Println(list.getByIndex(3))

	fmt.Println("Concatenation test")
	first := createList("First list")
	first.insertAtEnd(2)
	first.insertAtEnd(4)
	second := createList("Second list")
	second.insertAtEnd(1)
	second.insertAtEnd(3)
	first.concatList(second)
	printList(first.head)
}
