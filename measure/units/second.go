package units

import (
	"fmt"

	"gitlab.ozon.ru/at/measure/measure"
	"gitlab.ozon.ru/at/measure/measure/prefix/metric"
)

// Second ...
type Second struct {
	value[metric.Prefix]
}

// NewSecond ...
func NewSecond[T measure.Quantitiable](v T, p metric.Prefix) Second {
	return Second{value: newValue(v, p)}
}

// Sum ...
func (s Second) Sum(other Second) Second {
	s.value = s.value.Sum(other.value)

	return s
}

// ToPrefix ...
func (s Second) ToPrefix(prefix metric.Prefix) Second {
	s.value = s.value.ToPrefix(prefix)

	return s
}

// Multiply ...
func (s Second) Multiply(k measure.Quantity) Second {
	s.value = s.value.Multiply(k)

	return s
}

// Magnitude ...
func (s Second) Magnitude() string {
	return s.Prefix().String() + "s"
}

// Format ...
func (s Second) Format(prec int) string {
	return fmt.Sprintf("%s %s", s.value.Format(prec), s.Magnitude())
}
