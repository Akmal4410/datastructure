package nonlinear

import (
	"fmt"
	"os"

	"github.com/akmal4410/datastructure/utils"
)

func NonLinearDS() {
	var choice int

TryNonLinear:
	fmt.Println("Choose the Non Linear algorithm to work :\n1.Binary Search Tree \n2.Min Heap \n3.Go Back \n4.Exit")
	choice = utils.ScanValue(choice)
	switch choice {
	case 1:
		binarySerchTreeFunc()
		goto TryNonLinear
	case 2:
		minHeapFunc()
		goto TryNonLinear
	case 3:
		break
	case -1, 4:
		os.Exit(0)
	default:
		goto TryNonLinear
	}
}
