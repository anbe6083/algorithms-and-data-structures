package main

import "fmt"

func main() {

	// Create an unsorted array of 10 items
	arr := []int{43, 12, 8, 5, 10, 99, 1, 23, 7, 15}
	fmt.Println("Unsorted array:", arr)

	sorted := insertionSort(arr)
	fmt.Println("Sorted array: ", sorted)
}

func insertionSort(array []int) []int {
	for j := 1; j < len(array); j++ {
		key := array[j]
		i := j - 1

		for i >= 0 && array[i] > key {
			array[i+1] = array[i]
			i--
			array[i+1] = key
		}
	}
	return array
}

func insertionSortNonincreasing(array []int) []int {
	for j := 1; j < len(array); j++ {
		key := array[j]
		i := j - 1
		for i >= 0 && array[i] < key {
			array[i+1] = array[i]
			i--
			array[i+1] = key
		}
	}
	return array
}
