package liskov_substitution

import "fmt"

/**
 * Liskov Substitution Principle:
 *
 * Subtypes must be substitutable for their base types
 */

/**
 * In this example, the Penguin struct violates the Liskov Substitution Principle.
 *
 * The Bird struct has a Fly method, but the Penguin struct overrides the Fly method and panics.
 */
type Bird struct {
	Name string
}

func (b Bird) Fly() {
	fmt.Println(b.Name + " is flying")
}

type Penguin struct {
	Bird
}

func (p Penguin) Fly() {
	panic("Penguins can't fly")
}
