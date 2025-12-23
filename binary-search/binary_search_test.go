package main

import "testing"

func TestBinarySearch(t *testing.T) {

	t.Run("Should return true if the element exists in the array", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		v := 6
		want := true
		got := BinarySearch(arr, v)

		if got != want {
			t.Errorf("Binary search returned false, but expected true. V: %d, Array: %+v", v, arr)
		}
	})

	t.Run("Should return false if the element does not exist in the array", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		v := 11
		want := false
		got := BinarySearch(arr, v)

		if got != want {
			t.Errorf("Binary search returned false, but expected true. V: %d, Array: %+v", v, arr)
		}

	})
}
