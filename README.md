# Measure
Библиотека для типизированных измерений. Позволяет:
- на уровне кода видеть, с какими единицами производится работа
- производить арифметические операции только с "совместимыми" значениями
- работать со списком типизированных значений, как с абстрактными значениями (отсутствуют арифметические операции)

> Под измерением имеется в виду пара значений `(quantity, magnitude)`, где `magnitude` состоит из двух частей `(prefix, base_unit)`. 1 mCPU <=> `quantity` = 1, `prefix` = m, `base_unit` = CPU

## Доступные типы
- `Amount` — количество
- `CPU` — ядра
- `Second` — секунды
- `Byte` — байты

## Нюансы
- Существует два вида префиксов, каждый из которых детерминирован типом:
  - `Binary` — `Byte`
  - `Metric` — все остальные

## Пример использования
```go
package main

import (
	"gitlab.ozon.ru/at/measure/prefix/binary"
	"gitlab.ozon.ru/at/measure/units"
)

func main() {
	var memory = units.NewByte(500, binary.Tebi)
	var replicaCount = 3
	
	println(memory.Multiply(replicaCount).Format(2)) // -> 1.46 PiB
}
```
