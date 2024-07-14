package units

import (
	"fmt"

	"github.com/ythosa/gomeasure"
	"github.com/ythosa/gomeasure/prefix/metric"
)

// CPU represents cpu unit
type CPU struct {
	unit[metric.Prefix]
}

// NewCPU returns new instance of CPU
func NewCPU[T gomeasure.Quantitiable](v T, p metric.Prefix) CPU {
	return CPU{unit: newValue(v, p)}
}

// Sum returns sum of current and other CPU unit
func (c CPU) Sum(other CPU) CPU {
	c.unit = c.unit.Sum(other.unit)

	return c
}

// ToPrefix converts from current prefix to passed
func (c CPU) ToPrefix(prefix metric.Prefix) CPU {
	c.unit = c.unit.ToPrefix(prefix)

	return c
}

// Multiply returns current unit multiplied to k
func (c CPU) Multiply(k gomeasure.Quantity) CPU {
	c.unit = c.unit.Multiply(k)

	return c
}

// Magnitude returns string representation of magnitude (prefix + "cpu")
func (c CPU) Magnitude() string {
	return c.Prefix().String() + "CPU"
}

// Format returns formatted quantity with passed preciseness and appends magnitude
func (c CPU) Format(prec int) string {
	c.unit = c.unit.prettier()

	return fmt.Sprintf("%s %s", c.unit.Format(prec), c.Magnitude())
}
