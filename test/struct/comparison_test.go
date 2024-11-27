package struct_test

import "testing"

/**
 * Struct types are comparable if all their field types are comparable.
 * Two struct values are equal if their corresponding non-blank field values are equal.
 * The fields are compared in source order, and comparison stops as soon as two field values differ (or all fields have been compared).
 *
 * - https://go.dev/ref/spec#Comparison_operators
 */

func TestStructComparison(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	t.Run("Two struct values are equal if their corresponding non-blank field values are equal.", func(t *testing.T) {
		t.Parallel()
		a := person{"John", 30}
		b := person{"John", 30}
		if a != b {
			t.Errorf("%v should equals to %v", a, b)
		}
	})

	t.Run("Two struct values are not equal if their corresponding non-blank field values are not equal.", func(t *testing.T) {
		t.Parallel()
		a := person{"John", 30}
		b := person{"Doe", 30}
		if a == b {
			t.Errorf("%v should not equals to %v", a, b)
		}
	})
}
