package units

import (
	"fmt"

	"github.com/ythosa/gomeasure"
	"github.com/ythosa/gomeasure/prefix/metric"
)

// Amount represents amount unit
type Amount struct {
	unit[metric.Prefix]
}

// NewAmount returns new instance of Amount
func NewAmount[T gomeasure.Quantitiable](v T, p metric.Prefix) Amount {
	return Amount{unit: newValue(v, p)}
}

// Sum returns sum of current and other Amount unit
func (a Amount) Sum(other Amount) Amount {
	a.unit = a.unit.Sum(other.unit)

	return a
}

// ToPrefix converts from current prefix to passed
func (a Amount) ToPrefix(prefix metric.Prefix) Amount {
	a.unit = a.unit.ToPrefix(prefix)

	return a
}

// Multiply returns current unit multiplied to k
func (a Amount) Multiply(k gomeasure.Quantity) Amount {
	a.unit = a.unit.Multiply(k)

	return a
}

// Magnitude returns string representation of magnitude (prefix + "")
func (a Amount) Magnitude() string {
	return a.Prefix().String() + "шт."
}

// Format returns formatted quantity with passed preciseness and appends magnitude
func (a Amount) Format(prec int) string {
	a.unit = a.unit.prettier()

	return fmt.Sprintf("%s %s", a.unit.Format(prec), a.Magnitude())
}
