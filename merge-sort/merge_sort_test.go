package main

import (
	"reflect"
	"testing"

	tests "github.com/anbe6083/intro-to-algorithms"
)

func TestMergeSort(t *testing.T) {

	t.Run("Should sort an array of 4 element", func(t *testing.T) {
		arr := []int{5, 2, 1, 4}
		want := []int{1, 2, 4, 5}
		got := MergeSort(arr)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Array is not sorted: Got %+v, expected %+v", got, want)
		}
	})
	t.Run("Should sort an array of 100 elements", func(t *testing.T) {
		arr, want := tests.CreateRandomArray(100, 100)
		got := MergeSort(arr)
		tests.AssertTwoArraysSorted(t, got, want)
	})

	t.Run("Should sort an array of 10000 elements", func(t *testing.T) {
		arr, want := tests.CreateRandomArray(10000, 10000)
		got := MergeSort(arr)
		tests.AssertTwoArraysSorted(t, got, want)
	})

}
