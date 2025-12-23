package tests

import (
	"math/rand/v2"
	"reflect"
	"slices"
	"testing"
)

func CreateRandomArray(len int, intRange int) ([]int, []int) {
	var (
		unsorted []int
		sorted   []int
	)
	for i := 0; i < len; i++ {
		num := rand.IntN(intRange) + 1
		unsorted = append(unsorted, num)
		sorted = append(sorted, num)
	}
	slices.Sort(sorted)
	return unsorted, sorted
}

func AssertTwoArraysSorted(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Array is not sorted: Got %+v, expected %+v", got, want)
	}
}
