package metric

// Prefix ...
type Prefix float64

// Metric prefixes
const (
	Milli Prefix = 1e-3
	None  Prefix = 1
)

// String ...
func (m Prefix) String() string {
	switch m {
	case Milli:
		return "m"
	default:
		return ""
	}
}

// Value ...
func (m Prefix) Value() float64 {
	return float64(m)
}

// Scale returns a prefix like scale, so all prefixes in asc order
func (m Prefix) Scale() []Prefix {
	return []Prefix{Milli, None}
}
