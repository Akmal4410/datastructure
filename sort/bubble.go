package sort

import "fmt"

func bubbleSort(array []int) {
	fmt.Println("Bubble Sort ====>")
	fmt.Println("Array before soring :", array)
	_bubbleSort(array)
	fmt.Println("Array after sorting ", array)
}

// In Bubble sort the last element will get sorted first
func _bubbleSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		//The last element will get sorted in first itration so we can reduce the checking of last number
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
}
