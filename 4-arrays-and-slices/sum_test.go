package arraySlices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	checkListSumsTo := func(t testing.TB, numbers []int, sum int) {
		t.Helper()

		got := Sum(numbers)
		want := sum

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	}

	t.Run("Sum 4 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		sum := 10

		checkListSumsTo(t, numbers, sum)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{5, 5, 5})
	want := []int{6, 15}

	checkSums(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("Test sum tails with 2 slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{5, 5, 5})
		want := []int{5, 10}

		checkSums(t, got, want)
	})
	t.Run("Safely sum empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{9, 8, 10, 12})
		want := []int{0, 30}

		checkSums(t, got, want)
	})
}

func checkSums(t testing.TB, got, want []int) {
	t.Helper()

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
