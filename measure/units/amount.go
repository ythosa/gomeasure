package units

import (
	"fmt"

	"gitlab.ozon.ru/at/measure/measure"
	"gitlab.ozon.ru/at/measure/measure/prefix/metric"
)

// Amount ...
type Amount struct {
	value[metric.Prefix]
}

// NewAmount ...
func NewAmount[T measure.Quantitiable](v T, p metric.Prefix) Amount {
	return Amount{value: newValue(v, p)}
}

// Sum ...
func (a Amount) Sum(other Amount) Amount {
	a.value = a.value.Sum(other.value)

	return a
}

// ToPrefix ...
func (a Amount) ToPrefix(prefix metric.Prefix) Amount {
	a.value = a.value.ToPrefix(prefix)

	return a
}

// Multiply ...
func (a Amount) Multiply(k measure.Quantity) Amount {
	a.value = a.value.Multiply(k)

	return a
}

// Magnitude ...
func (a Amount) Magnitude() string {
	return "" + a.prefix.String()
}

// Format ...
func (a Amount) Format(prec int) string {
	a.value = a.value.Prettier()

	return fmt.Sprintf("%s %s", a.value.Format(prec), a.Magnitude())
}
