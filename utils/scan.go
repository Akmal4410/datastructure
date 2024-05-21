package utils

import "fmt"

func ScanValue(value int) int {
	_, err := fmt.Scan(&value)
	if err != nil {
		fmt.Println("Error Scan:", err)
		return -1
	}
	return value
}
