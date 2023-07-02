package measure

import "golang.org/x/exp/constraints"

// Quantitiable ...
type Quantitiable interface {
	constraints.Integer | constraints.Float
}
