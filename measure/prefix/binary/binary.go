package binary

// Prefix ...
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

// String ...
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

// Value ...
func (b Prefix) Value() float64 {
	return float64(b)
}

// Scale returns a prefix like scale, so all prefixes in asc order
func (b Prefix) Scale() []Prefix {
	return []Prefix{None, Kibi, Mebi, Gibi, Tebi, Pebi}
}
