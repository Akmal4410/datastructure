package linear

import (
	"fmt"
	"os"

	"github.com/akmal4410/datastructure/utils"
)

func stackLinkedListFunc() {
	var choice, temp int
	stack := stackLinkedList{}
TryLinear:
	fmt.Println("Choose Stack using Linkedlist operation :\n1.Display \n2.Push \n3.Pop \n4.Go Back \n5.Exit")
	choice = utils.ScanValue(choice)
	switch choice {
	case 1:
		stack.display()
		goto TryLinear
	case 2:
		fmt.Println("Enter the value to insert")
		temp = utils.ScanValue(temp)
		if temp == -1 {
			goto TryLinear
		}
		stack.push(temp)
		goto TryLinear
	case 3:
		pop := stack.pop()
		if pop == -1 {
			fmt.Println("Stack is underflow")
		} else {
			fmt.Printf("Poped value is %d\n", pop)
		}
		goto TryLinear
	case 4:
		break
	case -1, 5:
		os.Exit(0)
	default:
		goto TryLinear
	}
}

type stackLinkedList struct {
	top *node
}

func (s *stackLinkedList) push(data int) {
	newNode := &node{data: data}
	if s.top != nil {
		newNode.next = s.top
	}
	s.top = newNode
}

func (s *stackLinkedList) pop() int {
	if s.top == nil {
		return -1
	}
	data := s.top.data
	s.top = s.top.next
	return data
}

func (s *stackLinkedList) display() {
	fmt.Print("Displaing stack \n===============\n")
	if s.top == nil {
		fmt.Println("Empty Stack")
		return
	}
	currentNode := s.top
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}
