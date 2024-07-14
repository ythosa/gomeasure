// Package binary contains implementation of binary unit prefix such as Kibi, Mebi, etc.
package binary

// Prefix represents binary prefix
type Prefix float64

// Binary prefixes
const (
	None Prefix = 1 << (10 * iota)
	Kibi
	Mebi
	Gibi
	Tebi
	Pebi
)

// String returns prefix string value
func (b Prefix) String() string {
	switch b {
	case Kibi:
		return "Ki"
	case Mebi:
		return "Mi"
	case Gibi:
		return "Gi"
	case Tebi:
		return "Ti"
	case Pebi:
		return "Pi"
	default:
		return ""
	}
}

// Value returns prefix num value
func (b Prefix) Value() float64 {
	return float64(b)
}

// Scale returns a prefix like scale, so all prefixes in asc order
func (b Prefix) Scale() []Prefix {
	return []Prefix{None, Kibi, Mebi, Gibi, Tebi, Pebi}
}
