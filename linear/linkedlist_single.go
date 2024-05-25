package linear

import (
	"fmt"
	"os"

	"github.com/akmal4410/datastructure/utils"
)

func singlyLinkedListFunc() {
	var choice, data, temp int
	sList := singleLinkedList{}
TryLinear:
	fmt.Println("Choose Singly Linkedlist operation :\n1.Display \n2.Add New Node \n3.Push \n4.Push Before \n5.Push After \n6.Delete First \n7.Delete Last \n8.Delete Node \n9.Reverse \n10.Go Back \n11.Exit")
	choice = utils.ScanValue(choice)
	switch choice {
	case 1:
		sList.display()
		goto TryLinear
	case 2:
		fmt.Println("Enter the value to insert")
		data = utils.ScanValue(data)
		if data == -1 {
			goto TryLinear
		}
		sList.addNode(data)
		goto TryLinear
	case 3:
		fmt.Println("Enter the value to insert")
		data = utils.ScanValue(data)
		if data == -1 {
			goto TryLinear
		}
		sList.push(data)
		goto TryLinear
	case 4:
		fmt.Println("Enter the value to insert")
		data = utils.ScanValue(data)
		if data == -1 {
			goto TryLinear
		}
		fmt.Println("Enter the value to insert before")
		temp = utils.ScanValue(temp)
		if temp == -1 {
			goto TryLinear
		}
		sList.insertBefore(data, temp)
		goto TryLinear
	case 5:
		fmt.Println("Enter the value to insert")
		data = utils.ScanValue(data)
		if data == -1 {
			goto TryLinear
		}
		fmt.Println("Enter the value to insert after")
		temp = utils.ScanValue(temp)
		if temp == -1 {
			goto TryLinear
		}
		sList.insertAfter(data, temp)
		goto TryLinear
	case 6:
		sList.deleteFirst()
		goto TryLinear
	case 7:
		sList.deleteLast()
		goto TryLinear
	case 8:
		fmt.Println("Enter the value to delete")
		data = utils.ScanValue(data)
		if data == -1 {
			goto TryLinear
		}
		sList.deleteNode(data)
		goto TryLinear
	case 9:
		sList.reverse()
		goto TryLinear
	case 10:
		break
	case -1, 11:
		os.Exit(0)
	default:
		goto TryLinear
	}
}

type singleLinkedList struct {
	head, tail *node
}

func (l *singleLinkedList) display() {
	if l.head == nil {
		fmt.Println("Empty Linked List\n=================")
		return
	}
	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%d -> ", currentNode.data)
		currentNode = currentNode.next
	}
	if currentNode == nil {
		fmt.Println("nil")
	}
}

func (l *singleLinkedList) push(data int) {
	newNode := &node{data: data}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	newNode.next = l.head
	l.head = newNode
}

func (l *singleLinkedList) addNode(data int) {
	newNode := &node{data: data}
	if l.head == nil {
		l.head = newNode
	} else {
		l.tail.next = newNode
	}
	l.tail = newNode
}

func (l *singleLinkedList) insertBefore(data, before int) {
	currentNode := l.head
	var prev *node
	for currentNode != nil && currentNode.data != before {
		prev = currentNode
		currentNode = currentNode.next
	}
	if currentNode == nil {
		fmt.Printf("Cannot insert %d because %d not found\n", data, before)
		return
	}
	if prev == nil {
		l.push(data)
		return
	}
	newNode := &node{data: data}
	newNode.next = currentNode
	prev.next = newNode
}

func (l *singleLinkedList) insertAfter(data, after int) {
	currentNode := l.head
	for currentNode != nil && currentNode.data != after {
		currentNode = currentNode.next
	}
	if currentNode == nil {
		fmt.Printf("Cannot insert %d because %d not found\n", data, after)
		return
	}
	if currentNode.next == nil {
		l.addNode(data)
		return
	}
	newNode := &node{data: data}
	newNode.next = currentNode.next
	currentNode.next = newNode
}

func (l *singleLinkedList) deleteFirst() {
	if l.head == nil {
		fmt.Println("Cannot delete Linkedlist is empty")
		return
	}
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}
}

func (l *singleLinkedList) deleteLast() {
	if l.head == nil {
		fmt.Println("Cannot delete Linkedlist is empty")
		return
	}
	if l.head.next == nil {
		l.deleteFirst()
		return
	}
	currentNode := l.head
	for currentNode.next != l.tail {
		currentNode = currentNode.next
	}
	l.tail = currentNode
	currentNode.next = nil
}

func (l *singleLinkedList) deleteNode(data int) {
	if l.head == nil {
		fmt.Println("Cannot delete Linkedlist is empty")
		return
	}
	if l.head.data == data {
		l.deleteFirst()
		return
	}
	currentNode := l.head
	var prevNode *node
	for currentNode != nil && currentNode.data != data {
		prevNode = currentNode
		currentNode = currentNode.next
	}
	if currentNode == nil {
		fmt.Printf("Cannot delete %d, Not found\n", data)
		return
	}
	if currentNode == l.tail {
		l.deleteLast()
		return
	}
	prevNode.next = currentNode.next
}

func (l *singleLinkedList) reverse() {
	if l.head == nil {
		return
	}
	currentNode := l.head
	var prevNode, nextNode *node
	for currentNode != nil {
		nextNode = currentNode.next
		currentNode.next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}
	temp := l.head
	l.head = l.tail
	l.tail = temp

}
