package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestInsertionSort(t *testing.T) {

	t.Run("Insertion sort should sort correctly a slice of 10 unsorted items", func(t *testing.T) {

		// Create an unsorted array of 10 items
		arr := []int{43, 12, 8, 5, 10, 99, 1, 23, 7, 15}
		fmt.Println("Unsorted array:", arr)

		got := insertionSort(arr)

		want := []int{1, 5, 7, 8, 10, 12, 15, 23, 43, 99}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Array is not sorted correctly: Got %q, expected %q", got, want)
		}
	})

	t.Run("Should sort a slice of 1 unsorted items", func(t *testing.T) {

		// Create an unsorted array of 10 items
		arr := []int{4}
		fmt.Println("Unsorted array:", arr)

		got := insertionSort(arr)

		want := []int{4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Array is not sorted correctly: Got %q, expected %q", got, want)
		}
	})
	t.Run("Insertion sort should sort correctly a slice of 1000 unsorted items", func(t *testing.T) {
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
		got = insertionSort(got)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Array of %d elements is not sorted correctly", n)
		}
	})

	t.Run("Should do insertion sort in nonincreasing instead of nondecreasing order", func(t *testing.T) {
		// Create an unsorted array of 10 items
		arr := []int{43, 12, 8, 5, 10, 99, 1, 23, 7, 15}
		fmt.Println("Unsorted array:", arr)

		got := insertionSortNonincreasing(arr)

		want := []int{99, 43, 23, 15, 12, 10, 8, 7, 5, 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Array is not sorted correctly: Got %q, expected %q", got, want)
		}
	})

}
