package sort

import (
	"fmt"
	"os"

	"github.com/akmal4410/datastructure/utils"
)

func Sorting() {
	var len, choice, value int
	var arr []int
TrySorting:
	fmt.Println("Choose the sorting algorithm to work :\n1.Bubble Sort \n2.Selection Sort \n3.Insertion Sort \n4.Merge Sort \n5.Go Back \n6.Exit")
	choice = utils.ScanValue(choice)
Length:
	if choice != -1 && choice != 5 && choice != 6 {
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
	}
	switch choice {
	case 1:
		bubbleSort(arr)
		arr = nil
		goto TrySorting
	case 2:
		selectionSort(arr)
		arr = nil
		goto TrySorting
	case 3:
		insertionSort(arr)
		arr = nil
		goto TrySorting
	case 4:
		mergeSort(arr)
		arr = nil
		goto TrySorting
	case 5:
		break
	case -1, 6:
		os.Exit(0)
	default:
		goto TrySorting
	}
}
