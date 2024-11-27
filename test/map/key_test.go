package map_test

import "testing"

/**
 * KeyType may be any type that is comparable (more on this later)
 *
 * - https://go.dev/blog/maps
 */

func TestMapKey(t *testing.T) {
	t.Run("Array is comparable and can be used as a key in a map", func(t *testing.T) {
		t.Parallel()
		a := [3]int{1, 2, 3}
		b := [3]int{1, 2, 3}

		m := make(map[[3]int]string)
		m[a] = "John"
		if m[b] != "John" {
			t.Errorf("%v should equals to %v", m[b], "John")
		}
	})

	t.Run("Struct is comparable and can be used as a key in a map", func(t *testing.T) {
		t.Parallel()
		type person struct {
			name string
			age  int
		}

		a := person{"John", 30}
		b := person{"John", 30}

		m := make(map[person]string)
		m[a] = "John"
		if m[b] != "John" {
			t.Errorf("%v should equals to %v", m[b], "John")
		}
	})
}
