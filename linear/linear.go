package linear

import (
	"fmt"
	"os"

	"github.com/akmal4410/datastructure/utils"
)

func LinearDS() {
	var choice int

TryLinear:
	fmt.Println("Choose the Linear algorithm to work :\n1.Singly Linkedlist \n2.Doubly Linkedlist \n3.Stack using Array \n4.Stack using Linkedlist \n5.Queue using Array \n6.Queue using Linkedlist \n7.Go Back \n8.Exit")
	choice = utils.ScanValue(choice)
	switch choice {
	case 1:
		singlyLinkedListFunc()
		goto TryLinear
	case 2:
		doubleLinkedListFunc()
		goto TryLinear
	case 3:
		stackArrayFunc()
		goto TryLinear
	case 4:
		stackLinkedListFunc()
		goto TryLinear
	case 5:
		queueArrayFunc()
		goto TryLinear
	case 6:
		queueLinkedListFunc()
		goto TryLinear
	case 7:
		break
	case -1, 8:
		os.Exit(0)
	default:
		goto TryLinear
	}
}
