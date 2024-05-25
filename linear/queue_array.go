package linear

import (
	"fmt"
	"os"

	"github.com/akmal4410/datastructure/utils"
)

func queueArrayFunc() {
	var choice, temp int
	stack := queueArray{}
TryLinear:
	fmt.Println("Choose Queue using Array operation :\n1.Display \n2.Enqueue \n3.Dequeue \n4.Go Back \n5.Exit")
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
		stack.enqueue(temp)
		goto TryLinear
	case 3:
		pop := stack.dequeue()
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

type queueArray struct {
	array []int
}

func (s *queueArray) enqueue(data int) {
	s.array = append(s.array, data)
}

func (s *queueArray) dequeue() int {
	if len(s.array) == 0 || s.array == nil {
		return -1
	}
	element := s.array[0]
	s.array = s.array[1:]
	return element
}

func (s *queueArray) display() {
	fmt.Print("Displaing Queue \n===============\n")
	if len(s.array) == 0 || s.array == nil {
		fmt.Println("Empty Queue")
	} else {
		for i := 0; i < len(s.array); i++ {
			fmt.Printf("%d ", s.array[i])
		}
		fmt.Println()
	}
}
