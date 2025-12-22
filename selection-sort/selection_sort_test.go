package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestSelectionSort(t *testing.T) {
	t.Run("Should sort an array of 10 numbers using selection sort", func(t *testing.T) {
		arr := []int{4, 8, 1, 7, 3, 10, 2, 6, 9, 5}
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := SelectionSort(arr)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Selection sort did not sort properly. Got: %+v, expected %+v", got, want)
		}

	})

	t.Run("Should sort an array of 1 element correctly using selection sort", func(t *testing.T) {
		arr := []int{1}
		want := []int{1}
		got := SelectionSort(arr)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Selection sort did not sort properly. Got: %+v, expected %+v", got, want)
		}

	})
	t.Run("Should sort an array of 2 element correctly using selection sort", func(t *testing.T) {
		arr := []int{2, 1}
		want := []int{1, 2}
		got := SelectionSort(arr)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Selection sort did not sort properly. Got: %+v, expected %+v", got, want)
		}

	})
	t.Run("Selection sort should sort correctly a slice of 1000 unsorted items", func(t *testing.T) {
		// Seed random number generator
		rand.Seed(time.Now().UnixNano())

		// Create an unsorted slice of 1000 items
		n := 1000
		got := make([]int, n)
		for i := 0; i < n; i++ {
			got[i] = rand.Intn(10000) // Random integers between 0 and 9999
		}

		// Create a copy of the unsorted slice to determine the "want" value
		want := make([]int, n)
		copy(want, got)
		sort.Ints(want)

		// Run your insertionSort implementation
		got = SelectionSort(got)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Array of %d elements is not sorted correctly", n)
		}
	})
}
