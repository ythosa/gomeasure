// Package units provides implementations of gomeasure.Value
package units

import (
	"strconv"
	"strings"

	"github.com/samber/lo"
	"github.com/ythosa/gomeasure"
	"github.com/ythosa/gomeasure/prefix"
)

// unit is base class for any unit
type unit[P prefix.Prefix[P]] struct {
	quantity gomeasure.Quantity
	prefix   P
}

// newValue ...
func newValue[T gomeasure.Quantitiable, P prefix.Prefix[P]](quantity T, prefix P) unit[P] {
	return unit[P]{quantity: gomeasure.Quantity(quantity), prefix: prefix}
}

// Sum returns new instance of Value as sum of current and other
func (m unit[P]) Sum(other unit[P]) unit[P] {
	other = other.ToPrefix(m.prefix)
	m.quantity += other.quantity

	return m
}

// Multiply multiplies unit on k and returns new instance
func (m unit[P]) Multiply(k gomeasure.Quantity) unit[P] {
	m.quantity *= k

	return m
}

// ToPrefix leads unit to another prefix
func (m unit[P]) ToPrefix(prefix P) unit[P] {
	m.quantity = m.quantity * gomeasure.Quantity(m.prefix.Value()) / gomeasure.Quantity(prefix.Value())
	m.prefix = prefix

	return m
}

func (m unit[P]) Quantity() gomeasure.Quantity {
	return m.quantity
}

func (m unit[P]) Prefix() P {
	return m.prefix
}

func (m unit[P]) prettier() unit[P] {
	scale := m.prefix.Scale()
	if len(scale) == 0 {
		return m
	}

	sgn := gomeasure.Quantity(lo.Ternary(m.Quantity() > 0, 1, -1))
	result := m.ToPrefix(scale[0]).Multiply(sgn)

	for i := 1; i < len(scale); i++ {
		minimized := result.ToPrefix(scale[i])
		if int(minimized.Quantity()) > 0 {
			result = minimized
		} else {
			break
		}
	}

	return result.Multiply(sgn)
}

func (m unit[P]) Format(prec int) string {
	return m.spaced(m.stripTrailingZeros(strconv.FormatFloat(float64(m.quantity), 'f', prec, 64)))
}

func (m unit[P]) spaced(value string) string {
	parts := strings.Split(value, ".")
	integerPart := []rune(parts[0])

	return string(lo.Reverse([]rune(strings.Join(lo.ChunkString(string(lo.Reverse(integerPart)), 3), " "))))
}

func (m unit[P]) stripTrailingZeros(s string) string {
	if !strings.ContainsRune(s, '.') {
		return s
	}

	offset := len(s) - 1
	for offset > 0 {
		if s[offset] == '.' {
			offset--
			break
		}

		if s[offset] != '0' {
			break
		}
		offset--
	}

	return s[:offset+1]
}
