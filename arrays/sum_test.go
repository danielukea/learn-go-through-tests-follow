package main

import (
	"reflect"
	"testing"
)

func assertMessage(t *testing.T, got int, want int, numbers []int) {
	if got != want {
		t.Errorf("got %d want %d given: %v", got, want, numbers)
	}
}

func assertArraysEqual(t *testing.T, got []int, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		assertMessage(t, got, want, numbers)

	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assertMessage(t, got, want, numbers)
	})

}

func TestSumAll(t *testing.T) {

	t.Run("two collections of same size", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		assertArraysEqual(t, got, want)

	})
}

func TestSumAllTails(t *testing.T) {

	t.Run("two collections of same size", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assertArraysEqual(t, got, want)

	})

	t.Run("safely handle empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}
		assertArraysEqual(t, got, want)

	})
}
