package array_test

import "testing"

/**
 * Array types are comparable if their array element types are comparable.
 * Two array values are equal if their corresponding element values are equal.
 * The elements are compared in ascending index order, and comparison stops as soon as two element values differ (or all elements have been compared).
 *
 * - https://go.dev/ref/spec#Comparison_operators
 */

func TestArrayComparison(t *testing.T) {
	t.Run("Two array values are equal if their corresponding element values are equal", func(t *testing.T) {
		t.Parallel()
		a := [3]int{1, 2, 3}
		b := [3]int{1, 2, 3}
		if a != b {
			t.Errorf("%v should equals to %v", a, b)
		}
	})

	t.Run("Two array values are not equal if their corresponding element values are not equal", func(t *testing.T) {
		t.Parallel()
		a := [3]int{1, 2, 3}
		b := [3]int{1, 2, 4}
		if a == b {
			t.Errorf("%v should not equals to %v", a, b)
		}
	})
}
