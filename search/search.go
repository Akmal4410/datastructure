package search

import (
	"fmt"
	"os"

	"github.com/akmal4410/datastructure/utils"
)

func Searching() {
	var len, choice, value, target int
	var arr []int
TrySearching:
	fmt.Println("Choose the Searching algorithm to work :\n1.Linear Search \n2.Binary Search \n3.Go Back \n4.Exit")
	choice = utils.ScanValue(choice)
Length:
	if choice != -1 && choice != 3 && choice != 4 {

		fmt.Println("Enter the length of the array :")
		len = utils.ScanValue(len)
		if len == -1 {
			goto Length
		}
		fmt.Println("Enter the array elements : ")
		for i := 0; i < len; i++ {
			value = utils.ScanValue(value)
			arr = append(arr, value)
		}
		fmt.Println("Enter the Search Value :")
		target = utils.ScanValue(target)
		if target == -1 {
			goto TrySearching
		}
	}
	switch choice {
	case 1:
		linearSearch(arr, target)
		arr = nil
		goto TrySearching
	case 2:
		binarySerch(arr, target)
		arr = nil
		goto TrySearching
	case 3:
		break
	case -1, 4:
		os.Exit(0)
	default:
		goto TrySearching
	}
}
