package units

import (
	"strconv"
	"strings"

	"github.com/samber/lo"
	"gitlab.ozon.ru/at/measure"
	"gitlab.ozon.ru/at/measure/prefix"
)

// Value is type represents measurement
type value[P prefix.Prefix[P]] struct {
	quantity measure.Quantity
	prefix   P
}

// newValue ...
func newValue[T measure.Quantitiable, P prefix.Prefix[P]](quantity T, prefix P) value[P] {
	return value[P]{quantity: measure.Quantity(quantity), prefix: prefix}
}

// Sum returns new instance of Value as sum of current and other
func (m value[P]) Sum(other value[P]) value[P] {
	other = other.ToPrefix(m.prefix)
	m.quantity += other.quantity

	return m
}

// Multiply multiplies value on k and returns new instance
func (m value[P]) Multiply(k measure.Quantity) value[P] {
	m.quantity *= k

	return m
}

// ToPrefix leads value to another prefix
func (m value[P]) ToPrefix(prefix P) value[P] {
	m.quantity = m.quantity * measure.Quantity(m.prefix.Value()) / measure.Quantity(prefix.Value())
	m.prefix = prefix

	return m
}

func (m value[P]) Quantity() measure.Quantity {
	return m.quantity
}

func (m value[P]) Prefix() P {
	return m.prefix
}

func (m value[P]) Prettier() value[P] {
	scale := m.prefix.Scale()
	if len(scale) == 0 {
		return m
	}

	var result = m.ToPrefix(scale[0])
	for i := 1; i < len(scale); i++ {
		minimized := result.ToPrefix(scale[i])
		if int(minimized.Quantity()) > 0 {
			result = minimized
		} else {
			break
		}
	}

	return result
}

func (m value[P]) Format(prec int) string {
	return m.spaced(m.stripTrailingZeros(strconv.FormatFloat(float64(m.quantity), 'f', prec, 64)))
}

func (m value[P]) spaced(value string) string {
	parts := strings.Split(value, ".")
	integerPart := []rune(parts[0])

	return string(lo.Reverse([]rune(strings.Join(lo.ChunkString(string(lo.Reverse(integerPart)), 3), " "))))
}

func (m value[P]) stripTrailingZeros(s string) string {
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
