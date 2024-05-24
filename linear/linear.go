package linear

import (
	"fmt"
	"os"

	"github.com/akmal4410/datastructure/utils"
)

func LinearDS() {
	var choice int

TryLinear:
	fmt.Println("Choose the Linear algorithm to work :\n1.Singly Linkedlist \n2.Doubly Linkedlist \n3.Cirucular Linkedlist \n4.Stack using Array \n5.Stack using Linkedlist \n6.Queue using Array \n7.Queue using Linkedlist \n8.Go Back \n9.Exit")
	choice = utils.ScanValue(choice)
	switch choice {
	case 1:
		goto TryLinear
	case 2:
		goto TryLinear
	case 3:
		goto TryLinear
	case 4:
		stackArray()
		goto TryLinear
	case 8:
		break
	case -1, 9:
		os.Exit(0)
	default:
		goto TryLinear
	}
}
