// Package prefix provides implementations of magnitude prefixes
package prefix

import (
	"fmt"
)

// Prefix is interface for magnitude prefix
type Prefix[T any] interface {
	fmt.Stringer
	Value() float64
	Scale() []T
}
