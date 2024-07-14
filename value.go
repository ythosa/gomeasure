package gomeasure

// Value is common interface for units
type Value interface {
	Quantity() Quantity
	Magnitude() string
	Format(prec int) string
}
