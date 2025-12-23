package main

import "testing"

func TestBinarySearch(t *testing.T) {

	t.Run("Should return true if the element exists in the array", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		v := 6
		want := true
		got := BinarySearch(arr, v)
		AssertBinarySearch(t, got, want, v, arr)
	})

	t.Run("Should return false if the element does not exist in the array", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		v := 11
		want := false
		got := BinarySearch(arr, v)
		AssertBinarySearch(t, got, want, v, arr)

	})

	t.Run("Should return true if the element exists in slice of size 1", func(t *testing.T) {
		arr := []int{11}
		v := 11
		want := true
		got := BinarySearch(arr, v)
		AssertBinarySearch(t, got, want, v, arr)
	})
	t.Run("Should return false if the element does not exist in slice of size 1", func(t *testing.T) {
		arr := []int{9}
		v := 11
		want := false
		got := BinarySearch(arr, v)
		AssertBinarySearch(t, got, want, v, arr)
	})
	t.Run("Should return false if the element does not exist in slice of size 0", func(t *testing.T) {
		arr := []int{}
		v := 11
		want := false
		got := BinarySearch(arr, v)
		AssertBinarySearch(t, got, want, v, arr)
	})

}

func AssertBinarySearch(t testing.TB, got, want bool, v int, arr []int) {
	t.Helper()
	if got != want {
		t.Errorf("Binary search returned %t, but expected %t. V: %d, Array: %+v", got, want, v, arr)
	}
}
