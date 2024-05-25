package linear

import (
	"fmt"
	"os"

	"github.com/akmal4410/datastructure/utils"
)

func queueLinkedListFunc() {
	var choice, temp int
	queue := queueLinkedList{}
TryLinear:
	fmt.Println("Choose Queue using Linkedlist operation :\n1.Display \n2.Enqueue \n3.Dequeue \n4.Go Back \n5.Exit")
	choice = utils.ScanValue(choice)
	switch choice {
	case 1:
		queue.display()
		goto TryLinear
	case 2:
		fmt.Println("Enter the value to insert")
		temp = utils.ScanValue(temp)
		if temp == -1 {
			goto TryLinear
		}
		queue.push(temp)
		goto TryLinear
	case 3:
		pop := queue.pop()
		if pop == -1 {
			fmt.Println("Queue is underflow")
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

type queueLinkedList struct {
	front, rear *node
}

func (q *queueLinkedList) push(data int) {
	newNode := &node{data: data}
	if q.front == nil {
		q.front = newNode
	} else {
		q.rear.next = newNode
	}
	q.rear = newNode
}

func (q *queueLinkedList) pop() int {
	if q.front == nil {
		return -1
	}
	data := q.front.data
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	return data
}

func (s *queueLinkedList) display() {
	fmt.Print("Displaing Queue \n===============\n")
	if s.front == nil {
		fmt.Println("Empty Queue")
		return
	}
	currentNode := s.front
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}
