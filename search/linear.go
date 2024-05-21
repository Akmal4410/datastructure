package search

import "fmt"

func linearSearch(arr []int, target int) {
	fmt.Println("Linear Search ====>")
	for i, v := range arr {
		if v == target {
			fmt.Printf("Element %d is present at %d index\n", v, i)
			return
		}
	}
	fmt.Printf("Element %d is not present in the array\n", target)
}
