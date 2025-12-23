package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {

	t.Run("Should sort an array of 1 element", func(t *testing.T) {
		arr := []int{5, 2, 1, 4}
		want := []int{1, 2, 4, 5}
		got := MergeSort(arr)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Array is not sorted: Got %+v, expected %+v", got, want)
		}
	})
}
