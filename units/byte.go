package units

import (
	"fmt"

	"gitlab.ozon.ru/at/measure"
	"gitlab.ozon.ru/at/measure/prefix/binary"
)

// Byte ...
type Byte struct {
	value[binary.Prefix]
}

// NewByte ...
func NewByte[T measure.Quantitiable](v T, p binary.Prefix) Byte {
	return Byte{value: newValue(v, p)}
}

// Sum ...
func (b Byte) Sum(other Byte) Byte {
	b.value = b.value.Sum(other.value)

	return b
}

// ToPrefix ...
func (b Byte) ToPrefix(prefix binary.Prefix) Byte {
	b.value = b.value.ToPrefix(prefix)

	return b
}

// Multiply ...
func (b Byte) Multiply(k measure.Quantity) Byte {
	b.value = b.value.Multiply(k)

	return b
}

// Magnitude ...
func (b Byte) Magnitude() string {
	return b.Prefix().String() + "B"
}

// Format ...
func (b Byte) Format(prec int) string {
	b.value = b.value.Prettier()

	return fmt.Sprintf("%s %s", b.value.Format(prec), b.Magnitude())
}
