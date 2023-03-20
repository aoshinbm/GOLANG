package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Add(value int) {
	newNode := &Node{Value: value}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	fmt.Println(current)
	for current.Next != nil {
		current = current.Next
		fmt.Println(current)
	}
	current.Next = newNode
	fmt.Println(current.Next)
}

func (l *LinkedList) Append(value int) {
	newNode := &Node{Value: value, Next: l.Head}
	l.Head = newNode
}

func (l *LinkedList) Display() {
	if l.Head == nil {
		fmt.Println("Linked List is empty.")
		return
	}
	current := l.Head
	fmt.Println("Linked List values:")
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

func main() {
	list := &LinkedList{}
	list.Add(1)
	list.Add(2)
	//	list.Add(3)
	list.Append(0)
	list.Display()
}
