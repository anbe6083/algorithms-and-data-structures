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
