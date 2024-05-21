package sort

import "fmt"

func selectionSort(array []int) {
	fmt.Println("Array before soring :", array)
	sort(array)
	fmt.Println("Array after sorting ", array)
}

func sort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		smallIndex := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[smallIndex] {
				smallIndex = j
			}
		}
		if smallIndex != i {
			temp := array[i]
			array[i] = array[smallIndex]
			array[smallIndex] = temp
		}
	}
}
