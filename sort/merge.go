package sort

import "fmt"

func mergeSort(array []int) {
	fmt.Println("Merge Sort ====>")
	fmt.Println("Array before soring :", array)
	mergeArray := _mergeSort(array)
	fmt.Println("Array after sorting ", mergeArray)
}

func _mergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	middle := len(array) / 2
	leftArray := _mergeSort(array[0:middle])
	righArray := _mergeSort(array[middle:])
	return _mergeSortHelper(leftArray, righArray)
}

func _mergeSortHelper(leftArray, rightArray []int) []int {
	var mergeArray []int
	left := 0
	right := 0
	for left < len(leftArray) && right < len(rightArray) {
		if leftArray[left] < rightArray[right] {
			mergeArray = append(mergeArray, leftArray[left])
			left++
		} else {
			mergeArray = append(mergeArray, rightArray[right])
			right++
		}
	}
	for left < len(leftArray) {
		mergeArray = append(mergeArray, leftArray[left])
		left++
	}
	for right < len(rightArray) {
		mergeArray = append(mergeArray, rightArray[right])
		right++
	}
	return mergeArray
}
