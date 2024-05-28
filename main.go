package main

import (
	"fmt"
	"os"

	"github.com/akmal4410/datastructure/linear"
	"github.com/akmal4410/datastructure/nonlinear"
	"github.com/akmal4410/datastructure/search"
	"github.com/akmal4410/datastructure/sort"
	"github.com/akmal4410/datastructure/utils"
)

func main() {
	fmt.Println("Hai Welcome to Data Structure example\nWhat would you try to do: ?")
TryAgain:
	value := chooseDataStructure()
	switch value {
	case 1:
		linear.LinearDS()
		goto TryAgain
	case 2:
		nonlinear.NonLinearDS()
		goto TryAgain
	case 3:
		sort.Sorting()
		goto TryAgain
	case 4:
		search.Searching()
		goto TryAgain
	case -1, 5:
		os.Exit(0)
	}
}

func chooseDataStructure() int {
	var ans int
	fmt.Println("\nChoose the data structure to work:")
	fmt.Println("1.Linear  \n2.Non Linear \n3.Sorting \n4.Searching \n5.Exit")
	return utils.ScanValue(ans)
}
