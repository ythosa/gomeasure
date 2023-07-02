package prefix

import (
	"fmt"
)

// Prefix ...
type Prefix[T any] interface {
	fmt.Stringer
	Value() float64
	Scale() []T
}
