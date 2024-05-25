package linear

import (
	"fmt"
	"os"

	"github.com/akmal4410/datastructure/utils"
)

func stackArrayFunc() {
	var choice, temp int
	stack := stackArray{}
TryLinear:
	fmt.Println("Choose Stack using Array operation :\n1.Display \n2.Push \n3.Pop \n4.Go Back \n5.Exit")
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

type stackArray struct {
	array []int
}

func (s *stackArray) push(data int) {
	s.array = append(s.array, data)
}

func (s *stackArray) pop() int {
	if len(s.array) == 0 || s.array == nil {
		return -1
	}
	element := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return element
}

func (s *stackArray) display() {
	fmt.Print("Displaing stack \n===============\n")
	if len(s.array) == 0 || s.array == nil {
		fmt.Println("Empty Stack")
	} else {
		for i := len(s.array) - 1; i >= 0; i-- {
			fmt.Println(s.array[i])
		}
	}
}
