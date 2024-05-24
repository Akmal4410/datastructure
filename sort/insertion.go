package sort

import "fmt"

func insertionSort(array []int) {
	fmt.Println("Insertion Sort ====>")
	fmt.Println("Array before soring :", array)
	_insertionSort(array)
	fmt.Println("Array after sorting ", array)
}

func _insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		key := array[i]
		j := i - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j = j - 1
		}
		array[j+1] = key
	}
}
