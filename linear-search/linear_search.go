package main

func LinearSearch(array []int, v int) int {
	for i := 0; i < len(array); i++ {
		if array[i] == v {
			return i
		}
	}
	return -1
}
