package sort

import "fmt"

func hearpSort(array []int) {
	fmt.Println("Heap Sort ====>")
	fmt.Println("Array before soring :", array)
	_heapSort(array)
	fmt.Println("Array after sorting ", array)
}

func _heapSort(array []int) {
	heapify(array)
	n := len(array)
	for len := n - 1; len >= 0; len-- {
		temp := array[0]
		array[0] = array[len]
		array[len] = temp
		shiftDown(array, 0, len)
	}
}

func heapify(array []int) {
	if len(array) == 0 {
		return
	}
	len := len(array)
	for i := parentIndex(len - 1); i >= 0; i-- {
		shiftDown(array, i, len)
	}
}

func parentIndex(currentIndex int) int {
	return (currentIndex - 1) / 2
}

func shiftDown(array []int, currentIndex, len int) {
	endIndex := len - 1
	leftChildIndex := 2*currentIndex + 1
	rightChildIndex := 2*currentIndex + 2
	bigIndex := currentIndex
	if leftChildIndex <= endIndex && array[leftChildIndex] > array[bigIndex] {
		bigIndex = leftChildIndex
	}
	if rightChildIndex <= endIndex && array[rightChildIndex] > array[bigIndex] {
		bigIndex = rightChildIndex
	}
	if currentIndex != bigIndex {
		temp := array[bigIndex]
		array[bigIndex] = array[currentIndex]
		array[currentIndex] = temp
		shiftDown(array, bigIndex, len)
	}
}
