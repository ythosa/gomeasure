package units

import (
	"fmt"

	"github.com/ythosa/gomeasure"
	"github.com/ythosa/gomeasure/prefix/binary"
)

// Byte represents byte unit
type Byte struct {
	unit[binary.Prefix]
}

// NewByte returns new instance of Byte
func NewByte[T gomeasure.Quantitiable](v T, p binary.Prefix) Byte {
	return Byte{unit: newValue(v, p)}
}

// Sum returns sum of current and other CPU unit
func (b Byte) Sum(other Byte) Byte {
	b.unit = b.unit.Sum(other.unit)

	return b
}

// ToPrefix converts from current prefix to passed
func (b Byte) ToPrefix(prefix binary.Prefix) Byte {
	b.unit = b.unit.ToPrefix(prefix)

	return b
}

// Multiply returns current unit multiplied to k
func (b Byte) Multiply(k gomeasure.Quantity) Byte {
	b.unit = b.unit.Multiply(k)

	return b
}

// Magnitude returns string representation of magnitude (prefix + "B")
func (b Byte) Magnitude() string {
	return b.Prefix().String() + "B"
}

// Format returns formatted quantity with passed preciseness and appends magnitude
func (b Byte) Format(prec int) string {
	b.unit = b.unit.prettier()

	return fmt.Sprintf("%s %s", b.unit.Format(prec), b.Magnitude())
}
