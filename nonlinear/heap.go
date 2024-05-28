package nonlinear

import (
	"fmt"
	"os"

	"github.com/akmal4410/datastructure/utils"
)

func minHeapFunc() {
	var len, choice, value int
	var arr []int
Length:
	if choice != -1 && choice != 5 && choice != 6 {
		fmt.Println("Enter the length of the array for initial Heap Array:")
		len = utils.ScanValue(len)
		if len == -1 {
			goto Length
		}
		fmt.Println("Enter the array elements : ")
		for i := 0; i < len; i++ {
			value = utils.ScanValue(value)
			arr = append(arr, value)
		}
	}
	minHeap := minHeap{
		heap: arr,
	}
	minHeap.heapify()
TryMinHeap:
	fmt.Println("Choose the Min Heap to work :\n1.Display \n2.Insert \n3.Remove \n4.Go Back \n5.Exit")
	choice = utils.ScanValue(choice)

	switch choice {
	case 1:
		minHeap.display()
		arr = nil
		goto TryMinHeap
	case 2:
		fmt.Println("Enter the value to insert")
		value = utils.ScanValue(value)
		if value == -1 {
			goto TryMinHeap
		}
		minHeap.insert(value)
		arr = nil
		goto TryMinHeap
	case 3:
		minHeap.remove()
		arr = nil
		goto TryMinHeap

	case 4:
		break
	case -1, 5:
		os.Exit(0)
	default:
		goto TryMinHeap
	}
}

type minHeap struct {
	heap []int
}

func (heap *minHeap) heapify() {
	for i := heap.parentIndex(len(heap.heap) - 1); i >= 0; i-- {
		heap.shiftDown(i)
	}
}

func (heap *minHeap) insert(data int) {
	heap.heap = append(heap.heap, data)
	heap.shiftUp(len(heap.heap) - 1)
}

func (heap *minHeap) remove() {
	if heap.heap == nil || len(heap.heap) == 0 {
		return
	}
	heap.heap[0] = heap.heap[len(heap.heap)-1]
	heap.heap = heap.heap[:len(heap.heap)-1]
	heap.shiftDown(0)
}

func (heap *minHeap) display() {
	fmt.Println("Displaying Min Heap")
	if heap.heap == nil || len(heap.heap) == 0 {
		return
	}
	for _, v := range heap.heap {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func (heap *minHeap) shiftDown(currentIndex int) {
	endIndex := len(heap.heap) - 1
	leftChildIndex := heap.leftChildIndex(currentIndex)
	rightChildIndex := heap.rightChildIndex(currentIndex)
	smallIndex := currentIndex
	if leftChildIndex <= endIndex && heap.heap[leftChildIndex] < heap.heap[smallIndex] {
		smallIndex = leftChildIndex
	}
	if rightChildIndex <= endIndex && heap.heap[rightChildIndex] < heap.heap[smallIndex] {
		smallIndex = rightChildIndex
	}
	if currentIndex != smallIndex {
		// heap.heap[smallIndex], heap.heap[currentIndex] = heap.heap[currentIndex], heap.heap[smallIndex]
		temp := heap.heap[smallIndex]
		heap.heap[smallIndex] = heap.heap[currentIndex]
		heap.heap[currentIndex] = temp
		heap.shiftDown(smallIndex)
	}
}

func (heap *minHeap) shiftUp(currentIndex int) {
	parentIndex := heap.parentIndex(currentIndex)
	if currentIndex > 0 && heap.heap[currentIndex] < heap.heap[parentIndex] {
		// Swap elements
		temp := heap.heap[currentIndex]
		heap.heap[currentIndex] = heap.heap[parentIndex]
		heap.heap[parentIndex] = temp
		// Recursive call to continue shifting up
		heap.shiftUp(parentIndex)
	}
}

func (heap *minHeap) parentIndex(currentIndex int) int {
	return (currentIndex - 1) / 2
}

func (heap *minHeap) leftChildIndex(currentIndex int) int {
	return (currentIndex * 2) + 1
}
func (heap *minHeap) rightChildIndex(currentIndex int) int {
	return (currentIndex * 2) + 2
}
