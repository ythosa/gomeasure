package units

import (
	"fmt"

	"github.com/ythosa/gomeasure"
	"github.com/ythosa/gomeasure/prefix/metric"
)

// Second represents second unit
type Second struct {
	unit[metric.Prefix]
}

// NewSecond returns new instance of Second
func NewSecond[T gomeasure.Quantitiable](v T, p metric.Prefix) Second {
	return Second{unit: newValue(v, p)}
}

// Sum returns sum of current and other Second unit
func (s Second) Sum(other Second) Second {
	s.unit = s.unit.Sum(other.unit)

	return s
}

// ToPrefix converts from current prefix to passed
func (s Second) ToPrefix(prefix metric.Prefix) Second {
	s.unit = s.unit.ToPrefix(prefix)

	return s
}

// Multiply returns current unit multiplied to k
func (s Second) Multiply(k gomeasure.Quantity) Second {
	s.unit = s.unit.Multiply(k)

	return s
}

// Magnitude returns string representation of magnitude (prefix + s)
func (s Second) Magnitude() string {
	return s.Prefix().String() + "s"
}

// Format returns formatted quantity with passed preciseness and appends magnitude
func (s Second) Format(prec int) string {
	return fmt.Sprintf("%s %s", s.unit.Format(prec), s.Magnitude())
}
