package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assetEqual := func(t *testing.T, got int, expected int) {
		t.Helper()
		if got != expected {
			t.Errorf("expected %d did not match sum %d", expected, got)
		}
	}

	t.Run("Sum of variable size: ", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		sum := Sum(numbers)
		expected := 10
		assetEqual(t, sum, expected)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{2, 2})
	expected := []int{6, 4}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %d did not match sum %d", expected, got)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("testing non empty slice: ", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4}, []int{3, 4})
		expected := []int{9, 4}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %d did not match sum %d", expected, got)
		}
	})

	t.Run("empty slice: ", func(t *testing.T) {
		got := SumAllTails([]int{})
		expected := []int{0}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %d did not match sum %d", expected, got)
		}
	})
}
