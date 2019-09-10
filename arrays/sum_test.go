package main

import (
	"reflect"
	"testing"
)

func _checkSums(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("no tail - make the sums of some slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		_checkSums(t, got, want)
	})

	t.Run("no tail - safely sum empty slices", func(t *testing.T) {
		got := SumAll([]int{}, []int{3, 4, 5})
		want := []int{0, 12}
		_checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("with tail - make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		_checkSums(t, got, want)
	})

	t.Run("with tail - safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		_checkSums(t, got, want)
	})
}
