package search

import (
	"fmt"
)

func binarySerch(arr []int, target int) {
	fmt.Println("Binary Search ====>")
	_binarySerch(arr, 0, len(arr)-1, target)
}

func _binarySerch(arr []int, min, max, target int) {
	if max >= min {
		middle := (min + max) / 2
		if target == arr[middle] {
			fmt.Printf("Element %d is present at %d index\n", target, middle)
		} else if target < arr[middle] {
			_binarySerch(arr, min, middle-1, target)
		} else {
			_binarySerch(arr, middle+1, max, target)
		}
	} else {
		fmt.Printf("Element %d is not present in the array\n", target)
	}
}
