package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedListTable struct {
	head *Node
	tail *Node
}

func main() {
	fmt.Println("--------------------------------------------------------------")
	//add i.e. add last element in the linked list
	list := LinkedListTable{}
	list.AddElementsLast(30)
	list.AddElementsLast(33)
	list.AddElementsLast(98)
	list.AddElementsLast(67)
	list.AddElementsLast(1)
	list.Display()

	fmt.Println("--------------------------------------------------------------")
	//append i.e. add first in the linked list
	list.AppendElement(78)
	list.AppendElement(55)
	list.AppendElement(9)
	list.Display()

	fmt.Println("--------------------------------------------------------------")
	//insertion
	list.InsertElement(15)
	list.Display()

	fmt.Println("--------------------------------------------------------------")
	/*	//search element from the linkedlist
		var find int
		fmt.Println("Enter the element you want to search: ")
		fmt.Scanf("%d", &find)
		list.SearchElement(find)

		fmt.Println("--------------------------------------------------------------")
	*/
	//search element from the linkedlist
	var delete int
	fmt.Println("Enter the element you want to delete: ")
	fmt.Scanf("%d", &delete)
	list.DeleteElement(delete)
	list.Display()

}

//creating a linked list
func createLinkedList(value int) {
	head := &Node{data: 0}
	current := head
	for i := 1; i < 5; i++ {
		current.next = &Node{data: i}
		current = current.next
	}
	for node := head; node != nil; node = node.next {
		fmt.Println(node.data)

	}
}
func (l *LinkedListTable) AddElementsLast(value int) {
	newNode := &Node{data: value}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	//fmt.Println(current)

	for current.next != nil {
		current = current.next
		//fmt.Println(current)
	}
	current.next = newNode
	//fmt.Println(current.next)
	//fmt.Println(newNode)
}

func (l *LinkedListTable) SearchElement(search int) {
	if l.head == nil {
		fmt.Println("LinkedList is empty")
	}
	for l.head != nil {
		current := l.head
		//fmt.Println(current)
		for current.next != nil {
			current = current.next
			//fmt.Println(current)
			if current.next.data == search {
				fmt.Println("Element ", current.next.data, "is found ")
				break
			}
			fmt.Println("Element not found ")
		}
	}
}

func (l *LinkedListTable) DeleteElement(delete int) {
	if l.head == nil {
		fmt.Println("LinkedList is empty")
		return
	}
	if l.head.data == delete {
		//deleting the head node
		l.head = l.head.next
		return
	}
	current := l.head
	for current.next != nil {
		if current.next.data == delete {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

func (l *LinkedListTable) AppendElement(value int) {
	newNode := &Node{data: value, next: l.head}
	l.head = newNode
}

func (l *LinkedListTable) Display() {
	if l.head == nil {
		fmt.Println("Linked List is empty.")
		return
	}
	current := l.head
	fmt.Println("Linked List values:")
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

// Insert inserts new node at the end of  from linked list
func (l *LinkedListTable) InsertElement(location *Node, value int) {
	newNode := &Node{data: value}
	current := l.head
	for current != nil {
		if current.next == nil {
			current.next = newNode
			return
		}
		else{
			
				
			}
		}
	}
		//		current = current.next
		//		current.next = newNode
}
