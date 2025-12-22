package main

func main() {

}

func SelectionSort(arr []int) []int {
	for j := 0; j < len(arr)-1; j++ {
		smallest := arr[j]
		smallestIndex := j
		for i := j; i < len(arr); i++ {
			if arr[i] < smallest {
				smallest = arr[i]
				smallestIndex = i
			}
		}
		arr[smallestIndex] = arr[j]
		arr[j] = smallest
	}
	return arr
}
