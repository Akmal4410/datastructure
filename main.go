package main

import (
	"fmt"
	"strconv"

	"github.com/akmal4410/datastructure/sorting"
)

func main() {
	fmt.Println("Hai Welcome to Data Structure example\nWhat would you try to do: ?")
TryAgain:
	fmt.Println("1. Sorting")
	var strValue string
	strValue = scanValue(strValue)
	val, err := strconv.Atoi(strValue)
	if err != nil {
		fmt.Println("Error", err)
		fmt.Println("Do you want to try again ?\nEnter y - to continue")
		var tryAgain string
		_, err := fmt.Scan(&tryAgain)
		if err != nil {
			fmt.Println("Error Scan:", err)
			return
		}
		if tryAgain == "y" {
			goto TryAgain
		} else {
			return
		}
	}
	switch val {
	case 1:
		sorting.BubbleSort()
	case 2:
		sorting.SelectionSort()
	default:
		fmt.Printf("You Entered wrong choice: %d", val)
	}

}

func scanValue(strValue string) string {
	_, err := fmt.Scan(&strValue)
	if err != nil {
		fmt.Println("Error Scan:", err)
		return ""
	}
	return strValue
}
