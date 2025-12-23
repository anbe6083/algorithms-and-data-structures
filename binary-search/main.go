package main

func main() {

}

func BinarySearch(arr []int, v int) bool {
	if len(arr) < 1 {
		return false
	}
	mid := len(arr) / 2
	if arr[mid] == v {
		return true
	}
	if v < arr[mid] {
		return BinarySearch(arr[:mid], v)
	} else {
		return BinarySearch(arr[mid+1:], v)
	}
}

func BinarySearchIterative(arr []int, v int) bool {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == v {
			return true
		} else if v < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
