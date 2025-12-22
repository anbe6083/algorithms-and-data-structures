package main

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	t.Run("Should should for a value v and return the index", func(t *testing.T) {
		arr := []int{43, 12, 8, 5, 10, 99, 1, 23, 7, 15}
		fmt.Println("Unsorted array:", arr)
		v := 99
		got := LinearSearch(arr, v)
		want := 5
		if got != want {
			t.Errorf("Wrong index was returned, got %d, want %d", got, want)
		}
	})
	t.Run("Should return -1 if a value v doesnt exist in an array during linear search", func(t *testing.T) {
		arr := []int{43, 12, 8, 5, 10, 99, 1, 23, 7, 15}
		v := 88
		got := LinearSearch(arr, v)
		want := -1
		if got != want {
			t.Errorf("Wrong index was returned, got %d, want %d", got, want)
		}
	})
}
