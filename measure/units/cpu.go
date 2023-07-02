package units

import (
	"fmt"

	"gitlab.ozon.ru/at/measure/measure"
	"gitlab.ozon.ru/at/measure/measure/prefix/metric"
)

// CPU ...
type CPU struct {
	value[metric.Prefix]
}

// NewCPU ...
func NewCPU[T measure.Quantitiable](v T, p metric.Prefix) CPU {
	return CPU{value: newValue(v, p)}
}

// Sum ...
func (c CPU) Sum(other CPU) CPU {
	c.value = c.value.Sum(other.value)

	return c
}

// ToPrefix ...
func (c CPU) ToPrefix(prefix metric.Prefix) CPU {
	c.value = c.value.ToPrefix(prefix)

	return c
}

// Multiply ...
func (c CPU) Multiply(k measure.Quantity) CPU {
	c.value = c.value.Multiply(k)

	return c
}

// Magnitude ...
func (c CPU) Magnitude() string {
	return c.Prefix().String() + "CPU"
}

// Format ...
func (c CPU) Format(prec int) string {
	c.value = c.value.Prettier()

	return fmt.Sprintf("%s %s", c.value.Format(prec), c.Magnitude())
}
