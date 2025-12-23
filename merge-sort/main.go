package main

func main() {

}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	q := len(arr) / 2
	left := MergeSort(arr[:q])
	right := MergeSort(arr[q:])
	sorted := Merge(left, right)

	return sorted
}

func Merge(leftArr, rightArr []int) []int {
	arr := []int{}
	i, j := 0, 0
	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] <= rightArr[j] {
			arr = append(arr, leftArr[i])
			i++
		} else {
			arr = append(arr, rightArr[j])
			j++
		}
	}
	arr = append(arr, leftArr[i:]...)
	arr = append(arr, rightArr[j:]...)
	return arr
}
