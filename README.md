# gomeasure

Library for typed measurements. Allows:
- at the code level, see which units are being worked with
- perform arithmetic operations only with "compatible" values
- work with a list of typed values as with abstract values (there are no arithmetic operations)

> By measurement we mean pair of values: `(quantity, magnitude)`, where `magnitude` consists of two
> parts: `(prefix, base_unit)`. 1 mCPU <=> `quantity` = 1, `prefix` = m, `base unit` = CPU

## Available kinds
- `Amount` — quantity
- `CPU` — CPU count
- `Second` — seconds count
- `Byte` — bytes count

## Nuances
- There are two types of prefixes, each of which is type-determined:
    - `Binary` — `Byte`
    - `Metric` — all the others

## Usage example
```go
package main

import (
    "github.com/ythosa/gomeasure/units"
    "github.com/ythosa/gomeasure/prefix/binary"
)

func main() {
	var memory = units.NewByte(500, binary.Tebi)
	const replicaCount = 3

	println(memory.Multiply(replicaCount).Format(2)) // -> 1.46 PiB
}
```
