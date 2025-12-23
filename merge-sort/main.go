package main

func main() {

}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	leftArr := MergeSort(arr[:mid])
	rightArr := MergeSort(arr[mid:])
	return Merge(leftArr, rightArr)
}

func Merge(leftArr, rightArr []int) []int {
	var result []int

	i := 0
	j := 0

	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] <= rightArr[j] {
			result = append(result, leftArr[i])
			i++
		} else {
			result = append(result, rightArr[j])
			j++
		}
	}
	result = append(result, leftArr[i:]...)
	result = append(result, rightArr[j:]...)
	return result
}
