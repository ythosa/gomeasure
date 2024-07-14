// Package gomeasure contains common types for all units
package gomeasure

import "golang.org/x/exp/constraints"

// Quantitiable represents interface for values, which convertable to Quantity
type Quantitiable interface {
	constraints.Integer | constraints.Float
}
