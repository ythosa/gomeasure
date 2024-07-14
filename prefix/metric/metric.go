// Package metric contains implementation of binary unit prefix such as Milli, etc.
package metric

// Prefix represents metric prefix
type Prefix float64

// Metric prefixes
const (
	Milli Prefix = 1e-3
	None  Prefix = 1
)

// String returns prefix string value
func (m Prefix) String() string {
	switch m {
	case Milli:
		return "m"
	default:
		return ""
	}
}

// Value returns prefix num value
func (m Prefix) Value() float64 {
	return float64(m)
}

// Scale returns a prefix like scale, so all prefixes in asc order
func (m Prefix) Scale() []Prefix {
	return []Prefix{Milli, None}
}
