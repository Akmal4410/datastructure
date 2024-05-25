package linear

import (
	"fmt"
	"os"

	"github.com/akmal4410/datastructure/utils"
)

func queueArrayFunc() {
	var choice, temp int
	queue := queueArray{}
TryLinear:
	fmt.Println("Choose Queue using Array operation :\n1.Display \n2.Enqueue \n3.Dequeue \n4.Go Back \n5.Exit")
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
		queue.enqueue(temp)
		goto TryLinear
	case 3:
		pop := queue.dequeue()
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

type queueArray struct {
	array []int
}

func (q *queueArray) enqueue(data int) {
	q.array = append(q.array, data)
}

func (q *queueArray) dequeue() int {
	if len(q.array) == 0 || q.array == nil {
		return -1
	}
	element := q.array[0]
	q.array = q.array[1:]
	return element
}

func (q *queueArray) display() {
	fmt.Print("Displaing Queue \n===============\n")
	if len(q.array) == 0 || q.array == nil {
		fmt.Println("Empty Queue")
	} else {
		for i := 0; i < len(q.array); i++ {
			fmt.Printf("%d ", q.array[i])
		}
		fmt.Println()
	}
}
