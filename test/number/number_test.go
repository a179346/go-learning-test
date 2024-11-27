package number_test

import (
	"reflect"
	"testing"
)

func TestFloat(t *testing.T) {
	t.Run("Quotient is int when dividend and divisor are both int.", func(t *testing.T) {
		dividend := 10
		divisor := 3
		quotient := dividend / divisor

		kind := reflect.TypeOf(quotient).Kind()
		if kind != reflect.Int {
			t.Errorf("Expected quotient to be int, but got %s", kind)
		}

		if quotient != 3 {
			t.Errorf("Expected quotient to be 3, but got %d", quotient)
		}
	})

	t.Run("Quotient is float when dividend and divisor are both float", func(t *testing.T) {
		dividend := 10.0
		divisor := 3.0
		quotient := dividend / divisor

		kind := reflect.TypeOf(quotient).Kind()
		if kind != reflect.Float64 {
			t.Errorf("Expected quotient to be float64, but got %s", kind)
		}

		if quotient != 3.3333333333333335 {
			t.Errorf("Expected quotient to be 3.3333333333333335, but got %f", quotient)
		}
	})
}
